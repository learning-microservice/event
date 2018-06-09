package db

import (
	"github.com/jinzhu/gorm"
)

// TxFunc is ...
type TxFunc func(*gorm.DB) error

// WithTx is ...
func WithTx(f TxFunc) (err error) {
	if err = lazyInit(); err != nil {
		return
	}
	tx := cacheDB.Begin()
	if err = tx.Error; err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	err = f(tx)
	return
}

// WithReadOnly is ...
func WithReadOnly(f func(*gorm.DB) error) (err error) {
	if err = lazyInit(); err != nil {
		return
	}
	return f(cacheDB)
}
