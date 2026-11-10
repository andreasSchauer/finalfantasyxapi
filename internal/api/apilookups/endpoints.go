// Package apilookups ajiodaüdi
package apilookups

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

var E = &Endpoints{}

type Endpoints struct {
	Locations          handlerinput.Main[seedtypes.Location, apitypes.Location, apiresource.Named]
	Sublocations       handlerinput.Main[seedtypes.Sublocation, apitypes.Sublocation, apiresource.Named]
	Areas              handlerinput.Main[seedtypes.Area, apitypes.Area, apiresource.Area]
	MonsterFormations  handlerinput.Main[seedtypes.MonsterFormation, apitypes.MonsterFormation, apiresource.Unnamed]
	Shops              handlerinput.Main[seedtypes.Shop, apitypes.Shop, apiresource.Unnamed]
	Treasures          handlerinput.Main[seedtypes.Treasure, apitypes.Treasure, apiresource.Unnamed]
	Quests             handlerinput.Main[seedtypes.Quest, apitypes.Quest, apiresource.Quest]
	Sidequests         handlerinput.Main[seedtypes.Sidequest, apitypes.Sidequest, apiresource.Quest]
	Subquests          handlerinput.Main[seedtypes.Subquest, apitypes.Subquest, apiresource.Quest]
	ArenaCreations     handlerinput.Main[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.Named]
	BlitzballPrizes    handlerinput.Main[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.Named]
	Songs              handlerinput.Main[seedtypes.Song, apitypes.Song, apiresource.Named]
	FMVs               handlerinput.Main[seedtypes.FMV, apitypes.FMV, apiresource.Named]
	PlayerUnits        handlerinput.Main[seedtypes.PlayerUnit, apitypes.PlayerUnit, apiresource.Typed]
	Characters         handlerinput.Main[seedtypes.Character, apitypes.Character, apiresource.Named]
	Aeons              handlerinput.Main[seedtypes.Aeon, apitypes.Aeon, apiresource.Named]
	CharacterClasses   handlerinput.Main[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.Named]
	Monsters           handlerinput.Main[seedtypes.Monster, apitypes.Monster, apiresource.Named]
	Abilities          handlerinput.Main[seedtypes.Ability, apitypes.Ability, apiresource.Typed]
	PlayerAbilities    handlerinput.Main[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.Named]
	OverdriveAbilities handlerinput.Main[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.Named]
	ItemAbilities      handlerinput.Main[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.Named]
	TriggerCommands    handlerinput.Main[seedtypes.TriggerCommand, apitypes.TriggerCommand, apiresource.Named]
	MiscAbilities      handlerinput.Main[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.Named]
	EnemyAbilities     handlerinput.Main[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.Named]
	AeonCommands       handlerinput.Main[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.Named]
	OverdriveCommands  handlerinput.Main[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.Named]
	Overdrives         handlerinput.Main[seedtypes.Overdrive, apitypes.Overdrive, apiresource.Named]
	RonsoRages         handlerinput.Main[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.Named]
	Submenus           handlerinput.Main[seedtypes.Submenu, apitypes.Submenu, apiresource.Named]
	Topmenus           handlerinput.Main[seedtypes.Topmenu, apitypes.Topmenu, apiresource.Named]
	AllItems           handlerinput.Main[seedtypes.MasterItem, apitypes.MasterItem, apiresource.Typed]
	Items              handlerinput.Main[seedtypes.Item, apitypes.Item, apiresource.Named]
	KeyItems           handlerinput.Main[seedtypes.KeyItem, apitypes.KeyItem, apiresource.Named]
	Spheres            handlerinput.Main[seedtypes.Sphere, apitypes.Sphere, apiresource.Named]
	Primers            handlerinput.Main[seedtypes.Primer, apitypes.Primer, apiresource.Named]
	Mixes              handlerinput.Main[seedtypes.Mix, apitypes.Mix, apiresource.Named]
	AutoAbilities      handlerinput.Main[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.Named]
	EquipmentTables    handlerinput.Main[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.Unnamed]
	Equipment          handlerinput.Main[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.Named]
	CelestialWeapons   handlerinput.Main[seedtypes.CelestialWeapon, apitypes.CelestialWeapon, apiresource.Named]
	Stats              handlerinput.Main[seedtypes.Stat, apitypes.Stat, apiresource.Named]
	Properties         handlerinput.Main[seedtypes.Property, apitypes.Property, apiresource.Named]
	OverdriveModes     handlerinput.Main[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.Named]
	Elements           handlerinput.Main[seedtypes.Element, apitypes.Element, apiresource.Named]
	StatusConditions   handlerinput.Main[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.Named]
	Modifiers          handlerinput.Main[seedtypes.Modifier, apitypes.Modifier, apiresource.Named]
	AgilityTiers       handlerinput.Main[seedtypes.AgilityTier, apitypes.AgilityTier, apiresource.Unnamed]
	Enums              handlerinput.Enums
	Endpoints          handlerinput.Endpoints
	AlBhed             handlerinput.Service[servicetypes.AlBhedParams, servicetypes.AlBhedResponse]
	TurnOrder          handlerinput.Service[servicetypes.TurnOrderParams, servicetypes.TurnOrderResponse]
	DropChance         handlerinput.Service[servicetypes.DropChanceParams, servicetypes.DropChanceResponse]
	CritChance         handlerinput.Service[servicetypes.CritChanceParams, servicetypes.CritChanceResponse]
}

func (E *Endpoints) InitializeEndpointSlice() {
	E.Endpoints.Slice = []ts.EndpointName{
		E.Locations.Endpoint,
		E.Sublocations.Endpoint,
		E.Areas.Endpoint,
		E.MonsterFormations.Endpoint,
		E.Shops.Endpoint,
		E.Treasures.Endpoint,
		E.Quests.Endpoint,
		E.Sidequests.Endpoint,
		E.Subquests.Endpoint,
		E.ArenaCreations.Endpoint,
		E.BlitzballPrizes.Endpoint,
		E.Songs.Endpoint,
		E.FMVs.Endpoint,
		E.PlayerUnits.Endpoint,
		E.Characters.Endpoint,
		E.Aeons.Endpoint,
		E.CharacterClasses.Endpoint,
		E.Monsters.Endpoint,
		E.Abilities.Endpoint,
		E.PlayerAbilities.Endpoint,
		E.OverdriveAbilities.Endpoint,
		E.ItemAbilities.Endpoint,
		E.TriggerCommands.Endpoint,
		E.MiscAbilities.Endpoint,
		E.EnemyAbilities.Endpoint,
		E.AeonCommands.Endpoint,
		E.OverdriveCommands.Endpoint,
		E.Overdrives.Endpoint,
		E.RonsoRages.Endpoint,
		E.Submenus.Endpoint,
		E.Topmenus.Endpoint,
		E.AllItems.Endpoint,
		E.Items.Endpoint,
		E.Spheres.Endpoint,
		E.KeyItems.Endpoint,
		E.Primers.Endpoint,
		E.Mixes.Endpoint,
		E.AutoAbilities.Endpoint,
		E.EquipmentTables.Endpoint,
		E.Equipment.Endpoint,
		E.CelestialWeapons.Endpoint,
		E.Stats.Endpoint,
		E.Properties.Endpoint,
		E.OverdriveModes.Endpoint,
		E.Elements.Endpoint,
		E.StatusConditions.Endpoint,
		E.Modifiers.Endpoint,
		E.AgilityTiers.Endpoint,
		E.Enums.Endpoint,
		E.Endpoints.Endpoint,
		E.AlBhed.Endpoint,
		E.TurnOrder.Endpoint,
		E.DropChance.Endpoint,
		E.CritChance.Endpoint,
	}

	slices.Sort(E.Endpoints.Slice)
}
