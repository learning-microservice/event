package events

////////////////////////////////////////////
// Bookings
////////////////////////////////////////////

type Bookings []Booking

func (bs *Bookings) contains(attendeeID AttendeeID) bool {
	for _, book := range *bs {
		if book.attendeeID == attendeeID {
			return true
		}
	}
	return false
}

func (bs *Bookings) add(attendeeID AttendeeID) {
	if !bs.contains(attendeeID) {
		*bs = append(*bs, Booking{
			attendeeID: attendeeID,
		})
	}
}

////////////////////////////////////////////
// Booking
////////////////////////////////////////////

type Booking struct {
	id         uint
	attendeeID AttendeeID
	isCanceled bool // TODO 未登録の予約に対するキャンセルは削除
}
