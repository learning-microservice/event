package event

type ID uint

func (id *ID) IsNotEmpty() bool {
	return uint(*id) > 0
}
