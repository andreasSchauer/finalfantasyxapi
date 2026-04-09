package api

import "net/http"

func (cfg *Config) HandleLocations(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.locations)
}

func (cfg *Config) HandleSublocations(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.sublocations)
}

func (cfg *Config) HandleAreas(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.areas)
}



func (cfg *Config) HandleMonsterFormations(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.monsterFormations)
}

func (cfg *Config) HandleShops(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.shops)
}

func (cfg *Config) HandleTreasures(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.treasures)
}

func (cfg *Config) HandleQuests(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.quests)
}

func (cfg *Config) HandleSidequests(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.sidequests)
}

func (cfg *Config) HandleSubquests(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.subquests)
}

func (cfg *Config) HandleArenaCreations(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.arenaCreations)
}

func (cfg *Config) HandleBlitzballPrizes(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.blitzballPrizes)
}

func (cfg *Config) HandleSongs(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.songs)
}

func (cfg *Config) HandleFMVs(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.fmvs)
}



func (cfg *Config) HandlePlayerUnits(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.playerUnits)
}

func (cfg *Config) HandleCharacters(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.characters)
}

func (cfg *Config) HandleAeons(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.aeons)
}

func (cfg *Config) HandleCharacterClasses(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.characterClasses)
}

func (cfg *Config) HandleMonsters(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.e.monsters)
}



func (cfg *Config) HandleAbilities(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.abilities)
}

func (cfg *Config) HandlePlayerAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.e.playerAbilities)
}

func (cfg *Config) HandleOverdriveAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.e.overdriveAbilities)
}

func (cfg *Config) HandleItemAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.itemAbilities)
}

func (cfg *Config) HandleTriggerCommands(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.e.triggerCommands)
}

func (cfg *Config) HandleUnspecifiedAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.e.unspecifiedAbilities)
}

func (cfg *Config) HandleEnemyAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.e.enemyAbilities)
}



func (cfg *Config) HandleAeonCommands(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.aeonCommands)
}

func (cfg *Config) HandleOverdriveCommands(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.overdriveCommands)
}

func (cfg *Config) HandleOverdrives(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.overdrives)
}

func (cfg *Config) HandleRonsoRages(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.ronsoRages)
}

func (cfg *Config) HandleSubmenus(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.submenus)
}

func (cfg *Config) HandleTopmenus(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.topmenus)
}




func (cfg *Config) HandleAllItems(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.allItems)
}

func (cfg *Config) HandleItems(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.items)
}

func (cfg *Config) HandleKeyItems(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.keyItems)
}

func (cfg *Config) HandleSpheres(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.spheres)
}

func (cfg *Config) HandlePrimers(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.primers)
}

func (cfg *Config) HandleMixes(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.mixes)
}




func (cfg *Config) HandleAutoAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.autoAbilities)
}

func (cfg *Config) HandleEquipmentTables(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.equipmentTables)
}



func (cfg *Config) HandleOverdriveModes(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.overdriveModes)
}