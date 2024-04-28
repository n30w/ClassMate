package domain

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/n30w/Darkspace/internal/models"
)

type AssignmentStore interface {
	GetAssignmentById(assignmentid string) (*models.Assignment, error)
	InsertAssignment(assignment *models.Assignment) error
	DeleteAssignment(assignment *models.Assignment) error
	SubmitAssignment(assignment *models.Assignment) (*models.Assignment, error)
	ChangeAssignment(assignment *models.Assignment, updatedfield string, action string) (*models.Assignment, error)
}

type AssignmentService struct {
	store AssignmentStore
}

func NewAssignmentService(a AssignmentStore) *AssignmentService { return &AssignmentService{store: a} }

func (as *AssignmentService) ReadAssignment(assignmentid string) (*models.Assignment, error) {
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

	assignment, err := as.store.GetAssignmentById(assignmentid)
	if err != nil {
		return nil, err
	}
	if action == "submit" {
		if _, ok := updatedfield.(bool); !ok {
			return nil, fmt.Errorf("updated field is not of type bool, it is of type %T", updatedfield)
		}
		assignment, err := as.store.SubmitAssignment(assignment)
		if err != nil {
			return nil, err
		}
		return assignment, nil
	} else if action == "body" || action == "title" || action == "duedate" {
		if _, ok := updatedfield.(string); !ok {
			return nil, fmt.Errorf("updated field is not of type string, it is of type %T", updatedfield)
		}
		assignment, err := as.store.ChangeAssignment(assignment, updatedfield.(string), action)
		if err != nil {
			return nil, err
		}
		return assignment, nil
	} else {
		return nil, fmt.Errorf("%s is an invalid action", action)
	}
}

func (as *AssignmentService) DeleteAssignment(assignmentid string) error {
	assignment, err := as.store.GetAssignmentById(assignmentid)
	if err != nil {
		return err
	}
	err = as.store.DeleteAssignment(assignment)
	if err != nil {
		return err
	}
	return nil
}
