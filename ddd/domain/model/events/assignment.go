package events

////////////////////////////////////////////
// Assignments
////////////////////////////////////////////

type Assignments []Assignment

func (as *Assignments) contains(assigneeID AssigneeID) bool {
	for _, assign := range *as {
		if assign.assigneeID == assigneeID {
			return true
		}
	}
	return false
}

func (as *Assignments) add(assigneeID AssigneeID) {
	if !as.contains(assigneeID) {
		*as = append(*as, Assignment{
			assigneeID: assigneeID,
		})
	}
}

////////////////////////////////////////////
// Assignment
////////////////////////////////////////////

type Assignment struct {
	id         uint
	assigneeID AssigneeID
	isCanceled bool // TODO 未登録の予約に対するキャンセルは削除
}
