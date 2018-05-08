package event

////////////////////////////////////////////
// Bookings
////////////////////////////////////////////

type Bookings []Booking

func (bs *Bookings) add(id ID, attendeeID AttendeeID) error {
	for _, book := range *bs {
		if book.attendeeID == attendeeID {
			return errorBookingFailure("event already booked")
		}
	}
	*bs = append(*bs, Booking{
		eventID:    id,
		attendeeID: attendeeID,
	})
	return nil
}

////////////////////////////////////////////
// Booking
////////////////////////////////////////////

type Booking struct {
	id         uint64
	eventID    ID
	attendeeID AttendeeID
	isCanceled bool // TODO 未登録の予約に対するキャンセルは削除
}
