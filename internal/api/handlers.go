package api

import "net/http"

func (s *Server) HandleLocations(w http.ResponseWriter, r *http.Request) {
	E.Locations.routerNameID(s, w, r)
}

func (s *Server) HandleSublocations(w http.ResponseWriter, r *http.Request) {
	E.Sublocations.routerNameID(s, w, r)
}

func (s *Server) HandleAreas(w http.ResponseWriter, r *http.Request) {
	E.Areas.routerIDOnly(s, w, r)
}

func (s *Server) HandleMonsterFormations(w http.ResponseWriter, r *http.Request) {
	E.MonsterFormations.routerIDOnly(s, w, r)
}

func (s *Server) HandleShops(w http.ResponseWriter, r *http.Request) {
	E.Shops.routerIDOnly(s, w, r)
}

func (s *Server) HandleTreasures(w http.ResponseWriter, r *http.Request) {
	E.Treasures.routerIDOnly(s, w, r)
}

func (s *Server) HandleQuests(w http.ResponseWriter, r *http.Request) {
	E.Quests.routerIDOnly(s, w, r)
}

func (s *Server) HandleSidequests(w http.ResponseWriter, r *http.Request) {
	E.Sidequests.routerNameID(s, w, r)
}

func (s *Server) HandleSubquests(w http.ResponseWriter, r *http.Request) {
	E.Subquests.routerIDOnly(s, w, r)
}

func (s *Server) HandleArenaCreations(w http.ResponseWriter, r *http.Request) {
	E.ArenaCreations.routerNameID(s, w, r)
}

func (s *Server) HandleBlitzballPrizes(w http.ResponseWriter, r *http.Request) {
	E.BlitzballPrizes.routerIDOnly(s, w, r)
}

func (s *Server) HandleSongs(w http.ResponseWriter, r *http.Request) {
	E.Songs.routerNameID(s, w, r)
}

func (s *Server) HandleFMVs(w http.ResponseWriter, r *http.Request) {
	E.FMVs.routerNameID(s, w, r)
}

func (s *Server) HandlePlayerUnits(w http.ResponseWriter, r *http.Request) {
	E.PlayerUnits.routerNameID(s, w, r)
}

func (s *Server) HandleCharacters(w http.ResponseWriter, r *http.Request) {
	E.Characters.routerNameID(s, w, r)
}

func (s *Server) HandleAeons(w http.ResponseWriter, r *http.Request) {
	E.Aeons.routerNameID(s, w, r)
}

func (s *Server) HandleCharacterClasses(w http.ResponseWriter, r *http.Request) {
	E.CharacterClasses.routerNameID(s, w, r)
}

func (s *Server) HandleMonsters(w http.ResponseWriter, r *http.Request) {
	E.Monsters.routerNameVersion(s, w, r)
}

func (s *Server) HandleAbilities(w http.ResponseWriter, r *http.Request) {
	E.Abilities.routerIDOnly(s, w, r)
}

func (s *Server) HandlePlayerAbilities(w http.ResponseWriter, r *http.Request) {
	E.PlayerAbilities.routerNameVersion(s, w, r)
}

func (s *Server) HandleOverdriveAbilities(w http.ResponseWriter, r *http.Request) {
	E.OverdriveAbilities.routerNameVersion(s, w, r)
}

func (s *Server) HandleItemAbilities(w http.ResponseWriter, r *http.Request) {
	E.ItemAbilities.routerNameID(s, w, r)
}

func (s *Server) HandleTriggerCommands(w http.ResponseWriter, r *http.Request) {
	E.TriggerCommands.routerNameVersion(s, w, r)
}

func (s *Server) HandleMiscAbilities(w http.ResponseWriter, r *http.Request) {
	E.MiscAbilities.routerNameVersion(s, w, r)
}

func (s *Server) HandleEnemyAbilities(w http.ResponseWriter, r *http.Request) {
	E.EnemyAbilities.routerNameVersion(s, w, r)
}

func (s *Server) HandleAeonCommands(w http.ResponseWriter, r *http.Request) {
	E.AeonCommands.routerNameID(s, w, r)
}

func (s *Server) HandleOverdriveCommands(w http.ResponseWriter, r *http.Request) {
	E.OverdriveCommands.routerNameID(s, w, r)
}

func (s *Server) HandleOverdrives(w http.ResponseWriter, r *http.Request) {
	E.Overdrives.routerNameID(s, w, r)
}

func (s *Server) HandleRonsoRages(w http.ResponseWriter, r *http.Request) {
	E.RonsoRages.routerNameID(s, w, r)
}

func (s *Server) HandleSubmenus(w http.ResponseWriter, r *http.Request) {
	E.Submenus.routerNameID(s, w, r)
}

func (s *Server) HandleTopmenus(w http.ResponseWriter, r *http.Request) {
	E.Topmenus.routerNameID(s, w, r)
}

func (s *Server) HandleAllItems(w http.ResponseWriter, r *http.Request) {
	E.AllItems.routerNameID(s, w, r)
}

func (s *Server) HandleItems(w http.ResponseWriter, r *http.Request) {
	E.Items.routerNameID(s, w, r)
}

func (s *Server) HandleKeyItems(w http.ResponseWriter, r *http.Request) {
	E.KeyItems.routerNameID(s, w, r)
}

func (s *Server) HandleSpheres(w http.ResponseWriter, r *http.Request) {
	E.Spheres.routerNameID(s, w, r)
}

func (s *Server) HandlePrimers(w http.ResponseWriter, r *http.Request) {
	E.Primers.routerNameID(s, w, r)
}

func (s *Server) HandleMixes(w http.ResponseWriter, r *http.Request) {
	E.Mixes.routerNameID(s, w, r)
}

func (s *Server) HandleAutoAbilities(w http.ResponseWriter, r *http.Request) {
	E.AutoAbilities.routerNameID(s, w, r)
}

func (s *Server) HandleEquipmentTables(w http.ResponseWriter, r *http.Request) {
	E.EquipmentTables.routerIDOnly(s, w, r)
}

func (s *Server) HandleEquipment(w http.ResponseWriter, r *http.Request) {
	E.Equipment.routerNameID(s, w, r)
}

func (s *Server) HandleCelestialWeapons(w http.ResponseWriter, r *http.Request) {
	E.CelestialWeapons.routerNameID(s, w, r)
}

func (s *Server) HandleStats(w http.ResponseWriter, r *http.Request) {
	E.Stats.routerNameID(s, w, r)
}

func (s *Server) HandleProperties(w http.ResponseWriter, r *http.Request) {
	E.Properties.routerNameID(s, w, r)
}

func (s *Server) HandleOverdriveModes(w http.ResponseWriter, r *http.Request) {
	E.OverdriveModes.routerNameID(s, w, r)
}

func (s *Server) HandleElements(w http.ResponseWriter, r *http.Request) {
	E.Elements.routerNameID(s, w, r)
}

func (s *Server) HandleStatusConditions(w http.ResponseWriter, r *http.Request) {
	E.StatusConditions.routerNameID(s, w, r)
}

func (s *Server) HandleModifiers(w http.ResponseWriter, r *http.Request) {
	E.Modifiers.routerNameID(s, w, r)
}

func (s *Server) HandleAgilityTiers(w http.ResponseWriter, r *http.Request) {
	E.AgilityTiers.routerIDOnly(s, w, r)
}

func (s *Server) HandleEnums(w http.ResponseWriter, r *http.Request) {
	E.Enums.router(s, w, r)
}

func (s *Server) HandleEndpoints(w http.ResponseWriter, r *http.Request) {
	E.Endpoints.router(s, w, r)
}

func (s *Server) HandleAlBhedGet(w http.ResponseWriter, r *http.Request) {
	E.AlBhed.routerGet(s, w, r)
}

func (s *Server) HandleAlBhedPost(w http.ResponseWriter, r *http.Request) {
	E.AlBhed.routerPost(s, w, r)
}

func (s *Server) HandleTurnOrderGet(w http.ResponseWriter, r *http.Request) {
	E.TurnOrder.routerGet(s, w, r)
}

func (s *Server) HandleTurnOrderPost(w http.ResponseWriter, r *http.Request) {
	E.TurnOrder.routerPost(s, w, r)
}

func (s *Server) HandleDropChanceGet(w http.ResponseWriter, r *http.Request) {
	E.DropChance.routerGet(s, w, r)
}

func (s *Server) HandleDropChancePost(w http.ResponseWriter, r *http.Request) {
	E.DropChance.routerPost(s, w, r)
}

func (s *Server) HandleCritChanceGet(w http.ResponseWriter, r *http.Request) {
	E.CritChance.routerGet(s, w, r)
}

func (s *Server) HandleCritChancePost(w http.ResponseWriter, r *http.Request) {
	E.CritChance.routerPost(s, w, r)
}
