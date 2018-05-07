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
	isCanceled bool
}

func (a *Assignment) ID() uint64 {
	return a.id
}

func (a *Assignment) EventID() ID {
	return a.eventID
}

func (a *Assignment) AssigneeID() AssigneeID {
	return a.assigneeID
}

func (a *Assignment) IsNew() bool {
	return a.id == 0 && !a.isCanceled
}

func (a *Assignment) IsCanceled() bool {
	return a.id > 0 && a.isCanceled
}
