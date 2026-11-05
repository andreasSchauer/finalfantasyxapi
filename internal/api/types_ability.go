package api

import (
	"fmt"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

type Ability struct {
	ID                 int32                          `json:"id"`
	Name               string                         `json:"name"`
	Version            *int32                         `json:"version,omitempty"`
	Specification      *string                        `json:"specification,omitempty"`
	Type               database.AbilityType           `json:"type"`
	TypedAbility       apiresource.NamedAPIResource   `json:"typed_ability"`
	Rank               *int32                         `json:"rank"`
	AppearsInHelpBar   bool                           `json:"appears_in_help_bar"`
	CanCopycat         bool                           `json:"can_copycat"`
	Monsters           []apiresource.NamedAPIResource `json:"monsters"`
	BattleInteractions []BattleInteraction            `json:"battle_interactions"`
}

type EnemyAbility struct {
	ID                 int32                          `json:"id"`
	Name               string                         `json:"name"`
	Version            *int32                         `json:"version,omitempty"`
	Specification      *string                        `json:"specification,omitempty"`
	UntypedAbility     apiresource.TypedAPIResource   `json:"untyped_ability"`
	Rank               *int32                         `json:"rank"`
	AppearsInHelpBar   bool                           `json:"appears_in_help_bar"`
	CanCopycat         bool                           `json:"can_copycat"`
	Effect             *string                        `json:"effect"`
	Monsters           []apiresource.NamedAPIResource `json:"monsters"`
	BattleInteractions []BattleInteraction            `json:"battle_interactions"`
}

type ItemAbility struct {
	ID                  int32                          `json:"id"`
	Name                string                         `json:"name"`
	UntypedAbility      apiresource.TypedAPIResource   `json:"untyped_ability"`
	Rank                *int32                         `json:"rank"`
	AppearsInHelpBar    bool                           `json:"appears_in_help_bar"`
	CanCopycat          bool                           `json:"can_copycat"`
	Item                apiresource.NamedAPIResource   `json:"item"`
	Description         string                         `json:"description"`
	Effect              string                         `json:"effect"`
	Category            string                         `json:"category"`
	CanUseOutsideBattle bool                           `json:"can_use_outside_battle"`
	Cursor              string                         `json:"cursor"`
	RelatedStats        []apiresource.NamedAPIResource `json:"related_stats"`
	BattleInteractions  []BattleInteraction            `json:"battle_interactions"`
}

type OverdriveAbility struct {
	ID                 int32                          `json:"id"`
	Name               string                         `json:"name"`
	Version            *int32                         `json:"version,omitempty"`
	Specification      *string                        `json:"specification,omitempty"`
	UntypedAbility     apiresource.TypedAPIResource   `json:"untyped_ability"`
	Rank               *int32                         `json:"rank"`
	AppearsInHelpBar   bool                           `json:"appears_in_help_bar"`
	CanCopycat         bool                           `json:"can_copycat"`
	User               apiresource.NamedAPIResource   `json:"user"`
	OverdriveCommand   *apiresource.NamedAPIResource  `json:"overdrive_command"`
	Overdrives         []apiresource.NamedAPIResource `json:"overdrives"`
	RelatedStats       []apiresource.NamedAPIResource `json:"related_stats"`
	BattleInteractions []BattleInteraction            `json:"battle_interactions"`
}

type PlayerAbility struct {
	ID                    int32                                                     `json:"id"`
	Name                  string                                                    `json:"name"`
	Version               *int32                                                    `json:"version,omitempty"`
	Specification         *string                                                   `json:"specification,omitempty"`
	UntypedAbility        apiresource.TypedAPIResource                              `json:"untyped_ability"`
	Rank                  *int32                                                    `json:"rank"`
	AppearsInHelpBar      bool                                                      `json:"appears_in_help_bar"`
	CanCopycat            bool                                                      `json:"can_copycat"`
	Category              string                                                    `json:"category"`
	Description           *string                                                   `json:"description"`
	Effect                string                                                    `json:"effect"`
	Topmenu               *apiresource.NamedAPIResource                             `json:"topmenu"`
	Submenu               *apiresource.NamedAPIResource                             `json:"submenu"`
	OpenSubmenu           *apiresource.NamedAPIResource                             `json:"open_submenu,omitempty"`
	CanUseOutsideBattle   bool                                                      `json:"can_use_outside_battle"`
	MpCost                int32                                                     `json:"mp_cost"`
	Cursor                *string                                                   `json:"cursor"`
	RelatedStats          []apiresource.NamedAPIResource                            `json:"related_stats"`
	LearnedBy             []apiresource.NamedAPIResource                            `json:"learned_by"`
	StandardGridCharacter *apiresource.NamedAPIResource                             `json:"standard_grid_character,omitempty"`
	ExpertGridCharacter   *apiresource.NamedAPIResource                             `json:"expert_grid_character,omitempty"`
	AeonLearnItem         *apiresource.ResourceAmount[apiresource.TypedAPIResource] `json:"aeon_learn_item,omitempty"`
	Monsters              []apiresource.NamedAPIResource                            `json:"monsters"`
	BattleInteractions    []BattleInteraction                                       `json:"battle_interactions"`
}

func (a PlayerAbility) canUseAbility(unitName string) bool {
	for _, class := range a.LearnedBy {
		classLookup, _ := lookup.GetResourceByID(class.ID, lookup.L.CharClassesID)

		if slices.Contains(classLookup.Members, unitName) {
			return true
		}
	}

	return false
}

func (a PlayerAbility) getBattleInteractions() []BattleInteraction {
	return a.BattleInteractions
}

func (a PlayerAbility) String() string {
	return fmt.Sprintf("player ability '%s'", h.NameToString(a.Name, a.Version, nil))
}

type TriggerCommand struct {
	ID                 int32                            `json:"id"`
	Name               string                           `json:"name"`
	Version            *int32                           `json:"version,omitempty"`
	Specification      *string                          `json:"specification,omitempty"`
	UntypedAbility     apiresource.TypedAPIResource     `json:"untyped_ability"`
	Rank               *int32                           `json:"rank"`
	AppearsInHelpBar   bool                             `json:"appears_in_help_bar"`
	CanCopycat         bool                             `json:"can_copycat"`
	Description        string                           `json:"description"`
	Effect             string                           `json:"effect"`
	Topmenu            *apiresource.NamedAPIResource    `json:"topmenu"`
	Cursor             string                           `json:"cursor"`
	RelatedStats       []apiresource.NamedAPIResource   `json:"related_stats"`
	UsedBy             []apiresource.NamedAPIResource   `json:"used_by"`
	MonsterFormations  []apiresource.UnnamedAPIResource `json:"monster_formations"`
	BattleInteractions []BattleInteraction              `json:"battle_interactions"`
}

func (a TriggerCommand) canUseAbility(unitName string) bool {
	for _, class := range a.UsedBy {
		classLookup, _ := lookup.GetResourceByID(class.ID, lookup.L.CharClassesID)

		if slices.Contains(classLookup.Members, unitName) {
			return true
		}
	}

	return false
}

func (a TriggerCommand) getBattleInteractions() []BattleInteraction {
	return a.BattleInteractions
}

func (a TriggerCommand) String() string {
	return fmt.Sprintf("trigger command '%s'", h.NameToString(a.Name, a.Version, nil))
}

type MiscAbility struct {
	ID                 int32                          `json:"id"`
	Name               string                         `json:"name"`
	Version            *int32                         `json:"version,omitempty"`
	Specification      *string                        `json:"specification,omitempty"`
	UntypedAbility     apiresource.TypedAPIResource   `json:"untyped_ability"`
	Rank               *int32                         `json:"rank"`
	AppearsInHelpBar   bool                           `json:"appears_in_help_bar"`
	CanCopycat         bool                           `json:"can_copycat"`
	Description        string                         `json:"description"`
	Effect             string                         `json:"effect"`
	Topmenu            *apiresource.NamedAPIResource  `json:"topmenu"`
	Submenu            *apiresource.NamedAPIResource  `json:"submenu"`
	OpenSubmenu        *apiresource.NamedAPIResource  `json:"open_submenu,omitempty"`
	Cursor             *string                        `json:"cursor"`
	LearnedBy          []apiresource.NamedAPIResource `json:"learned_by"`
	BattleInteractions []BattleInteraction            `json:"battle_interactions"`
}

func (a MiscAbility) canUseAbility(unitName string) bool {
	for _, class := range a.LearnedBy {
		classLookup, _ := lookup.GetResourceByID(class.ID, lookup.L.CharClassesID)

		if slices.Contains(classLookup.Members, unitName) {
			return true
		}
	}

	return false
}

func (a MiscAbility) getBattleInteractions() []BattleInteraction {
	return a.BattleInteractions
}

func (a MiscAbility) String() string {
	return fmt.Sprintf("misc ability '%s'", h.NameToString(a.Name, a.Version, nil))
}
