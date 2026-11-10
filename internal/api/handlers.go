package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
)

func (srv *Server) HandleLocations(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Locations)
}

func (srv *Server) HandleSublocations(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Sublocations)
}

func (srv *Server) HandleAreas(w http.ResponseWriter, r *http.Request) {
	routeIDOnly(w, r, apilookups.E.Areas)
}

func (srv *Server) HandleMonsterFormations(w http.ResponseWriter, r *http.Request) {
	routeIDOnly(w, r, apilookups.E.MonsterFormations)
}

func (srv *Server) HandleShops(w http.ResponseWriter, r *http.Request) {
	routeIDOnly(w, r, apilookups.E.Shops)
}

func (srv *Server) HandleTreasures(w http.ResponseWriter, r *http.Request) {
	routeIDOnly(w, r, apilookups.E.Treasures)
}

func (srv *Server) HandleQuests(w http.ResponseWriter, r *http.Request) {
	routeIDOnly(w, r, apilookups.E.Quests)
}

func (srv *Server) HandleSidequests(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Sidequests)
}

func (srv *Server) HandleSubquests(w http.ResponseWriter, r *http.Request) {
	routeIDOnly(w, r, apilookups.E.Subquests)
}

func (srv *Server) HandleArenaCreations(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.ArenaCreations)
}

func (srv *Server) HandleBlitzballPrizes(w http.ResponseWriter, r *http.Request) {
	routeIDOnly(w, r, apilookups.E.BlitzballPrizes)
}

func (srv *Server) HandleSongs(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Songs)
}

func (srv *Server) HandleFMVs(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.FMVs)
}

func (srv *Server) HandlePlayerUnits(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.PlayerUnits)
}

func (srv *Server) HandleCharacters(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Characters)
}

func (srv *Server) HandleAeons(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Aeons)
}

func (srv *Server) HandleCharacterClasses(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.CharacterClasses)
}

func (srv *Server) HandleMonsters(w http.ResponseWriter, r *http.Request) {
	routeNameVersion(w, r, apilookups.E.Monsters)
}

func (srv *Server) HandleAbilities(w http.ResponseWriter, r *http.Request) {
	routeIDOnly(w, r, apilookups.E.Abilities)
}

func (srv *Server) HandlePlayerAbilities(w http.ResponseWriter, r *http.Request) {
	routeNameVersion(w, r, apilookups.E.PlayerAbilities)
}

func (srv *Server) HandleOverdriveAbilities(w http.ResponseWriter, r *http.Request) {
	routeNameVersion(w, r, apilookups.E.OverdriveAbilities)
}

func (srv *Server) HandleItemAbilities(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.ItemAbilities)
}

func (srv *Server) HandleTriggerCommands(w http.ResponseWriter, r *http.Request) {
	routeNameVersion(w, r, apilookups.E.TriggerCommands)
}

func (srv *Server) HandleMiscAbilities(w http.ResponseWriter, r *http.Request) {
	routeNameVersion(w, r, apilookups.E.MiscAbilities)
}

func (srv *Server) HandleEnemyAbilities(w http.ResponseWriter, r *http.Request) {
	routeNameVersion(w, r, apilookups.E.EnemyAbilities)
}

func (srv *Server) HandleAeonCommands(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.AeonCommands)
}

func (srv *Server) HandleOverdriveCommands(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.OverdriveCommands)
}

func (srv *Server) HandleOverdrives(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Overdrives)
}

func (srv *Server) HandleRonsoRages(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.RonsoRages)
}

func (srv *Server) HandleSubmenus(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Submenus)
}

func (srv *Server) HandleTopmenus(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Topmenus)
}

func (srv *Server) HandleAllItems(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.AllItems)
}

func (srv *Server) HandleItems(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Items)
}

func (srv *Server) HandleKeyItems(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.KeyItems)
}

func (srv *Server) HandleSpheres(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Spheres)
}

func (srv *Server) HandlePrimers(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Primers)
}

func (srv *Server) HandleMixes(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Mixes)
}

func (srv *Server) HandleAutoAbilities(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.AutoAbilities)
}

func (srv *Server) HandleEquipmentTables(w http.ResponseWriter, r *http.Request) {
	routeIDOnly(w, r, apilookups.E.EquipmentTables)
}

func (srv *Server) HandleEquipment(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Equipment)
}

func (srv *Server) HandleCelestialWeapons(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.CelestialWeapons)
}

func (srv *Server) HandleStats(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Stats)
}

func (srv *Server) HandleProperties(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Properties)
}

func (srv *Server) HandleOverdriveModes(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.OverdriveModes)
}

func (srv *Server) HandleElements(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Elements)
}

func (srv *Server) HandleStatusConditions(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.StatusConditions)
}

func (srv *Server) HandleModifiers(w http.ResponseWriter, r *http.Request) {
	routeNameID(w, r, apilookups.E.Modifiers)
}

func (srv *Server) HandleAgilityTiers(w http.ResponseWriter, r *http.Request) {
	routeIDOnly(w, r, apilookups.E.AgilityTiers)
}

func (srv *Server) HandleEnums(w http.ResponseWriter, r *http.Request) {
	routeEnums(w, r, apilookups.E.Enums)
}

func (srv *Server) HandleEndpoints(w http.ResponseWriter, r *http.Request) {
	routeEndpoints(w, r, apilookups.E.Endpoints)
}

func (srv *Server) HandleAlBhedGet(w http.ResponseWriter, r *http.Request) {
	routeServiceGet(w, r, apilookups.E.AlBhed)
}

func (srv *Server) HandleAlBhedPost(w http.ResponseWriter, r *http.Request) {
	routeServicePost(w, r, apilookups.E.AlBhed)
}

func (srv *Server) HandleTurnOrderGet(w http.ResponseWriter, r *http.Request) {
	routeServiceGet(w, r, apilookups.E.TurnOrder)
}

func (srv *Server) HandleTurnOrderPost(w http.ResponseWriter, r *http.Request) {
	routeServicePost(w, r, apilookups.E.TurnOrder)
}

func (srv *Server) HandleDropChanceGet(w http.ResponseWriter, r *http.Request) {
	routeServiceGet(w, r, apilookups.E.DropChance)
}

func (srv *Server) HandleDropChancePost(w http.ResponseWriter, r *http.Request) {
	routeServicePost(w, r, apilookups.E.DropChance)
}

func (srv *Server) HandleCritChanceGet(w http.ResponseWriter, r *http.Request) {
	routeServiceGet(w, r, apilookups.E.CritChance)
}

func (srv *Server) HandleCritChancePost(w http.ResponseWriter, r *http.Request) {
	routeServicePost(w, r, apilookups.E.CritChance)
}
