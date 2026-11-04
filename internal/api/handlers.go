package api

import "net/http"

func (srv *Server) HandleLocations(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Locations)
}

func (srv *Server) HandleSublocations(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Sublocations)
}

func (srv *Server) HandleAreas(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(srv, w, r, E.Areas)
}

func (srv *Server) HandleMonsterFormations(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(srv, w, r, E.MonsterFormations)
}

func (srv *Server) HandleShops(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(srv, w, r, E.Shops)
}

func (srv *Server) HandleTreasures(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(srv, w, r, E.Treasures)
}

func (srv *Server) HandleQuests(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(srv, w, r, E.Quests)
}

func (srv *Server) HandleSidequests(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Sidequests)
}

func (srv *Server) HandleSubquests(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(srv, w, r, E.Subquests)
}

func (srv *Server) HandleArenaCreations(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.ArenaCreations)
}

func (srv *Server) HandleBlitzballPrizes(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(srv, w, r, E.BlitzballPrizes)
}

func (srv *Server) HandleSongs(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Songs)
}

func (srv *Server) HandleFMVs(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.FMVs)
}

func (srv *Server) HandlePlayerUnits(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.PlayerUnits)
}

func (srv *Server) HandleCharacters(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Characters)
}

func (srv *Server) HandleAeons(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Aeons)
}

func (srv *Server) HandleCharacterClasses(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.CharacterClasses)
}

func (srv *Server) HandleMonsters(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(srv, w, r, E.Monsters)
}

func (srv *Server) HandleAbilities(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(srv, w, r, E.Abilities)
}

func (srv *Server) HandlePlayerAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(srv, w, r, E.PlayerAbilities)
}

func (srv *Server) HandleOverdriveAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(srv, w, r, E.OverdriveAbilities)
}

func (srv *Server) HandleItemAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.ItemAbilities)
}

func (srv *Server) HandleTriggerCommands(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(srv, w, r, E.TriggerCommands)
}

func (srv *Server) HandleMiscAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(srv, w, r, E.MiscAbilities)
}

func (srv *Server) HandleEnemyAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(srv, w, r, E.EnemyAbilities)
}

func (srv *Server) HandleAeonCommands(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.AeonCommands)
}

func (srv *Server) HandleOverdriveCommands(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.OverdriveCommands)
}

func (srv *Server) HandleOverdrives(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Overdrives)
}

func (srv *Server) HandleRonsoRages(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.RonsoRages)
}

func (srv *Server) HandleSubmenus(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Submenus)
}

func (srv *Server) HandleTopmenus(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Topmenus)
}

func (srv *Server) HandleAllItems(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.AllItems)
}

func (srv *Server) HandleItems(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Items)
}

func (srv *Server) HandleKeyItems(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.KeyItems)
}

func (srv *Server) HandleSpheres(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Spheres)
}

func (srv *Server) HandlePrimers(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Primers)
}

func (srv *Server) HandleMixes(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Mixes)
}

func (srv *Server) HandleAutoAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.AutoAbilities)
}

func (srv *Server) HandleEquipmentTables(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(srv, w, r, E.EquipmentTables)
}

func (srv *Server) HandleEquipment(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Equipment)
}

func (srv *Server) HandleCelestialWeapons(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.CelestialWeapons)
}

func (srv *Server) HandleStats(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Stats)
}

func (srv *Server) HandleProperties(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Properties)
}

func (srv *Server) HandleOverdriveModes(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.OverdriveModes)
}

func (srv *Server) HandleElements(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Elements)
}

func (srv *Server) HandleStatusConditions(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.StatusConditions)
}

func (srv *Server) HandleModifiers(w http.ResponseWriter, r *http.Request) {
	routerNameID(srv, w, r, E.Modifiers)
}

func (srv *Server) HandleAgilityTiers(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(srv, w, r, E.AgilityTiers)
}

func (srv *Server) HandleEnums(w http.ResponseWriter, r *http.Request) {
	routerEnums(srv, w, r, E.Enums)
}

func (srv *Server) HandleEndpoints(w http.ResponseWriter, r *http.Request) {
	routerEndpoints(srv, w, r, E.Endpoints)
}

func (srv *Server) HandleAlBhedGet(w http.ResponseWriter, r *http.Request) {
	routerServiceGet(srv, w, r, E.AlBhed)
}

func (srv *Server) HandleAlBhedPost(w http.ResponseWriter, r *http.Request) {
	routerServicePost(srv, w, r, E.AlBhed)
}

func (srv *Server) HandleTurnOrderGet(w http.ResponseWriter, r *http.Request) {
	routerServiceGet(srv, w, r, E.TurnOrder)
}

func (srv *Server) HandleTurnOrderPost(w http.ResponseWriter, r *http.Request) {
	routerServicePost(srv, w, r, E.TurnOrder)
}

func (srv *Server) HandleDropChanceGet(w http.ResponseWriter, r *http.Request) {
	routerServiceGet(srv, w, r, E.DropChance)
}

func (srv *Server) HandleDropChancePost(w http.ResponseWriter, r *http.Request) {
	routerServicePost(srv, w, r, E.DropChance)
}

func (srv *Server) HandleCritChanceGet(w http.ResponseWriter, r *http.Request) {
	routerServiceGet(srv, w, r, E.CritChance)
}

func (srv *Server) HandleCritChancePost(w http.ResponseWriter, r *http.Request) {
	routerServicePost(srv, w, r, E.CritChance)
}
