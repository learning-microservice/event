package mysql

import (
	"github.com/learning-microservice/event/ddd/domain/context"
)

type txManager struct {
	mySQL *mySQL
}

func (*txManager) WithTx(f func(context.Session) error) (err error) {
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
	err = f(&mySQL{tx})
	return
}

func (*txManager) WithReadOnly(f func(context.Session) error) (err error) {
	if err = lazyInit(); err != nil {
		return
	}
	return f(cacheDB)
}

func NewTxManager() context.Transaction {
	return &txManager{}
}
