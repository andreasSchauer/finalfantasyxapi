package api

import "net/http"

func (cfg *Config) HandleLocations(w http.ResponseWriter, r *http.Request) {
	E.Locations.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleSublocations(w http.ResponseWriter, r *http.Request) {
	E.Sublocations.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleAreas(w http.ResponseWriter, r *http.Request) {
	E.Areas.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleMonsterFormations(w http.ResponseWriter, r *http.Request) {
	E.MonsterFormations.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleShops(w http.ResponseWriter, r *http.Request) {
	E.Shops.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleTreasures(w http.ResponseWriter, r *http.Request) {
	E.Treasures.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleQuests(w http.ResponseWriter, r *http.Request) {
	E.Quests.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleSidequests(w http.ResponseWriter, r *http.Request) {
	E.Sidequests.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleSubquests(w http.ResponseWriter, r *http.Request) {
	E.Subquests.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleArenaCreations(w http.ResponseWriter, r *http.Request) {
	E.ArenaCreations.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleBlitzballPrizes(w http.ResponseWriter, r *http.Request) {
	E.BlitzballPrizes.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleSongs(w http.ResponseWriter, r *http.Request) {
	E.Songs.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleFMVs(w http.ResponseWriter, r *http.Request) {
	E.FMVs.routerNameID(cfg, w, r)
}

func (cfg *Config) HandlePlayerUnits(w http.ResponseWriter, r *http.Request) {
	E.PlayerUnits.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleCharacters(w http.ResponseWriter, r *http.Request) {
	E.Characters.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleAeons(w http.ResponseWriter, r *http.Request) {
	E.Aeons.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleCharacterClasses(w http.ResponseWriter, r *http.Request) {
	E.CharacterClasses.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleMonsters(w http.ResponseWriter, r *http.Request) {
	E.Monsters.routerNameVersion(cfg, w, r)
}

func (cfg *Config) HandleAbilities(w http.ResponseWriter, r *http.Request) {
	E.Abilities.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandlePlayerAbilities(w http.ResponseWriter, r *http.Request) {
	E.PlayerAbilities.routerNameVersion(cfg, w, r)
}

func (cfg *Config) HandleOverdriveAbilities(w http.ResponseWriter, r *http.Request) {
	E.OverdriveAbilities.routerNameVersion(cfg, w, r)
}

func (cfg *Config) HandleItemAbilities(w http.ResponseWriter, r *http.Request) {
	E.ItemAbilities.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleTriggerCommands(w http.ResponseWriter, r *http.Request) {
	E.TriggerCommands.routerNameVersion(cfg, w, r)
}

func (cfg *Config) HandleMiscAbilities(w http.ResponseWriter, r *http.Request) {
	E.MiscAbilities.routerNameVersion(cfg, w, r)
}

func (cfg *Config) HandleEnemyAbilities(w http.ResponseWriter, r *http.Request) {
	E.EnemyAbilities.routerNameVersion(cfg, w, r)
}

func (cfg *Config) HandleAeonCommands(w http.ResponseWriter, r *http.Request) {
	E.AeonCommands.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleOverdriveCommands(w http.ResponseWriter, r *http.Request) {
	E.OverdriveCommands.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleOverdrives(w http.ResponseWriter, r *http.Request) {
	E.Overdrives.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleRonsoRages(w http.ResponseWriter, r *http.Request) {
	E.RonsoRages.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleSubmenus(w http.ResponseWriter, r *http.Request) {
	E.Submenus.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleTopmenus(w http.ResponseWriter, r *http.Request) {
	E.Topmenus.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleAllItems(w http.ResponseWriter, r *http.Request) {
	E.AllItems.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleItems(w http.ResponseWriter, r *http.Request) {
	E.Items.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleKeyItems(w http.ResponseWriter, r *http.Request) {
	E.KeyItems.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleSpheres(w http.ResponseWriter, r *http.Request) {
	E.Spheres.routerNameID(cfg, w, r)
}

func (cfg *Config) HandlePrimers(w http.ResponseWriter, r *http.Request) {
	E.Primers.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleMixes(w http.ResponseWriter, r *http.Request) {
	E.Mixes.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleAutoAbilities(w http.ResponseWriter, r *http.Request) {
	E.AutoAbilities.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleEquipmentTables(w http.ResponseWriter, r *http.Request) {
	E.EquipmentTables.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleEquipment(w http.ResponseWriter, r *http.Request) {
	E.Equipment.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleCelestialWeapons(w http.ResponseWriter, r *http.Request) {
	E.CelestialWeapons.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleStats(w http.ResponseWriter, r *http.Request) {
	E.Stats.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleProperties(w http.ResponseWriter, r *http.Request) {
	E.Properties.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleOverdriveModes(w http.ResponseWriter, r *http.Request) {
	E.OverdriveModes.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleElements(w http.ResponseWriter, r *http.Request) {
	E.Elements.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleStatusConditions(w http.ResponseWriter, r *http.Request) {
	E.StatusConditions.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleModifiers(w http.ResponseWriter, r *http.Request) {
	E.Modifiers.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleAgilityTiers(w http.ResponseWriter, r *http.Request) {
	E.AgilityTiers.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleEnums(w http.ResponseWriter, r *http.Request) {
	E.Enums.router(cfg, w, r)
}

func (cfg *Config) HandleEndpoints(w http.ResponseWriter, r *http.Request) {
	E.Endpoints.router(cfg, w, r)
}

func (cfg *Config) HandleAlBhedGet(w http.ResponseWriter, r *http.Request) {
	E.AlBhed.routerGet(cfg, w, r)
}

func (cfg *Config) HandleAlBhedPost(w http.ResponseWriter, r *http.Request) {
	E.AlBhed.routerPost(cfg, w, r)
}

func (cfg *Config) HandleTurnOrderGet(w http.ResponseWriter, r *http.Request) {
	E.TurnOrder.routerGet(cfg, w, r)
}

func (cfg *Config) HandleTurnOrderPost(w http.ResponseWriter, r *http.Request) {
	E.TurnOrder.routerPost(cfg, w, r)
}

func (cfg *Config) HandleDropChanceGet(w http.ResponseWriter, r *http.Request) {
	E.DropChance.routerGet(cfg, w, r)
}

func (cfg *Config) HandleDropChancePost(w http.ResponseWriter, r *http.Request) {
	E.DropChance.routerPost(cfg, w, r)
}

func (cfg *Config) HandleCritChanceGet(w http.ResponseWriter, r *http.Request) {
	E.CritChance.routerGet(cfg, w, r)
}

func (cfg *Config) HandleCritChancePost(w http.ResponseWriter, r *http.Request) {
	E.CritChance.routerPost(cfg, w, r)
}
