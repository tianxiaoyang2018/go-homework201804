package postgres

import (
	"context"
	"time"

	"github.com/p1cn/tantan-backend-common/config"
	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/util"
	"github.com/sony/gobreaker"
	pg "gopkg.in/pg.v3"
)

type pgDBWithCB struct {
	*pgDB
	cb *util.CircuitBreaker
}

type pgDBWithSonyCB struct {
	*pgDB
	cb *gobreaker.CircuitBreaker
}

func newDBWrapperWithCB(dbName string, ic infoCollector, pgConfig config.PostgreSql) DBWrapper {
	if pgConfig.CircuitBreaker.Disabled {
		opt := configPostgreSqlToPGOptions(pgConfig)
		return newPgDB(ic, dbName, pg.Connect(opt), opt)
	}

	if len(pgConfig.CircuitBreaker.Type) > 0 {
		slog.Info("install original circuit breaker for db %v.", pgConfig.Database)
		return newDBWrapperWithOriginalCB(dbName, ic, pgConfig)
	}
	// default is sony/CircuitBreaker
	slog.Info("install sony circuit breaker for db %v.", pgConfig.Database)
	return newDBWrapperWithSonyCB(dbName, ic, pgConfig)
}

func newDBWrapperWithSonyCB(dbName string, ic infoCollector, pgConfig config.PostgreSql) DBWrapper {
	opt := configPostgreSqlToPGOptions(pgConfig)

	var totalFailureThreshold uint32 = 5
	if pgConfig.CircuitBreaker.TotalFailureThreshold > 0 {
		totalFailureThreshold = pgConfig.CircuitBreaker.TotalFailureThreshold
	}

	var openStateDuration uint32 = 10
	if pgConfig.CircuitBreaker.RetryDuration > 0 {
		openStateDuration = pgConfig.CircuitBreaker.RetryDuration
	}

	var interval uint32 = 10
	if pgConfig.CircuitBreaker.ClearCountInterval > 0 {
		interval = pgConfig.CircuitBreaker.ClearCountInterval
	}

	var maxRequest uint32 = 0
	if pgConfig.CircuitBreaker.MaxRequestsOnHalfOpenStatus > 0 {
		maxRequest = uint32(pgConfig.CircuitBreaker.MaxRequestsOnHalfOpenStatus)
	}

	return &pgDBWithSonyCB{
		pgDB: newPgDB(ic, dbName, pg.Connect(opt), opt),

		cb: gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name: pgConfig.Database + ":" + pgConfig.Address,

			// Interval is the cyclic period of the closed state
			// for the CircuitBreaker to clear the internal Counts.
			// If Interval is 0, the CircuitBreaker doesn't clear internal Counts during the closed state.
			Interval: time.Duration(interval) * time.Second,

			// MaxRequests is the maximum number of requests allowed to pass through
			// when the CircuitBreaker is half-open.
			// If MaxRequests is 0, the CircuitBreaker allows only 1 request.
			MaxRequests: maxRequest,

			// Timeout is the period of the open state,
			// after which the state of the CircuitBreaker becomes half-open.
			// If Timeout is 0, the timeout value of the CircuitBreaker is set to 60 seconds.
			Timeout: time.Duration(openStateDuration) * time.Second,

			// ReadyToTrip is called with a copy of Counts whenever a request fails in the closed state.
			// If ReadyToTrip returns true, the CircuitBreaker will be placed into the open state.
			// If ReadyToTrip is nil, default ReadyToTrip is used.
			// Default ReadyToTrip returns true when the number of consecutive failures is more than 5.
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				return counts.TotalFailures >= totalFailureThreshold
			},

			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				slog.Info("breakername=%v, %v->%v ", name, from, to)
			},
		}),
	}
}

// remove this function if sony circuit breaker work well.
func newDBWrapperWithOriginalCB(dbName string, ic infoCollector, pgConfig config.PostgreSql) DBWrapper {
	opt := configPostgreSqlToPGOptions(pgConfig)
	cbThreshold, cbRetryDur := int64(10), time.Second*10
	if thr := pgConfig.CircuitBreaker.TotalFailureThreshold; thr != 0 {
		cbThreshold = int64(thr)
	}

	if dur := pgConfig.CircuitBreaker.RetryDuration; dur != 0 {
		cbRetryDur = time.Duration(dur) * time.Second
	}

	pgDB := newPgDB(ic, dbName, pg.Connect(opt), opt)

	return &pgDBWithCB{pgDB: pgDB, cb: util.NewCircuitBreaker(cbThreshold, cbRetryDur, pgDB)}
}

func (pdb *pgDBWithSonyCB) ExecWithContext(ctx context.Context, q string, args ...interface{}) (pg.Result, error) {
	return pdb.Exec(q, args...)
}

func (pdb *pgDBWithSonyCB) Exec(q string, args ...interface{}) (pg.Result, error) {
	var result pg.Result
	var err error
	_, err = pdb.cb.Execute(func() (interface{}, error) {
		result, err = pdb.pgDB.Exec(q, args...)
		return result, err
	})
	return result, err
}

func (pdb *pgDBWithSonyCB) ExecOneWithContext(ctx context.Context, q string, args ...interface{}) (pg.Result, error) {
	return pdb.ExecOne(q, args...)
}

func (pdb *pgDBWithSonyCB) ExecOne(q string, args ...interface{}) (pg.Result, error) {
	var result pg.Result
	var err error
	_, err = pdb.cb.Execute(func() (interface{}, error) {
		result, err = pdb.pgDB.ExecOne(q, args...)
		return result, err
	})
	return result, err
}

func (pdb *pgDBWithSonyCB) QueryWithContext(ctx context.Context, coll pg.Collection, q string, args ...interface{}) (pg.Result, error) {
	return pdb.Query(coll, q, args...)
}

func (pdb *pgDBWithSonyCB) Query(coll pg.Collection, q string, args ...interface{}) (pg.Result, error) {
	var result pg.Result
	var err error
	_, err = pdb.cb.Execute(func() (interface{}, error) {
		result, err = pdb.pgDB.Query(coll, q, args...)
		return result, err
	})
	return result, err
}

func (pdb *pgDBWithSonyCB) QueryOneWithContext(ctx context.Context, record interface{}, q string, args ...interface{}) (pg.Result, error) {
	return pdb.QueryOne(record, q, args...)
}

func (pdb *pgDBWithSonyCB) QueryOne(record interface{}, q string, args ...interface{}) (pg.Result, error) {
	var result pg.Result
	var err error
	_, err = pdb.cb.Execute(func() (interface{}, error) {
		result, err = pdb.pgDB.QueryOne(record, q, args...)
		return result, err
	})
	return result, err
}

///////////////////// original circuit breaker function begin ///////////////////////////////

func (pdb *pgDBWithCB) ExecWithContext(ctx context.Context, q string, args ...interface{}) (pg.Result, error) {
	return pdb.Exec(q, args...)
}

func (pdb *pgDBWithCB) Exec(q string, args ...interface{}) (pg.Result, error) {
	var result pg.Result
	var err error
	err = pdb.cb.Run(func() (error, bool) {
		result, err = pdb.pgDB.Exec(q, args...)
		return err, false
	})
	return result, err
}

func (pdb *pgDBWithCB) ExecOneWithContext(ctx context.Context, q string, args ...interface{}) (pg.Result, error) {
	return pdb.ExecOne(q, args...)
}

func (pdb *pgDBWithCB) ExecOne(q string, args ...interface{}) (pg.Result, error) {
	var result pg.Result
	var err error
	err = pdb.cb.Run(func() (error, bool) {
		result, err = pdb.pgDB.ExecOne(q, args...)
		return err, false
	})
	return result, err
}

func (pdb *pgDBWithCB) QueryWithContext(ctx context.Context, coll pg.Collection, q string, args ...interface{}) (pg.Result, error) {
	return pdb.Query(coll, q, args...)
}

func (pdb *pgDBWithCB) Query(coll pg.Collection, q string, args ...interface{}) (pg.Result, error) {
	var result pg.Result
	var err error
	err = pdb.cb.Run(func() (error, bool) {
		result, err = pdb.pgDB.Query(coll, q, args...)
		return err, false
	})
	return result, err
}

func (pdb *pgDBWithCB) QueryOneWithContext(ctx context.Context, record interface{}, q string, args ...interface{}) (pg.Result, error) {
	return pdb.QueryOne(record, q, args...)

}

func (pdb *pgDBWithCB) QueryOne(record interface{}, q string, args ...interface{}) (pg.Result, error) {
	var result pg.Result
	var err error
	err = pdb.cb.Run(func() (error, bool) {
		result, err = pdb.pgDB.QueryOne(record, q, args...)
		return err, false
	})
	return result, err
}
