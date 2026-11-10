// Package apitypes ajsidoaüdsj
package apitypes

import (
	"fmt"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

type UserAbility interface {
	CanUseAbility(string) bool
	GetBattleInteractions() []BattleInteraction
	fmt.Stringer
}

type Ability struct {
	ID                 int32                `json:"id"`
	Name               string               `json:"name"`
	Version            *int32               `json:"version,omitempty"`
	Specification      *string              `json:"specification,omitempty"`
	Type               database.AbilityType `json:"type"`
	TypedAbility       apiresource.Named    `json:"typed_ability"`
	Rank               *int32               `json:"rank"`
	AppearsInHelpBar   bool                 `json:"appears_in_help_bar"`
	CanCopycat         bool                 `json:"can_copycat"`
	Monsters           []apiresource.Named  `json:"monsters"`
	BattleInteractions []BattleInteraction  `json:"battle_interactions"`
}

type EnemyAbility struct {
	ID                 int32               `json:"id"`
	Name               string              `json:"name"`
	Version            *int32              `json:"version,omitempty"`
	Specification      *string             `json:"specification,omitempty"`
	UntypedAbility     apiresource.Typed   `json:"untyped_ability"`
	Rank               *int32              `json:"rank"`
	AppearsInHelpBar   bool                `json:"appears_in_help_bar"`
	CanCopycat         bool                `json:"can_copycat"`
	Effect             *string             `json:"effect"`
	Monsters           []apiresource.Named `json:"monsters"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

type ItemAbility struct {
	ID                  int32               `json:"id"`
	Name                string              `json:"name"`
	UntypedAbility      apiresource.Typed   `json:"untyped_ability"`
	Rank                *int32              `json:"rank"`
	AppearsInHelpBar    bool                `json:"appears_in_help_bar"`
	CanCopycat          bool                `json:"can_copycat"`
	Item                apiresource.Named   `json:"item"`
	Description         string              `json:"description"`
	Effect              string              `json:"effect"`
	Category            string              `json:"category"`
	CanUseOutsideBattle bool                `json:"can_use_outside_battle"`
	Cursor              string              `json:"cursor"`
	RelatedStats        []apiresource.Named `json:"related_stats"`
	BattleInteractions  []BattleInteraction `json:"battle_interactions"`
}

type OverdriveAbility struct {
	ID                 int32               `json:"id"`
	Name               string              `json:"name"`
	Version            *int32              `json:"version,omitempty"`
	Specification      *string             `json:"specification,omitempty"`
	UntypedAbility     apiresource.Typed   `json:"untyped_ability"`
	Rank               *int32              `json:"rank"`
	AppearsInHelpBar   bool                `json:"appears_in_help_bar"`
	CanCopycat         bool                `json:"can_copycat"`
	User               apiresource.Named   `json:"user"`
	OverdriveCommand   *apiresource.Named  `json:"overdrive_command"`
	Overdrives         []apiresource.Named `json:"overdrives"`
	RelatedStats       []apiresource.Named `json:"related_stats"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

type PlayerAbility struct {
	ID                    int32                                          `json:"id"`
	Name                  string                                         `json:"name"`
	Version               *int32                                         `json:"version,omitempty"`
	Specification         *string                                        `json:"specification,omitempty"`
	UntypedAbility        apiresource.Typed                              `json:"untyped_ability"`
	Rank                  *int32                                         `json:"rank"`
	AppearsInHelpBar      bool                                           `json:"appears_in_help_bar"`
	CanCopycat            bool                                           `json:"can_copycat"`
	Category              string                                         `json:"category"`
	Description           *string                                        `json:"description"`
	Effect                string                                         `json:"effect"`
	Topmenu               *apiresource.Named                             `json:"topmenu"`
	Submenu               *apiresource.Named                             `json:"submenu"`
	OpenSubmenu           *apiresource.Named                             `json:"open_submenu,omitempty"`
	CanUseOutsideBattle   bool                                           `json:"can_use_outside_battle"`
	MpCost                int32                                          `json:"mp_cost"`
	Cursor                *string                                        `json:"cursor"`
	RelatedStats          []apiresource.Named                            `json:"related_stats"`
	LearnedBy             []apiresource.Named                            `json:"learned_by"`
	StandardGridCharacter *apiresource.Named                             `json:"standard_grid_character,omitempty"`
	ExpertGridCharacter   *apiresource.Named                             `json:"expert_grid_character,omitempty"`
	AeonLearnItem         *apiresource.ResourceAmount[apiresource.Typed] `json:"aeon_learn_item,omitempty"`
	Monsters              []apiresource.Named                            `json:"monsters"`
	BattleInteractions    []BattleInteraction                            `json:"battle_interactions"`
}

func (a PlayerAbility) CanUseAbility(unitName string) bool {
	for _, class := range a.LearnedBy {
		classLookup, _ := lookup.GetResourceByID(class.ID, lookup.L.CharClassesID)

		if slices.Contains(classLookup.Members, unitName) {
			return true
		}
	}

	return false
}

func (a PlayerAbility) GetBattleInteractions() []BattleInteraction {
	return a.BattleInteractions
}

func (a PlayerAbility) String() string {
	return fmt.Sprintf("player ability '%s'", h.NameToString(a.Name, a.Version, nil))
}

type TriggerCommand struct {
	ID                 int32                 `json:"id"`
	Name               string                `json:"name"`
	Version            *int32                `json:"version,omitempty"`
	Specification      *string               `json:"specification,omitempty"`
	UntypedAbility     apiresource.Typed     `json:"untyped_ability"`
	Rank               *int32                `json:"rank"`
	AppearsInHelpBar   bool                  `json:"appears_in_help_bar"`
	CanCopycat         bool                  `json:"can_copycat"`
	Description        string                `json:"description"`
	Effect             string                `json:"effect"`
	Topmenu            *apiresource.Named    `json:"topmenu"`
	Cursor             string                `json:"cursor"`
	RelatedStats       []apiresource.Named   `json:"related_stats"`
	UsedBy             []apiresource.Named   `json:"used_by"`
	MonsterFormations  []apiresource.Unnamed `json:"monster_formations"`
	BattleInteractions []BattleInteraction   `json:"battle_interactions"`
}

func (a TriggerCommand) CanUseAbility(unitName string) bool {
	for _, class := range a.UsedBy {
		classLookup, _ := lookup.GetResourceByID(class.ID, lookup.L.CharClassesID)

		if slices.Contains(classLookup.Members, unitName) {
			return true
		}
	}

	return false
}

func (a TriggerCommand) GetBattleInteractions() []BattleInteraction {
	return a.BattleInteractions
}

func (a TriggerCommand) String() string {
	return fmt.Sprintf("trigger command '%s'", h.NameToString(a.Name, a.Version, nil))
}

type MiscAbility struct {
	ID                 int32               `json:"id"`
	Name               string              `json:"name"`
	Version            *int32              `json:"version,omitempty"`
	Specification      *string             `json:"specification,omitempty"`
	UntypedAbility     apiresource.Typed   `json:"untyped_ability"`
	Rank               *int32              `json:"rank"`
	AppearsInHelpBar   bool                `json:"appears_in_help_bar"`
	CanCopycat         bool                `json:"can_copycat"`
	Description        string              `json:"description"`
	Effect             string              `json:"effect"`
	Topmenu            *apiresource.Named  `json:"topmenu"`
	Submenu            *apiresource.Named  `json:"submenu"`
	OpenSubmenu        *apiresource.Named  `json:"open_submenu,omitempty"`
	Cursor             *string             `json:"cursor"`
	LearnedBy          []apiresource.Named `json:"learned_by"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

func (a MiscAbility) CanUseAbility(unitName string) bool {
	for _, class := range a.LearnedBy {
		classLookup, _ := lookup.GetResourceByID(class.ID, lookup.L.CharClassesID)

		if slices.Contains(classLookup.Members, unitName) {
			return true
		}
	}

	return false
}

func (a MiscAbility) GetBattleInteractions() []BattleInteraction {
	return a.BattleInteractions
}

func (a MiscAbility) String() string {
	return fmt.Sprintf("misc ability '%s'", h.NameToString(a.Name, a.Version, nil))
}
