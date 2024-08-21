// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/syoi-org/judy/ent/judge"
	"github.com/syoi-org/judy/ent/submission"
)

// CreateJudgeInput represents a mutation input for creating judges.
type CreateJudgeInput struct {
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	Name          string
	Code          string
	Type          *judge.Type
	Configuration string
	ProblemIDs    []int
}

// Mutate applies the CreateJudgeInput on the JudgeMutation builder.
func (i *CreateJudgeInput) Mutate(m *JudgeMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
	m.SetCode(i.Code)
	if v := i.Type; v != nil {
		m.SetType(*v)
	}
	m.SetConfiguration(i.Configuration)
	if v := i.ProblemIDs; len(v) > 0 {
		m.AddProblemIDs(v...)
	}
}

// SetInput applies the change-set in the CreateJudgeInput on the JudgeCreate builder.
func (c *JudgeCreate) SetInput(i CreateJudgeInput) *JudgeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateJudgeInput represents a mutation input for updating judges.
type UpdateJudgeInput struct {
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	Name             *string
	Code             *string
	Type             *judge.Type
	Configuration    *string
	ClearProblems    bool
	AddProblemIDs    []int
	RemoveProblemIDs []int
}

// Mutate applies the UpdateJudgeInput on the JudgeMutation builder.
func (i *UpdateJudgeInput) Mutate(m *JudgeMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Code; v != nil {
		m.SetCode(*v)
	}
	if v := i.Type; v != nil {
		m.SetType(*v)
	}
	if v := i.Configuration; v != nil {
		m.SetConfiguration(*v)
	}
	if i.ClearProblems {
		m.ClearProblems()
	}
	if v := i.AddProblemIDs; len(v) > 0 {
		m.AddProblemIDs(v...)
	}
	if v := i.RemoveProblemIDs; len(v) > 0 {
		m.RemoveProblemIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateJudgeInput on the JudgeUpdate builder.
func (c *JudgeUpdate) SetInput(i UpdateJudgeInput) *JudgeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateJudgeInput on the JudgeUpdateOne builder.
func (c *JudgeUpdateOne) SetInput(i UpdateJudgeInput) *JudgeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateProblemInput represents a mutation input for creating problems.
type CreateProblemInput struct {
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	Name          string
	Code          string
	SubmissionIDs []int
	JudgeID       int
}

// Mutate applies the CreateProblemInput on the ProblemMutation builder.
func (i *CreateProblemInput) Mutate(m *ProblemMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
	m.SetCode(i.Code)
	if v := i.SubmissionIDs; len(v) > 0 {
		m.AddSubmissionIDs(v...)
	}
	m.SetJudgeID(i.JudgeID)
}

// SetInput applies the change-set in the CreateProblemInput on the ProblemCreate builder.
func (c *ProblemCreate) SetInput(i CreateProblemInput) *ProblemCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateProblemInput represents a mutation input for updating problems.
type UpdateProblemInput struct {
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
	Name                *string
	Code                *string
	ClearSubmissions    bool
	AddSubmissionIDs    []int
	RemoveSubmissionIDs []int
	JudgeID             *int
}

// Mutate applies the UpdateProblemInput on the ProblemMutation builder.
func (i *UpdateProblemInput) Mutate(m *ProblemMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Code; v != nil {
		m.SetCode(*v)
	}
	if i.ClearSubmissions {
		m.ClearSubmissions()
	}
	if v := i.AddSubmissionIDs; len(v) > 0 {
		m.AddSubmissionIDs(v...)
	}
	if v := i.RemoveSubmissionIDs; len(v) > 0 {
		m.RemoveSubmissionIDs(v...)
	}
	if v := i.JudgeID; v != nil {
		m.SetJudgeID(*v)
	}
}

// SetInput applies the change-set in the UpdateProblemInput on the ProblemUpdate builder.
func (c *ProblemUpdate) SetInput(i UpdateProblemInput) *ProblemUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateProblemInput on the ProblemUpdateOne builder.
func (c *ProblemUpdateOne) SetInput(i UpdateProblemInput) *ProblemUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateSubmissionInput represents a mutation input for creating submissions.
type CreateSubmissionInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Status    *submission.Status
	Verdict   *submission.Verdict
	TestCount *int
	ProblemID int
}

// Mutate applies the CreateSubmissionInput on the SubmissionMutation builder.
func (i *CreateSubmissionInput) Mutate(m *SubmissionMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Verdict; v != nil {
		m.SetVerdict(*v)
	}
	if v := i.TestCount; v != nil {
		m.SetTestCount(*v)
	}
	m.SetProblemID(i.ProblemID)
}

// SetInput applies the change-set in the CreateSubmissionInput on the SubmissionCreate builder.
func (c *SubmissionCreate) SetInput(i CreateSubmissionInput) *SubmissionCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateSubmissionInput represents a mutation input for updating submissions.
type UpdateSubmissionInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Status    *submission.Status
	Verdict   *submission.Verdict
	TestCount *int
	ProblemID *int
}

// Mutate applies the UpdateSubmissionInput on the SubmissionMutation builder.
func (i *UpdateSubmissionInput) Mutate(m *SubmissionMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Verdict; v != nil {
		m.SetVerdict(*v)
	}
	if v := i.TestCount; v != nil {
		m.SetTestCount(*v)
	}
	if v := i.ProblemID; v != nil {
		m.SetProblemID(*v)
	}
}

// SetInput applies the change-set in the UpdateSubmissionInput on the SubmissionUpdate builder.
func (c *SubmissionUpdate) SetInput(i UpdateSubmissionInput) *SubmissionUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateSubmissionInput on the SubmissionUpdateOne builder.
func (c *SubmissionUpdateOne) SetInput(i UpdateSubmissionInput) *SubmissionUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
