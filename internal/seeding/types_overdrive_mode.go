package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type OverdriveMode struct {
	ID             int32
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Effect         string          `json:"effect"`
	Type           string          `json:"type"`
	FillRate       *float32        `json:"fill_rate"`
	ActionsToLearn []ActionToLearn `json:"actions_to_learn"`
}

func (o OverdriveMode) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", o),
		o.Name,
		o.Description,
		o.Effect,
		o.Type,
		h.DerefOrNil(o.FillRate),
	}
}

func (o OverdriveMode) ToKeyFields() []any {
	return []any{
		o.Name,
	}
}

func (o OverdriveMode) GetID() int32 {
	return o.ID
}

func (o OverdriveMode) GetResParamsNamed() ResParamsNamed {
	return ResParamsNamed{
		ID:   o.ID,
		Name: o.Name,
	}
}

func (o OverdriveMode) Error() string {
	return fmt.Sprintf("overdrive mode %s", o.Name)
}

type ActionToLearn struct {
	ID     int32
	UserID int32
	User   string `json:"user"`
	Amount int32  `json:"amount"`
}

func (a ActionToLearn) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.UserID,
		a.Amount,
	}
}

func (a ActionToLearn) GetID() int32 {
	return a.ID
}

func (a *ActionToLearn) SetID(id int32) {
	a.ID = id
}

func (a ActionToLearn) GetName() string {
	return a.User
}

func (a ActionToLearn) GetVersion() *int32 {
	return nil
}

func (a ActionToLearn) GetVal() int32 {
	return a.Amount
}

func (a ActionToLearn) Error() string {
	return fmt.Sprintf("action to learn with user: %s, amount: %d", a.User, a.Amount)
}
