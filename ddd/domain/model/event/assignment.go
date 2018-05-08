package event

////////////////////////////////////////////
// Assignments
////////////////////////////////////////////

type Assignments []Assignment

func (as *Assignments) add(id ID, assigneeID AssigneeID) error {
	for _, assign := range *as {
		if assign.assigneeID == assigneeID {
			return errorBookingFailure("event already booked")
		}
	}
	*as = append(*as, Assignment{
		eventID:    id,
		assigneeID: assigneeID,
	})
	return nil
}

////////////////////////////////////////////
// Assignment
////////////////////////////////////////////

type Assignment struct {
	id         uint64
	eventID    ID
	assigneeID AssigneeID
	isCanceled bool // TODO 未登録の予約に対するキャンセルは削除
}
