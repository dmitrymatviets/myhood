package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

const ContextTransactionKey = "db.transaction"

type DatabaseConfig struct {
	Server       string   `envconfig:"server"`
	ReadReplicas []string `envconfig:"readreplicas"`
	User         string   `envconfig:"user"`
	Password     string   `envconfig:"password"`
	Port         int      `envconfig:"port"`
	DatabaseName string   `envconfig:"database"`
	MaxOpenConns int      `envconfig:"maxopenconns"`
	MaxIdleConns int      `envconfig:"maxidleconns"`
}

type Database struct {
	*sqlx.DB
	readReplicaDbs []*sqlx.DB
}

func (d *Database) GetReadReplicaDb() *sqlx.DB {
	rand.Seed(time.Now().UnixNano())
	return d.readReplicaDbs[rand.Intn(len(d.readReplicaDbs))]
}

func NewDatabase(c DatabaseConfig) (*Database, error) {
	database := &Database{}
	db, err := newDb(c)
	if err != nil {
		return nil, err
	}
	database.DB = db
	for _, replicaHost := range c.ReadReplicas {
		cfgClone := c
		cfgClone.Server = replicaHost
		replicaDb, err := newDb(c)
		if err != nil {
			return nil, err
		}
		database.readReplicaDbs = append(database.readReplicaDbs, replicaDb)
	}
	return database, nil
}

func newDb(c DatabaseConfig) (*sqlx.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.User,
		c.Password,
		c.Server,
		c.Port,
		c.DatabaseName,
	)

	db, err := sqlx.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)

	return db, nil
}

func (db *Database) commit(ctx *context.Context) error {
	value := (*ctx).Value(ContextTransactionKey)
	*ctx = context.WithValue(*ctx, ContextTransactionKey, nil)
	tx, ok := value.(*sqlx.Tx)
	if ok {
		return tx.Commit()
	} else {
		return errors.New("cannot commit: transaction not found")
	}
}

func (db *Database) rollback(ctx *context.Context) error {
	value := (*ctx).Value(ContextTransactionKey)
	*ctx = context.WithValue(*ctx, ContextTransactionKey, nil)
	tx, ok := value.(*sqlx.Tx)
	if ok {
		return tx.Rollback()
	} else {
		return errors.New("cannot rollback: transaction not found")
	}
}

func (db *Database) transaction(ctx *context.Context) (*sqlx.Tx, error) {
	var err error
	value := (*ctx).Value(ContextTransactionKey)
	tx, ok := value.(*sqlx.Tx)
	if ok {
		return tx, nil
	}

	tx, err = db.Beginx()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	*ctx = context.WithValue(*ctx, ContextTransactionKey, tx)
	return tx, nil
}

func (db *Database) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	_, err = db.transaction(&ctx)
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			_ = db.rollback(&ctx)
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			_ = db.rollback(&ctx)
		}
	}()

	err = fn(ctx)
	if err == nil {
		err = db.commit(&ctx)
		if err == sql.ErrTxDone {
			return nil
		}
	}
	return err
}

type TxOrDb interface {
	sqlx.Execer
	sqlx.ExecerContext
	sqlx.Queryer
	sqlx.QueryerContext
	sqlx.Preparer
	sqlx.PreparerContext
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	DriverName() string
	Get(dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// абстракция для tx или db, из контекста
// если в контексте есть транзакция, то возвращает ее (как интерфейс TxOrDb)
// если транзакции нет, то возвращает db (как интерфейс TxOrDb)
func (db *Database) TxOrDbFromContext(ctx context.Context) TxOrDb {
	value := (ctx).Value(ContextTransactionKey)
	tx, ok := value.(*sqlx.Tx)
	if ok {
		return tx
	}
	return db
}
