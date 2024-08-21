// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateJudge implements createJudge operation.
//
// Creates a new Judge and persists it to storage.
//
// POST /judges
func (UnimplementedHandler) CreateJudge(ctx context.Context, req *CreateJudgeReq) (r CreateJudgeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateProblem implements createProblem operation.
//
// Creates a new Problem and persists it to storage.
//
// POST /problems
func (UnimplementedHandler) CreateProblem(ctx context.Context, req *CreateProblemReq) (r CreateProblemRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateSubmission implements createSubmission operation.
//
// Creates a new Submission and persists it to storage.
//
// POST /submissions
func (UnimplementedHandler) CreateSubmission(ctx context.Context, req *CreateSubmissionReq) (r CreateSubmissionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteJudge implements deleteJudge operation.
//
// Deletes the Judge with the requested ID.
//
// DELETE /judges/{id}
func (UnimplementedHandler) DeleteJudge(ctx context.Context, params DeleteJudgeParams) (r DeleteJudgeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteProblem implements deleteProblem operation.
//
// Deletes the Problem with the requested ID.
//
// DELETE /problems/{id}
func (UnimplementedHandler) DeleteProblem(ctx context.Context, params DeleteProblemParams) (r DeleteProblemRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteSubmission implements deleteSubmission operation.
//
// Deletes the Submission with the requested ID.
//
// DELETE /submissions/{id}
func (UnimplementedHandler) DeleteSubmission(ctx context.Context, params DeleteSubmissionParams) (r DeleteSubmissionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListJudge implements listJudge operation.
//
// List Judges.
//
// GET /judges
func (UnimplementedHandler) ListJudge(ctx context.Context, params ListJudgeParams) (r ListJudgeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListJudgeProblems implements listJudgeProblems operation.
//
// List attached Problems.
//
// GET /judges/{id}/problems
func (UnimplementedHandler) ListJudgeProblems(ctx context.Context, params ListJudgeProblemsParams) (r ListJudgeProblemsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListProblem implements listProblem operation.
//
// List Problems.
//
// GET /problems
func (UnimplementedHandler) ListProblem(ctx context.Context, params ListProblemParams) (r ListProblemRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListProblemSubmissions implements listProblemSubmissions operation.
//
// List attached Submissions.
//
// GET /problems/{id}/submissions
func (UnimplementedHandler) ListProblemSubmissions(ctx context.Context, params ListProblemSubmissionsParams) (r ListProblemSubmissionsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListSubmission implements listSubmission operation.
//
// List Submissions.
//
// GET /submissions
func (UnimplementedHandler) ListSubmission(ctx context.Context, params ListSubmissionParams) (r ListSubmissionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadJudge implements readJudge operation.
//
// Finds the Judge with the requested ID and returns it.
//
// GET /judges/{id}
func (UnimplementedHandler) ReadJudge(ctx context.Context, params ReadJudgeParams) (r ReadJudgeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadProblem implements readProblem operation.
//
// Finds the Problem with the requested ID and returns it.
//
// GET /problems/{id}
func (UnimplementedHandler) ReadProblem(ctx context.Context, params ReadProblemParams) (r ReadProblemRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadProblemJudge implements readProblemJudge operation.
//
// Find the attached Judge of the Problem with the given ID.
//
// GET /problems/{id}/judge
func (UnimplementedHandler) ReadProblemJudge(ctx context.Context, params ReadProblemJudgeParams) (r ReadProblemJudgeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadSubmission implements readSubmission operation.
//
// Finds the Submission with the requested ID and returns it.
//
// GET /submissions/{id}
func (UnimplementedHandler) ReadSubmission(ctx context.Context, params ReadSubmissionParams) (r ReadSubmissionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadSubmissionProblem implements readSubmissionProblem operation.
//
// Find the attached Problem of the Submission with the given ID.
//
// GET /submissions/{id}/problem
func (UnimplementedHandler) ReadSubmissionProblem(ctx context.Context, params ReadSubmissionProblemParams) (r ReadSubmissionProblemRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateJudge implements updateJudge operation.
//
// Updates a Judge and persists changes to storage.
//
// PATCH /judges/{id}
func (UnimplementedHandler) UpdateJudge(ctx context.Context, req *UpdateJudgeReq, params UpdateJudgeParams) (r UpdateJudgeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateProblem implements updateProblem operation.
//
// Updates a Problem and persists changes to storage.
//
// PATCH /problems/{id}
func (UnimplementedHandler) UpdateProblem(ctx context.Context, req *UpdateProblemReq, params UpdateProblemParams) (r UpdateProblemRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateSubmission implements updateSubmission operation.
//
// Updates a Submission and persists changes to storage.
//
// PATCH /submissions/{id}
func (UnimplementedHandler) UpdateSubmission(ctx context.Context, req *UpdateSubmissionReq, params UpdateSubmissionParams) (r UpdateSubmissionRes, _ error) {
	return r, ht.ErrNotImplemented
}
