package context

type Transaction interface {
	WithTx(func(Session) error) error
	WithReadOnly(func(Session) error) error
}

type Session interface{}
