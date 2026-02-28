package api

import "net/http"


func (cfg *Config) HandleAeons(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.aeons)
}


func (cfg *Config) HandleAreas(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.areas)
}


func (cfg *Config) HandleArenaCreations(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.arenaCreations)
}


func (cfg *Config) HandleBlitzballPrizes(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.blitzballPrizes)
}

func (cfg *Config) HandleCharacterClasses(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.characterClasses)
}

func (cfg *Config) HandleCharacters(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.characters)
}


func (cfg *Config) HandleFMVs(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.fmvs)
}


func (cfg *Config) HandleLocations(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.locations)
}


func (cfg *Config) HandleMonsterFormations(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.monsterFormations)
}


func (cfg *Config) HandleMonsters(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.e.monsters)
}


func (cfg *Config) HandleOverdriveModes(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.overdriveModes)
}


func (cfg *Config) HandleEnemyAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.e.enemyAbilities)
}


func (cfg *Config) HandleItemAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.itemAbilities)
}


func (cfg *Config) HandleOtherAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.e.otherAbilities)
}


func (cfg *Config) HandleOverdriveAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.e.overdriveAbilities)
}


func (cfg *Config) HandlePlayerAbilities(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.e.playerAbilities)
}


func (cfg *Config) HandleTriggerCommands(w http.ResponseWriter, r *http.Request) {
	routerNameVersion(cfg, w, r, cfg.e.triggerCommands)
}


func (cfg *Config) HandleShops(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.shops)
}


func (cfg *Config) HandleSidequests(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.sidequests)
}


func (cfg *Config) HandleSongs(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.songs)
}


func (cfg *Config) HandleSublocations(w http.ResponseWriter, r *http.Request) {
	routerNameOrID(cfg, w, r, cfg.e.sublocations)
}


func (cfg *Config) HandleSubquests(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.subquests)
}


func (cfg *Config) HandleTreasures(w http.ResponseWriter, r *http.Request) {
	routerIdOnly(cfg, w, r, cfg.e.treasures)
}