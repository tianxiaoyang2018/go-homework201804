package postgres

import (
	"context"

	"github.com/p1cn/tantan-backend-common/config"

	pg "gopkg.in/pg.v3"
)

type DbOperation int

const (
	DbRead DbOperation = 1 << (1 + iota)
	DbWrite
)

type DBWrapper interface {
	Exec(q string, args ...interface{}) (pg.Result, error)
	ExecOne(q string, args ...interface{}) (pg.Result, error)
	Query(coll pg.Collection, q string, args ...interface{}) (pg.Result, error)
	QueryOne(record interface{}, q string, args ...interface{}) (pg.Result, error)

	ExecWithContext(ctx context.Context, q string, args ...interface{}) (pg.Result, error)
	ExecOneWithContext(ctx context.Context, q string, args ...interface{}) (pg.Result, error)
	QueryWithContext(ctx context.Context, coll pg.Collection, q string, args ...interface{}) (pg.Result, error)
	QueryOneWithContext(ctx context.Context, record interface{}, q string, args ...interface{}) (pg.Result, error)

	RunInTransaction(fn func(tx *pg.Tx) error) error
}

func newPGDBWrapper(pgConfig config.PostgreSql) DBWrapper {
	opt := configPostgreSqlToPGOptions(pgConfig)
	return &pgDB{DB: pg.Connect(opt), options: opt}
}
