package memory

import (
	"testing"

	"github.com/learning-microservice/event/ddd/domain"
)

type txManager struct{}

func (t *txManager) WithTx(f func(domain.Session) error) (err error) {
	err = f(t)
	return
}

func (t *txManager) WithReadOnly(f func(domain.Session) error) (err error) {
	return f(t)
}

func NewTxManager(*testing.T) domain.TxContext {
	return &txManager{}
}
