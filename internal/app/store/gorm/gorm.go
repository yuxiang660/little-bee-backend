// Wrapper of gorm. Do not use gorm directly in the project.

package gorm

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	gorm "github.com/jinzhu/gorm"
	"github.com/yuxiang660/little-bee-server/internal/app/store"

	// gorm inject
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type options struct {
	debug        bool
	DBType       string
	DSN          string
	maxLifetime  int
	maxOpenConns int
	maxIdleConns int
}

// Option defines function signature to set options.
type Option func(*options)

// SetDebug returns an action to set debug flag.
func SetDebug(debug bool) Option {
	return func(o *options) {
		o.debug = debug
	}
}

// SetDBType returns an action to set database type.
func SetDBType(DBType string) Option {
	return func(o *options) {
		o.DBType = DBType
	}
}

// SetDSN returns an action to set DSN for database connenction.
func SetDSN(DSN string) Option {
	return func(o *options) {
		o.DSN = DSN
	}
}

// SetMaxLifetime returns an action to set max life time for a connenction.
func SetMaxLifetime(maxLifetime int) Option {
	return func(o *options) {
		o.maxLifetime = maxLifetime
	}
}

// SetMaxOpenConns returns an action to set max number of connections.
func SetMaxOpenConns(maxOpenConns int) Option {
	return func(o *options) {
		o.maxOpenConns = maxOpenConns
	}
}

// SetMaxIdleConns returns an action to set max number of connections in the idle connection pool.
func SetMaxIdleConns(maxIdleConns int) Option {
	return func(o *options) {
		o.maxIdleConns = maxIdleConns
	}
}

type storeGorm struct {
	db *gorm.DB
}

// New creates an autherJWT object based on user configuration.
func New(opts ...Option) (store.Store, error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	switch o.DBType {
	case "sqlite3":
		_ = os.MkdirAll(filepath.Dir(o.DSN), 0777)
	default:
		return nil, errors.New("Unknown Database")
	}

	db, err := gorm.Open(o.DBType, o.DSN)
	if err != nil {
		return nil, err
	}

	if o.debug {
		db = db.Debug()
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(o.maxIdleConns)
	db.DB().SetMaxOpenConns(o.maxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(o.maxLifetime) * time.Second)

	return &storeGorm{db: db}, nil
}

// Close close current db connection.  If database connection is not an io.Closer, returns an error.
func (s *storeGorm)Close() error{
	return s.db.Close()
}