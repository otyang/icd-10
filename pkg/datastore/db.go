package datastore

import (
	"database/sql"
	"log"
	"runtime"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	_ "github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

const (
	DriverPg     = "pg"
	DriverSqlite = sqliteshim.ShimName
)

type (
	OrmDB   = *bun.DB
	OrmDbTx = bun.IDB
)

// disable ssl on localhost
// dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"

func NewDBConnection(dbDriver, dsn string, dbPoolMax int, printQueriesToStdout bool) OrmDB {
	_dbh, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}

	var db *bun.DB
	switch dbDriver {
	case DriverSqlite:
		db = bun.NewDB(_dbh, sqlitedialect.New(), bun.WithDiscardUnknownColumns())
	case DriverPg:
		db = bun.NewDB(_dbh, pgdialect.New(), bun.WithDiscardUnknownColumns())
	default:
		log.Fatalf("unknown db driver: %s", dbDriver)
	}

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(printQueriesToStdout),
	))

	db.SetMaxOpenConns(dbPoolMax * runtime.GOMAXPROCS(0))
	db.SetMaxIdleConns(dbPoolMax * runtime.GOMAXPROCS(0))

	return db
}
