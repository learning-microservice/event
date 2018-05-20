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

func ToTags(data []byte) Tags {
	var t Tags
	if len(data) > 0 {
		if err := json.Unmarshal(data, &t); err != nil {
			// TODO warning log ?
			panic(err)
		}
	}
	return t
}
