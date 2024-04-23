package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/n30w/Darkspace/internal/models"
)

type AssignmentStore interface {
	GetAssignmentById(assignmentid string) (*models.Assignment, error)
	InsertAssignment(assignment *models.Assignment) error
	DeleteAssignment(assignment *models.Assignment) error
	ChangeAssignmentDueDate(assignment *models.Assignment, duedate time.Time) (*models.Assignment, error)
	ChangeAssignmentTitle(assignment *models.Assignment, title string) (*models.Assignment, error)
	ChangeAssignmentBody(assignment *models.Assignment, body string) (*models.Assignment, error)
}

type AssignmentService struct {
	store AssignmentStore
}

func NewAssignmentService(a AssignmentStore) *AssignmentService { return &AssignmentService{store: a} }

func (as *AssignmentService) ValidateID(assignmentid string) bool {
	return true
}

func (as *AssignmentService) ReadAssignment(assignmentid string) (*models.Assignment, error) {
	if !as.ValidateID(assignmentid) {
		return nil, fmt.Errorf("invalid assignment ID: %s", assignmentid)
	}
	assignment, err := as.store.GetAssignmentById(assignmentid)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

func (as *AssignmentService) CreateAssignment(assignment *models.Assignment) (*models.Assignment, error) {
	assignment.ID = uuid.New().String()
	err := as.store.InsertAssignment(assignment)
	if err != nil {
		return nil, err
	}

	return assignment, nil
}

func (as *AssignmentService) UpdateAssignment(assignmentid string, updatedfield interface{}, action string) (*models.Assignment, error) {
	if !as.ValidateID(assignmentid) {
		return nil, fmt.Errorf("invalid assignment ID: %s", assignmentid)
	}
	assignment, err := as.store.GetAssignmentById(assignmentid)
	if err != nil {
		return nil, err
	}
	if action == "title" {
		if _, ok := updatedfield.(string); !ok {
			return nil, fmt.Errorf("updated field is not of type string, it is of type %T", updatedfield)
		}
		assignment, err := as.store.ChangeAssignmentTitle(assignment, updatedfield.(string))
		if err != nil {
			return nil, err
		}
		return assignment, nil
	} else if action == "body" {
		if _, ok := updatedfield.(string); !ok {
			return nil, fmt.Errorf("updated field is not of type string, it is of type %T", updatedfield)
		}
		assignment, err := as.store.ChangeAssignmentBody(assignment, updatedfield.(string))
		if err != nil {
			return nil, err
		}
		return assignment, nil
	} else if action == "duedate" {
		if _, ok := updatedfield.(time.Time); !ok {
			return nil, fmt.Errorf("updated field is not of type string, it is of type %T", updatedfield)
		}
		assignment, err := as.store.ChangeAssignmentDueDate(assignment, updatedfield.(time.Time))
		if err != nil {
			return nil, err
		}
		return assignment, nil
	} else {
		return nil, fmt.Errorf("%s is an invalid action", action)
	}
}
