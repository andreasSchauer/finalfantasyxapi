package api

import "net/http"

func (cfg *Config) HandleLocations(w http.ResponseWriter, r *http.Request) {
	cfg.E.Locations.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleSublocations(w http.ResponseWriter, r *http.Request) {
	cfg.E.Sublocations.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleAreas(w http.ResponseWriter, r *http.Request) {
	cfg.E.Areas.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleMonsterFormations(w http.ResponseWriter, r *http.Request) {
	cfg.E.MonsterFormations.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleShops(w http.ResponseWriter, r *http.Request) {
	cfg.E.Shops.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleTreasures(w http.ResponseWriter, r *http.Request) {
	cfg.E.Treasures.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleQuests(w http.ResponseWriter, r *http.Request) {
	cfg.E.Quests.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleSidequests(w http.ResponseWriter, r *http.Request) {
	cfg.E.Sidequests.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleSubquests(w http.ResponseWriter, r *http.Request) {
	cfg.E.Subquests.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleArenaCreations(w http.ResponseWriter, r *http.Request) {
	cfg.E.ArenaCreations.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleBlitzballPrizes(w http.ResponseWriter, r *http.Request) {
	cfg.E.BlitzballPrizes.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleSongs(w http.ResponseWriter, r *http.Request) {
	cfg.E.Songs.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleFMVs(w http.ResponseWriter, r *http.Request) {
	cfg.E.FMVs.routerNameID(cfg, w, r)
}

func (cfg *Config) HandlePlayerUnits(w http.ResponseWriter, r *http.Request) {
	cfg.E.PlayerUnits.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleCharacters(w http.ResponseWriter, r *http.Request) {
	cfg.E.Characters.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleAeons(w http.ResponseWriter, r *http.Request) {
	cfg.E.Aeons.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleCharacterClasses(w http.ResponseWriter, r *http.Request) {
	cfg.E.CharacterClasses.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleMonsters(w http.ResponseWriter, r *http.Request) {
	cfg.E.Monsters.routerNameVersion(cfg, w, r)
}

func (cfg *Config) HandleAbilities(w http.ResponseWriter, r *http.Request) {
	cfg.E.Abilities.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandlePlayerAbilities(w http.ResponseWriter, r *http.Request) {
	cfg.E.PlayerAbilities.routerNameVersion(cfg, w, r)
}

func (cfg *Config) HandleOverdriveAbilities(w http.ResponseWriter, r *http.Request) {
	cfg.E.OverdriveAbilities.routerNameVersion(cfg, w, r)
}

func (cfg *Config) HandleItemAbilities(w http.ResponseWriter, r *http.Request) {
	cfg.E.ItemAbilities.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleTriggerCommands(w http.ResponseWriter, r *http.Request) {
	cfg.E.TriggerCommands.routerNameVersion(cfg, w, r)
}

func (cfg *Config) HandleMiscAbilities(w http.ResponseWriter, r *http.Request) {
	cfg.E.MiscAbilities.routerNameVersion(cfg, w, r)
}

func (cfg *Config) HandleEnemyAbilities(w http.ResponseWriter, r *http.Request) {
	cfg.E.EnemyAbilities.routerNameVersion(cfg, w, r)
}

func (cfg *Config) HandleAeonCommands(w http.ResponseWriter, r *http.Request) {
	cfg.E.AeonCommands.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleOverdriveCommands(w http.ResponseWriter, r *http.Request) {
	cfg.E.OverdriveCommands.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleOverdrives(w http.ResponseWriter, r *http.Request) {
	cfg.E.Overdrives.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleRonsoRages(w http.ResponseWriter, r *http.Request) {
	cfg.E.RonsoRages.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleSubmenus(w http.ResponseWriter, r *http.Request) {
	cfg.E.Submenus.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleTopmenus(w http.ResponseWriter, r *http.Request) {
	cfg.E.Topmenus.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleAllItems(w http.ResponseWriter, r *http.Request) {
	cfg.E.AllItems.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleItems(w http.ResponseWriter, r *http.Request) {
	cfg.E.Items.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleKeyItems(w http.ResponseWriter, r *http.Request) {
	cfg.E.KeyItems.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleSpheres(w http.ResponseWriter, r *http.Request) {
	cfg.E.Spheres.routerNameID(cfg, w, r)
}

func (cfg *Config) HandlePrimers(w http.ResponseWriter, r *http.Request) {
	cfg.E.Primers.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleMixes(w http.ResponseWriter, r *http.Request) {
	cfg.E.Mixes.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleAutoAbilities(w http.ResponseWriter, r *http.Request) {
	cfg.E.AutoAbilities.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleEquipmentTables(w http.ResponseWriter, r *http.Request) {
	cfg.E.EquipmentTables.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleEquipment(w http.ResponseWriter, r *http.Request) {
	cfg.E.Equipment.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleCelestialWeapons(w http.ResponseWriter, r *http.Request) {
	cfg.E.CelestialWeapons.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleStats(w http.ResponseWriter, r *http.Request) {
	cfg.E.Stats.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleProperties(w http.ResponseWriter, r *http.Request) {
	cfg.E.Properties.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleOverdriveModes(w http.ResponseWriter, r *http.Request) {
	cfg.E.OverdriveModes.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleElements(w http.ResponseWriter, r *http.Request) {
	cfg.E.Elements.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleStatusConditions(w http.ResponseWriter, r *http.Request) {
	cfg.E.StatusConditions.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleModifiers(w http.ResponseWriter, r *http.Request) {
	cfg.E.Modifiers.routerNameID(cfg, w, r)
}

func (cfg *Config) HandleAgilityTiers(w http.ResponseWriter, r *http.Request) {
	cfg.E.AgilityTiers.routerIDOnly(cfg, w, r)
}

func (cfg *Config) HandleEnums(w http.ResponseWriter, r *http.Request) {
	cfg.E.Enums.router(cfg, w, r)
}

func (cfg *Config) HandleEndpoints(w http.ResponseWriter, r *http.Request) {
	cfg.E.Endpoints.router(cfg, w, r)
}

func (cfg *Config) HandleAlBhedGet(w http.ResponseWriter, r *http.Request) {
	cfg.E.AlBhed.routerGet(cfg, w, r)
}

func (cfg *Config) HandleAlBhedPost(w http.ResponseWriter, r *http.Request) {
	cfg.E.AlBhed.routerPost(cfg, w, r)
}

func (cfg *Config) HandleTurnOrderGet(w http.ResponseWriter, r *http.Request) {
	cfg.E.TurnOrder.routerGet(cfg, w, r)
}

func (cfg *Config) HandleTurnOrderPost(w http.ResponseWriter, r *http.Request) {
	cfg.E.TurnOrder.routerPost(cfg, w, r)
}

func (cfg *Config) HandleDropChanceGet(w http.ResponseWriter, r *http.Request) {
	cfg.E.DropChance.routerGet(cfg, w, r)
}

func (cfg *Config) HandleDropChancePost(w http.ResponseWriter, r *http.Request) {
	cfg.E.DropChance.routerPost(cfg, w, r)
}

func (cfg *Config) HandleCritChanceGet(w http.ResponseWriter, r *http.Request) {
	cfg.E.CritChance.routerGet(cfg, w, r)
}

func (cfg *Config) HandleCritChancePost(w http.ResponseWriter, r *http.Request) {
	cfg.E.CritChance.routerPost(cfg, w, r)
}
