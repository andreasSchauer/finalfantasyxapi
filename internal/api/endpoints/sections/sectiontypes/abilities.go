package sectiontypes

type AbilitySimple struct {
	ID                 int32                     `json:"id"`
	URL                string                    `json:"url"`
	Name               string                    `json:"name"`
	Version            *int32                    `json:"version,omitempty"`
	Specification      *string                   `json:"specification,omitempty"`
	Type               *string                   `json:"type,omitempty"`
	Rank               *int32                    `json:"rank"`
	BattleInteractions []BattleInteractionSimple `json:"battle_interactions"`
}

func (a AbilitySimple) GetURL() string {
	return a.URL
}

type PlayerAbilitySimple struct {
	ID                 int32                     `json:"id"`
	URL                string                    `json:"url"`
	Name               string                    `json:"name"`
	Version            *int32                    `json:"version,omitempty"`
	Specification      *string                   `json:"specification,omitempty"`
	Type               *string                   `json:"type,omitempty"`
	Rank               *int32                    `json:"rank"`
	MpCost             int32                     `json:"mp_cost"`
	BattleInteractions []BattleInteractionSimple `json:"battle_interactions"`
}

func (a PlayerAbilitySimple) GetURL() string {
	return a.URL
}
