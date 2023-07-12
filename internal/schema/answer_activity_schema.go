package schema

// AcceptAnswerOperationInfo accept answer operation info
type AcceptAnswerOperationInfo struct {
	QuestionObjectID string
	QuestionUserID   string
	AnswerObjectID   string
	AnswerUserID     string

	// vote activity info
	Activities []*AcceptAnswerActivity
}

// AcceptAnswerActivity accept answer activity
type AcceptAnswerActivity struct {
	ActivityType     int
	ActivityUserID   string
	TriggerUserID    string
	OriginalObjectID string
	Rank             int
}

func (v *AcceptAnswerActivity) HasRank() int {
	if v.Rank != 0 {
		return 1
	}
	return 0
}

func (a *AcceptAnswerOperationInfo) GetUserIDs() (userIDs []string) {
	for _, act := range a.Activities {
		userIDs = append(userIDs, act.ActivityUserID)
	}
	return userIDs
}
