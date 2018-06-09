package generator

import (
	"github.com/learning-microservice/event/mvc/commons/types/event"
	"github.com/rs/xid"
)

// NextUID is ...
func NextUID() event.UID {
	return event.UID(xid.New().String())
}
