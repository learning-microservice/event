package domain

type TxContext interface {
	WithTx(func(Session) error) error
	WithReadOnly(func(Session) error) error
}

type Session interface {
	//OperatorID() string
}
