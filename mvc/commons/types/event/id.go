package event

// ID is ...
type ID uint

// IsEmpty is ...
func (id ID) IsEmpty() bool {
	return uint(id) == 0
}

// IsNotEmpty is ...
func (id ID) IsNotEmpty() bool {
	return !id.IsEmpty()
}
