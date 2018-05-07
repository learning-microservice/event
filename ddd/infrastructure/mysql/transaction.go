package mysql

import (
	"github.com/learning-microservice/core/db"
	"github.com/learning-microservice/event/ddd/domain/context"
)

type transaction struct {
}

func (*transaction) WithTx(f func(context.Session) error) (err error) {
	tx := db.MySQL().Begin()
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
	err = f(&db.DB{tx})
	return
}

func (*transaction) WithReadOnly(f func(context.Session) error) error {
	return f(db.MySQL())
}

func NewTransactionContext() context.Transaction {
	return &transaction{}
}
