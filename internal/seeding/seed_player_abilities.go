package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type PlayerAbility struct {
	ID int32
	Ability
	TopmenuID           *int32
	SubmenuID           *int32
	OpenSubmenuID       *int32
	StandardGridCharID  *int32
	ExpertGridCharID    *int32
	Description         *string             `json:"description"`
	Effect              string              `json:"effect"`
	RelatedStats        []string            `json:"related_stats"`
	Category            string              `json:"category"`
	Topmenu             *string             `json:"topmenu"`
	Submenu             *string             `json:"submenu"`
	OpenSubmenu         *string             `json:"open_submenu"`
	LearnedBy           []string            `json:"learned_by"`
	StandardGridPos     *string             `json:"standard_grid_pos"`
	ExpertGridPos       *string             `json:"expert_grid_pos"`
	CanUseOutsideBattle bool                `json:"can_use_outside_battle"`
	AeonLearnItem       *ItemAmount         `json:"aeon_learn_item"`
	MPCost              int32               `json:"mp_cost"`
	Cursor              *string             `json:"cursor"`
	BattleInteractions  []BattleInteraction `json:"battle_interactions"`
}

func (p PlayerAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", p),
		p.Ability.ID,
		h.DerefOrNil(p.Description),
		p.Effect,
		p.Category,
		h.DerefOrNil(p.TopmenuID),
		p.CanUseOutsideBattle,
		p.MPCost,
		h.DerefOrNil(p.Cursor),
		h.DerefOrNil(p.SubmenuID),
		h.DerefOrNil(p.OpenSubmenuID),
		h.DerefOrNil(p.StandardGridCharID),
		h.DerefOrNil(p.ExpertGridCharID),
		h.ObjPtrToID(p.AeonLearnItem),
	}
}

func (p PlayerAbility) ToKeyFields() []any {
	return []any{
		p.Ability.Name,
		h.DerefOrNil(p.Ability.Version),
	}
}

func (p PlayerAbility) GetID() int32 {
	return p.ID
}

func (p PlayerAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        p.Name,
		Version:     p.Version,
		AbilityType: string(database.AbilityTypePlayerAbility),
	}
}

func (p PlayerAbility) Error() string {
	return fmt.Sprintf("player ability '%s'", h.NameToString(p.Name, p.Version, p.Specification))
}

func (p PlayerAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            p.ID,
		Name:          p.Name,
		Version:       p.Version,
		Specification: p.Specification,
	}
}

func (p PlayerAbility) GetItemAmount() ItemAmount {
	itemAmtPtr := p.AeonLearnItem

	if itemAmtPtr == nil {
		return ItemAmount{}
	}

	return *itemAmtPtr
}

func (l *Lookup) seedPlayerAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/player_abilities.json"

	var playerAbilities []PlayerAbility

	err := loadJSONFile(string(srcPath), &playerAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, playerAbility := range playerAbilities {
			var err error
			playerAbility.Type = database.AbilityTypePlayerAbility

			playerAbility.Ability, err = seedObjAssignID(qtx, playerAbility.Ability, l.seedAbility)
			if err != nil {
				return h.NewErr(playerAbility.Error(), err)
			}

			dbPlayerAbility, err := qtx.CreatePlayerAbility(context.Background(), database.CreatePlayerAbilityParams{
				DataHash:            generateDataHash(playerAbility),
				AbilityID:           playerAbility.Ability.ID,
				Description:         h.GetNullString(playerAbility.Description),
				Effect:              playerAbility.Effect,
				Category:            database.PlayerAbilityCategory(playerAbility.Category),
				CanUseOutsideBattle: playerAbility.CanUseOutsideBattle,
				MpCost:              playerAbility.MPCost,
				Cursor:              database.ToNullTargetType(playerAbility.Cursor),
			})
			if err != nil {
				return h.NewErr(playerAbility.Error(), err, "couldn't create player ability")
			}

			playerAbility.ID = dbPlayerAbility.ID
			key := Key(playerAbility)
			l.PlayerAbilities[key] = playerAbility
			l.PlayerAbilitiesID[playerAbility.ID] = playerAbility
		}
		return nil
	})
}

func (l *Lookup) seedPlayerAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/player_abilities.json"

	var playerAbilities []PlayerAbility

	err := loadJSONFile(string(srcPath), &playerAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range playerAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			playerAbility, err := GetResource(abilityRef.Untyped(), l.PlayerAbilities)
			if err != nil {
				return err
			}

			err = l.seedPlayerAbilityFKs(qtx, playerAbility)
			if err != nil {
				return h.NewErr(playerAbility.Error(), err)
			}

			err = l.seedPlayerAbilityRelatedStats(qtx, playerAbility)
			if err != nil {
				return h.NewErr(playerAbility.Error(), err)
			}

			err = l.seedPlayerAbilityLearnedBy(qtx, playerAbility)
			if err != nil {
				return h.NewErr(playerAbility.Error(), err)
			}

			l.currentAbility = playerAbility.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, playerAbility.BattleInteractions)
			if err != nil {
				return h.NewErr(playerAbility.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedPlayerAbilityFKs(qtx *database.Queries, ability PlayerAbility) error {
	var err error

	ability.TopmenuID, err = assignFKPtr(ability.Topmenu, l.Topmenus)
	if err != nil {
		return err
	}

	ability.SubmenuID, err = assignFKPtr(ability.Submenu, l.Submenus)
	if err != nil {
		return err
	}

	ability.OpenSubmenuID, err = assignFKPtr(ability.OpenSubmenu, l.Submenus)
	if err != nil {
		return err
	}

	ability.StandardGridCharID, err = assignFKPtr(ability.StandardGridPos, l.Characters)
	if err != nil {
		return err
	}

	ability.ExpertGridCharID, err = assignFKPtr(ability.ExpertGridPos, l.Characters)
	if err != nil {
		return err
	}

	ability.AeonLearnItem, err = seedObjPtrAssignFK(qtx, ability.AeonLearnItem, l.seedItemAmount)
	if err != nil {
		return err
	}

	err = qtx.UpdatePlayerAbility(context.Background(), database.UpdatePlayerAbilityParams{
		DataHash:           generateDataHash(ability),
		TopmenuID:          h.GetNullInt32(ability.TopmenuID),
		SubmenuID:          h.GetNullInt32(ability.SubmenuID),
		OpenSubmenuID:      h.GetNullInt32(ability.OpenSubmenuID),
		StandardGridCharID: h.GetNullInt32(ability.StandardGridCharID),
		ExpertGridCharID:   h.GetNullInt32(ability.ExpertGridCharID),
		AeonLearnItemID:    h.ObjPtrToNullInt32ID(ability.AeonLearnItem),
		ID:                 ability.ID,
	})
	if err != nil {
		return h.NewErr("", err, "couldn't update player ability")
	}

	return nil
}

func (l *Lookup) seedPlayerAbilityRelatedStats(qtx *database.Queries, ability PlayerAbility) error {
	for _, jsonStat := range ability.RelatedStats {
		junction, err := createJunction(ability, jsonStat, l.Stats)
		if err != nil {
			return err
		}

		err = qtx.CreatePlayerAbilitiesRelatedStatsJunction(context.Background(), database.CreatePlayerAbilitiesRelatedStatsJunctionParams{
			DataHash:        generateDataHash(junction),
			PlayerAbilityID: junction.ParentID,
			StatID:          junction.ChildID,
		})
		if err != nil {
			return h.NewErr(jsonStat, err, "couldn't junction related stat")
		}
	}

	return nil
}

func (l *Lookup) seedPlayerAbilityLearnedBy(qtx *database.Queries, ability PlayerAbility) error {
	for _, charClass := range ability.LearnedBy {
		junction, err := createJunction(ability, charClass, l.CharClasses)
		if err != nil {
			return err
		}

		err = qtx.CreatePlayerAbilitiesLearnedByJunction(context.Background(), database.CreatePlayerAbilitiesLearnedByJunctionParams{
			DataHash:         generateDataHash(junction),
			PlayerAbilityID:  junction.ParentID,
			CharacterClassID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(charClass, err, "couldn't junction 'learned by' class")
		}
	}

	return nil
}

func (l *Lookup) loop5SeedPlayerAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractPlayerAbilities()
	if err != nil {
		return err
	}

	params := database.CreatePlayerAbilityBulkParams{
		DataHash:            make([]string, len(abilities)),
		AbilityID:           make([]int32, len(abilities)),
		Description:         make([]sql.NullString, len(abilities)),
		Effect:              make([]string, len(abilities)),
		Category:            make([]database.PlayerAbilityCategory, len(abilities)),
		CanUseOutsideBattle: make([]bool, len(abilities)),
		MpCost:              make([]int32, len(abilities)),
		Cursor:              make([]database.NullTargetType, len(abilities)),
		TopmenuID:           make([]sql.NullInt32, len(abilities)),
		SubmenuID:           make([]sql.NullInt32, len(abilities)),
		OpenSubmenuID:       make([]sql.NullInt32, len(abilities)),
		StdGridCharID:       make([]sql.NullInt32, len(abilities)),
		ExpGridCharID:       make([]sql.NullInt32, len(abilities)),
		AeonLearnItemID:     make([]sql.NullInt32, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = generateDataHash(a)
		params.AbilityID[i] = a.Ability.ID
		params.Description[i] = h.GetNullString(a.Description)
		params.Effect[i] = a.Effect
		params.Category[i] = database.PlayerAbilityCategory(a.Category)
		params.CanUseOutsideBattle[i] = a.CanUseOutsideBattle
		params.MpCost[i] = a.MPCost
		params.Cursor[i] = database.ToNullTargetType(a.Cursor)
		params.TopmenuID[i] = h.GetNullInt32(a.TopmenuID)
		params.SubmenuID[i] = h.GetNullInt32(a.SubmenuID)
		params.OpenSubmenuID[i] = h.GetNullInt32(a.OpenSubmenuID)
		params.StdGridCharID[i] = h.GetNullInt32(a.StandardGridCharID)
		params.ExpGridCharID[i] = h.GetNullInt32(a.ExpertGridCharID)
		params.AeonLearnItemID[i] = h.ObjPtrToNullInt32ID(a.AeonLearnItem)
	}

	dbRows, err := qtx.CreatePlayerAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create player abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		l.json.playerAbilities[i].ID = row.ID
		key := Key(abilities[i])
		l.PlayerAbilities[key] = abilities[i]
		l.PlayerAbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractPlayerAbilities() ([]PlayerAbility, error) {
	abilities := []PlayerAbility{}
	var err error

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		ability.Ability.ID, err = l.getHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		ability.TopmenuID, err = assignFKPtr(ability.Topmenu, l.Topmenus)
		if err != nil {
			return nil, err
		}

		ability.SubmenuID, err = assignFKPtr(ability.Submenu, l.Submenus)
		if err != nil {
			return nil, err
		}

		ability.OpenSubmenuID, err = assignFKPtr(ability.OpenSubmenu, l.Submenus)
		if err != nil {
			return nil, err
		}

		ability.StandardGridCharID, err = assignFKPtr(ability.StandardGridPos, l.Characters)
		if err != nil {
			return nil, err
		}

		ability.ExpertGridCharID, err = assignFKPtr(ability.ExpertGridPos, l.Characters)
		if err != nil {
			return nil, err
		}

		if ability.AeonLearnItem != nil {
			ability.AeonLearnItem.ID, err = l.getHashID(ability.AeonLearnItem)
			if err != nil {
				return nil, err
			}
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, l.Hashes), nil
}

func (l *Lookup) completePlayerAbilities() error {
	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		err := l.completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		l.PlayerAbilities[Key(ability)] = *ability
		l.PlayerAbilitiesID[ability.ID] = *ability
	}

	return nil
}