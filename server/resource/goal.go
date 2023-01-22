package resource

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/paidoff/db"
)

type GoalResource struct {
	goalStore db.GoalStore
}

func NewGoalResource(goalStore db.GoalStore) *GoalResource {
	return &GoalResource{
		goalStore: goalStore,
	}
}

func (this *GoalResource) GetGoal(req Request) *Response {
	return nil
}

func (this *GoalResource) GetGoals(req Request) *Response {
	return nil
}

func (this *GoalResource) UpdateGoal(req Request) *Response {
	return nil
}

func (this *GoalResource) CreateGoal(req Request) *Response {
	accountId := mustGetAccountId(req)

	var goalCreateRequest GoalCreateRequest

	err := json.Unmarshal(req.Body, &goalCreateRequest)

	if err != nil {
		//return 500
	}

	//do validations

	timeNow := int64(123)

	newGoal := db.Goal{
		Id:            "gen-uuid",
		AccountId:     accountId,
		CategoryId:    goalCreateRequest.CategoryId,
		Month:         goalCreateRequest.Month,
		Year:          goalCreateRequest.Year,
		Name:          goalCreateRequest.Name,
		CurrentAmount: goalCreateRequest.CurrentAmount,
		TargetAmount:  goalCreateRequest.TargetAmount,
		GoalType:      goalCreateRequest.GoalType,
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
	}

	err = this.goalStore.Insert(newGoal)

	if err != nil {
		//return 500
	}

	goalResponse := this.toGoalResponse(newGoal)
	responseBody, err := json.Marshal(&goalResponse)

	if err != nil {
		//return 500
	}

	return &Response{
		Body:   responseBody,
		Status: http.StatusOK,
	}
}

func (this *GoalResource) DeleteGoal(req Request) *Response {
	return nil
}

func (this *GoalResource) toGoalResponse(goal db.Goal) *GoalResponse {
	return nil
}
