package domain

import (
	"fmt"

	// "github.com/google/uuid"
	"github.com/n30w/Darkspace/internal/models"
)

type AssignmentStore interface {
	GetAssignmentById(assignmentid string) (*models.Assignment, error)
	GetAssignmentsByCourse(courseid string) ([]string, error)
	InsertIntoCourseAssignments(a *models.Assignment) (
		*models.Assignment,
		error,
	)
	InsertAssignmentIntoUser(a *models.Assignment) (*models.Assignment, error)
	InsertAssignment(assignment *models.Assignment) (*models.Assignment, error)
	DeleteAssignmentByID(assignmentid string) error
	ChangeAssignment(
		assignment *models.Assignment,
		updatedfield string,
		action string,
	) (*models.Assignment, error)
}

type AssignmentService struct {
	store AssignmentStore
}

func NewAssignmentService(a AssignmentStore) *AssignmentService { return &AssignmentService{store: a} }

// ReadAssignment uses an Assignment's ID to retrieve it from
// the database. Options can also be passed in that specify
// what types of data transformations can be done, for example
// changing the date to a readable format.
func (as *AssignmentService) ReadAssignment(
	assignmentId string,
	opts ...func(assignment *models.Assignment) error,
) (
	*models.Assignment,
	error,
) {
	assignment, err := as.store.GetAssignmentById(assignmentId)
	if err != nil {
		return nil, err
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			err := opt(assignment)
			if err != nil {
				return nil, fmt.Errorf("option transform error: ", err)
			}
		}
	}

	return assignment, nil
}

// RetrieveAssignments retrieves an assignment using a specific
// Course ID. It returns a slice of all the assignments in a course.
func (as *AssignmentService) RetrieveAssignments(courseid string) (
	[]string,
	error,
) {
	assignmentIds, err := as.store.GetAssignmentsByCourse(courseid)
	if err != nil {
		return nil, err
	}
	return assignmentIds, nil
}

func (as *AssignmentService) CreateAssignment(assignment *models.Assignment) (
	*models.Assignment,
	error,
) {
	assignment, err := as.store.InsertAssignment(assignment)
	if err != nil {
		return nil, err
	}

	assignment, err = as.store.InsertIntoCourseAssignments(assignment)
	if err != nil {
		return nil, err
	}

	assignment, err = as.store.InsertAssignmentIntoUser(assignment)
	if err != nil {
		return nil, err
	}

	return assignment, nil
}

func (as *AssignmentService) UpdateAssignment(
	assignmentid string,
	updatedfield interface{},
	action string,
) (*models.Assignment, error) {

	assignment, err := as.store.GetAssignmentById(assignmentid)
	if err != nil {
		return nil, err
	}

	if action == "body" || action == "title" || action == "duedate" {
		if _, ok := updatedfield.(string); !ok {
			return nil, fmt.Errorf(
				"updated field is not of type string, it is of type %T",
				updatedfield,
			)
		}
		assignment, err := as.store.ChangeAssignment(
			assignment,
			updatedfield.(string),
			action,
		)
		if err != nil {
			return nil, err
		}
		return assignment, nil
	} else {
		return nil, fmt.Errorf("%s is an invalid action", action)
	}
}

func (as *AssignmentService) DeleteAssignment(assignmentid string) error {
	err := as.store.DeleteAssignmentByID(assignmentid)
	if err != nil {
		return err
	}
	return nil
}
