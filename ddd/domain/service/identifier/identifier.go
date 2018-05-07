package identifier

import (
	"github.com/rs/xid"
)

var (
	guid = xid.New()
)

func GenerateID() string {
	return guid.String()
}
