package api

import "net/http"

func (cfg *Config) HandleLocations(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Locations)
}

func (cfg *Config) HandleSublocations(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Sublocations)
}

func (cfg *Config) HandleAreas(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(cfg, w, r, cfg.E.Areas)
}

func (cfg *Config) HandleMonsterFormations(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(cfg, w, r, cfg.E.MonsterFormations)
}

func (cfg *Config) HandleShops(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(cfg, w, r, cfg.E.Shops)
}

func (cfg *Config) HandleTreasures(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(cfg, w, r, cfg.E.Treasures)
}

func (cfg *Config) HandleQuests(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(cfg, w, r, cfg.E.Quests)
}

func (cfg *Config) HandleSidequests(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Sidequests)
}

func (cfg *Config) HandleSubquests(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(cfg, w, r, cfg.E.Subquests)
}

func (cfg *Config) HandleArenaCreations(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.ArenaCreations)
}

func (cfg *Config) HandleBlitzballPrizes(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(cfg, w, r, cfg.E.BlitzballPrizes)
}

func (cfg *Config) HandleSongs(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Songs)
}

func (cfg *Config) HandleFMVs(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.FMVs)
}

func (cfg *Config) HandlePlayerUnits(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.PlayerUnits)
}

func (cfg *Config) HandleCharacters(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Characters)
}

func (cfg *Config) HandleAeons(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Aeons)
}

func (cfg *Config) HandleCharacterClasses(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.CharacterClasses)
}

func (cfg *Config) HandleMonsters(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.E.Monsters)
}

func (cfg *Config) HandleAbilities(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(cfg, w, r, cfg.E.Abilities)
}

func (cfg *Config) HandlePlayerAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.E.PlayerAbilities)
}

func (cfg *Config) HandleOverdriveAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.E.OverdriveAbilities)
}

func (cfg *Config) HandleItemAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.ItemAbilities)
}

func (cfg *Config) HandleTriggerCommands(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.E.TriggerCommands)
}

func (cfg *Config) HandleMiscAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.E.MiscAbilities)
}

func (cfg *Config) HandleEnemyAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.E.EnemyAbilities)
}

func (cfg *Config) HandleAeonCommands(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.AeonCommands)
}

func (cfg *Config) HandleOverdriveCommands(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.OverdriveCommands)
}

func (cfg *Config) HandleOverdrives(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Overdrives)
}

func (cfg *Config) HandleRonsoRages(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.RonsoRages)
}

func (cfg *Config) HandleSubmenus(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Submenus)
}

func (cfg *Config) HandleTopmenus(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Topmenus)
}

func (cfg *Config) HandleAllItems(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.AllItems)
}

func (cfg *Config) HandleItems(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Items)
}

func (cfg *Config) HandleKeyItems(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.KeyItems)
}

func (cfg *Config) HandleSpheres(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Spheres)
}

func (cfg *Config) HandlePrimers(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Primers)
}

func (cfg *Config) HandleMixes(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Mixes)
}

func (cfg *Config) HandleAutoAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.AutoAbilities)
}

func (cfg *Config) HandleEquipmentTables(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(cfg, w, r, cfg.E.EquipmentTables)
}

func (cfg *Config) HandleEquipment(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Equipment)
}

func (cfg *Config) HandleCelestialWeapons(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.CelestialWeapons)
}

func (cfg *Config) HandleStats(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Stats)
}

func (cfg *Config) HandleProperties(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Properties)
}

func (cfg *Config) HandleOverdriveModes(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.OverdriveModes)
}

func (cfg *Config) HandleElements(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Elements)
}

func (cfg *Config) HandleStatusConditions(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.StatusConditions)
}

func (cfg *Config) HandleModifiers(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.E.Modifiers)
}

func (cfg *Config) HandleAgilityTiers(w http.ResponseWriter, r *http.Request) {
	routerIDOnly(cfg, w, r, cfg.E.AgilityTiers)
}

func (cfg *Config) HandleEnums(w http.ResponseWriter, r *http.Request) {
	routerEnums(cfg, w, r, cfg.E.Enums)
}

func (cfg *Config) HandleEndpoints(w http.ResponseWriter, r *http.Request) {
	routerEndpoints(cfg, w, r, cfg.E.Endpoints)
}

func (cfg *Config) HandleAlBhedGet(w http.ResponseWriter, r *http.Request) {
	routerServiceGet(cfg, w, r, cfg.E.AlBhed)
}

func (cfg *Config) HandleAlBhedPost(w http.ResponseWriter, r *http.Request) {
	routerServicePost(cfg, w, r, cfg.E.AlBhed)
}

func (cfg *Config) HandleTurnOrderGet(w http.ResponseWriter, r *http.Request) {
	routerServiceGet(cfg, w, r, cfg.E.TurnOrder)
}

func (cfg *Config) HandleTurnOrderPost(w http.ResponseWriter, r *http.Request) {
	routerServicePost(cfg, w, r, cfg.E.TurnOrder)
}

func (cfg *Config) HandleDropChanceGet(w http.ResponseWriter, r *http.Request) {
	routerServiceGet(cfg, w, r, cfg.E.DropChance)
}

func (cfg *Config) HandleDropChancePost(w http.ResponseWriter, r *http.Request) {
	routerServicePost(cfg, w, r, cfg.E.DropChance)
}

func (cfg *Config) HandleCritChanceGet(w http.ResponseWriter, r *http.Request) {
	routerServiceGet(cfg, w, r, cfg.E.CritChance)
}

func (cfg *Config) HandleCritChancePost(w http.ResponseWriter, r *http.Request) {
	routerServicePost(cfg, w, r, cfg.E.CritChance)
}
