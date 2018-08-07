package postgres

import (
	"context"
	"fmt"
	"time"

	pg "gopkg.in/pg.v3"

	"github.com/p1cn/tantan-backend-common/config"
	slog "github.com/p1cn/tantan-backend-common/log"
)

type pgDB struct {
	*pg.DB
	options       *pg.Options
	infoCollector infoCollector
	Name          string
}

func newPgDB(ic infoCollector, name string, db *pg.DB, opt *pg.Options) *pgDB {
	return &pgDB{
		DB:            db,
		options:       opt,
		infoCollector: ic,
		Name:          name,
	}
}

func configPostgreSqlToPGOptions(pgConfig config.PostgreSql) *pg.Options {
	idle := pgConfig.IdleTimeout
	if idle == 0 {
		idle = 180
	} else if idle < 0 {
		idle = 0
	}

	return &pg.Options{
		Host:               pgConfig.Address,
		Port:               pgConfig.Port,
		User:               pgConfig.User,
		Database:           pgConfig.Database,
		Password:           pgConfig.Password,
		SSL:                false,
		PoolSize:           pgConfig.PoolSize,
		IdleTimeout:        time.Duration(idle) * time.Second,
		PoolTimeout:        time.Duration(pgConfig.PoolTimeout) * time.Second,
		IdleCheckFrequency: time.Duration(pgConfig.IdleCheckFrequency) * time.Second,
	}
}

func (pdb *pgDB) String() string {
	return fmt.Sprintf("DB %s Addr %s:%s", pdb.options.Database, pdb.options.Host, pdb.options.Port)
}

func (pdb *pgDB) ExecWithContext(ctx context.Context, q string, args ...interface{}) (pg.Result, error) {
	return pdb.Exec(q, args...)
}

func (pdb *pgDB) Exec(q string, args ...interface{}) (pg.Result, error) {

	record := pdb.infoCollector.Timer()

	result, err := pdb.DB.Exec(q, toDbTimes(args...)...)

	if err == pg.ErrNoRows {
		err = nil
	}

	if err != nil {
		slog.Err("%s %s %v err %s", pdb, q, args, err.Error())
		record(pdb.Name, getFuncName(q), "error")
	} else {
		record(pdb.Name, getFuncName(q), "OK")
	}

	return result, err
}

func (pdb *pgDB) ExecOneWithContext(ctx context.Context, q string, args ...interface{}) (pg.Result, error) {
	return pdb.ExecOne(q, args...)
}

func (pdb *pgDB) ExecOne(q string, args ...interface{}) (pg.Result, error) {

	record := pdb.infoCollector.Timer()

	result, err := pdb.DB.ExecOne(q, toDbTimes(args...)...)

	if err == pg.ErrNoRows {
		err = nil
	}

	if err != nil {
		slog.Err("%s %s %v err %s", pdb, q, args, err.Error())
		record(pdb.Name, getFuncName(q), "error")
	} else {
		record(pdb.Name, getFuncName(q), "OK")
	}

	return result, err
}

func (pdb *pgDB) QueryWithContext(ctx context.Context, coll pg.Collection, q string, args ...interface{}) (pg.Result, error) {
	return pdb.Query(coll, q, args...)
}

func (pdb *pgDB) Query(coll pg.Collection, q string, args ...interface{}) (pg.Result, error) {
	record := pdb.infoCollector.Timer()

	result, err := pdb.DB.Query(coll, q, toDbTimes(args...)...)

	if err == pg.ErrNoRows {
		err = nil
	}

	if err != nil {
		slog.Err("%s %s %v err %s", pdb, q, args, err.Error())
		record(pdb.Name, getFuncName(q), "error")
	} else {
		record(pdb.Name, getFuncName(q), "OK")
	}

	return result, err
}

func (pdb *pgDB) QueryOneWithContext(ctx context.Context, record interface{}, q string, args ...interface{}) (pg.Result, error) {
	return pdb.QueryOne(record, q, args...)
}

func (pdb *pgDB) QueryOne(record interface{}, q string, args ...interface{}) (pg.Result, error) {
	recordMetrics := pdb.infoCollector.Timer()

	result, err := pdb.DB.QueryOne(record, q, toDbTimes(args...)...)

	if err == pg.ErrNoRows {
		err = nil
	}
	if err != nil {
		slog.Err("%s %s %v err %s", pdb, q, args, err.Error())
		recordMetrics(pdb.Name, getFuncName(q), "error")
	} else {
		recordMetrics(pdb.Name, getFuncName(q), "OK")
	}

	return result, err
}
