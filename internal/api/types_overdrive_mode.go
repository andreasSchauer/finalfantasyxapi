package api


type OverdriveMode struct {
	ID          int32            `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Effect      string           `json:"effect"`
	Type        NamedAPIResource `json:"type"`
	FillRate    *float32         `json:"fill_rate,omitempty"`
	Actions     []ActionAmount   `json:"actions"`
}


type ActionAmount struct {
	User   NamedAPIResource `json:"user"`
	Amount int32            `json:"amount"`
}

func convertActionAmount(res NamedAPIResource, amount int32) ActionAmount {
	return ActionAmount{
		User:   res,
		Amount: amount,
	}
}

func (a ActionAmount) GetAPIResource() APIResource {
	return a.User
}

func (a ActionAmount) GetName() string {
	return a.User.Name
}

func (a ActionAmount) GetVersion() *int32 {
	return nil
}

func (a ActionAmount) GetVal() int32 {
	return a.Amount
}


type ModeAmount struct {
	OverdriveMode	NamedAPIResource	`json:"overdrive_mode"`
	Amount			int32				`json:"amount"`
}

func convertModeAmount(res NamedAPIResource, amount int32) ModeAmount {
	return ModeAmount{
		OverdriveMode:  res,
		Amount: 		amount,
	}
}

func (m ModeAmount) GetAPIResource() APIResource {
	return m.OverdriveMode
}

func (m ModeAmount) GetName() string {
	return m.OverdriveMode.Name
}

func (m ModeAmount) GetVersion() *int32 {
	return nil
}

func (m ModeAmount) GetVal() int32 {
	return m.Amount
}