package event

import (
	"encoding/json"
)

type Tags []string

func (t Tags) JSON() []byte {
	if len(t) > 0 {
		b, _ := json.Marshal(t)
		return b
	} else {
		return []byte("[]")
	}
}
