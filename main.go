package main

import (
	"log"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
)

func main() {
	const port = "8080"

	err := api.LoadEnvFromRoot()
	if err != nil {
		log.Fatal(err)
	}

	var server api.Server
	err = server.Init()
	if err != nil {
		log.Fatal(err)
	}

	server.InitGlobals()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/healthz", server.Cfg.HandlerReadiness)
	mux.HandleFunc("POST /admin/reset", server.Cfg.HandlerResetDatabase)

	mux.HandleFunc("GET /api/locations/", server.Cfg.HandleLocations)
	mux.HandleFunc("GET /api/sublocations/", server.Cfg.HandleSublocations)
	mux.HandleFunc("GET /api/areas/", server.Cfg.HandleAreas)

	mux.HandleFunc("GET /api/monster-formations/", server.Cfg.HandleMonsterFormations)
	mux.HandleFunc("GET /api/shops/", server.Cfg.HandleShops)
	mux.HandleFunc("GET /api/treasures/", server.Cfg.HandleTreasures)
	mux.HandleFunc("GET /api/quests/", server.Cfg.HandleQuests)
	mux.HandleFunc("GET /api/sidequests/", server.Cfg.HandleSidequests)
	mux.HandleFunc("GET /api/subquests/", server.Cfg.HandleSubquests)
	mux.HandleFunc("GET /api/arena-creations/", server.Cfg.HandleArenaCreations)
	mux.HandleFunc("GET /api/blitzball-prizes/", server.Cfg.HandleBlitzballPrizes)
	mux.HandleFunc("GET /api/fmvs/", server.Cfg.HandleFMVs)
	mux.HandleFunc("GET /api/songs/", server.Cfg.HandleSongs)

	mux.HandleFunc("GET /api/player-units/", server.Cfg.HandlePlayerUnits)
	mux.HandleFunc("GET /api/characters/", server.Cfg.HandleCharacters)
	mux.HandleFunc("GET /api/aeons/", server.Cfg.HandleAeons)
	mux.HandleFunc("GET /api/character-classes/", server.Cfg.HandleCharacterClasses)
	mux.HandleFunc("GET /api/monsters/", server.Cfg.HandleMonsters)

	mux.HandleFunc("GET /api/abilities/", server.Cfg.HandleAbilities)
	mux.HandleFunc("GET /api/player-abilities/", server.Cfg.HandlePlayerAbilities)
	mux.HandleFunc("GET /api/overdrive-abilities/", server.Cfg.HandleOverdriveAbilities)
	mux.HandleFunc("GET /api/item-abilities/", server.Cfg.HandleItemAbilities)
	mux.HandleFunc("GET /api/trigger-commands/", server.Cfg.HandleTriggerCommands)
	mux.HandleFunc("GET /api/misc-abilities/", server.Cfg.HandleMiscAbilities)
	mux.HandleFunc("GET /api/enemy-abilities/", server.Cfg.HandleEnemyAbilities)

	mux.HandleFunc("GET /api/aeon-commands/", server.Cfg.HandleAeonCommands)
	mux.HandleFunc("GET /api/overdrive-commands/", server.Cfg.HandleOverdriveCommands)
	mux.HandleFunc("GET /api/overdrives/", server.Cfg.HandleOverdrives)
	mux.HandleFunc("GET /api/ronso-rages/", server.Cfg.HandleRonsoRages)
	mux.HandleFunc("GET /api/submenus/", server.Cfg.HandleSubmenus)
	mux.HandleFunc("GET /api/topmenus/", server.Cfg.HandleTopmenus)

	mux.HandleFunc("GET /api/all-items/", server.Cfg.HandleAllItems)
	mux.HandleFunc("GET /api/items/", server.Cfg.HandleItems)
	mux.HandleFunc("GET /api/key-items/", server.Cfg.HandleKeyItems)
	mux.HandleFunc("GET /api/spheres/", server.Cfg.HandleSpheres)
	mux.HandleFunc("GET /api/primers/", server.Cfg.HandlePrimers)
	mux.HandleFunc("GET /api/mixes/", server.Cfg.HandleMixes)

	mux.HandleFunc("GET /api/auto-abilities/", server.Cfg.HandleAutoAbilities)
	mux.HandleFunc("GET /api/equipment-tables/", server.Cfg.HandleEquipmentTables)
	mux.HandleFunc("GET /api/equipment/", server.Cfg.HandleEquipment)
	mux.HandleFunc("GET /api/celestial-weapons/", server.Cfg.HandleCelestialWeapons)

	mux.HandleFunc("GET /api/stats/", server.Cfg.HandleStats)
	mux.HandleFunc("GET /api/properties/", server.Cfg.HandleProperties)
	mux.HandleFunc("GET /api/overdrive-modes/", server.Cfg.HandleOverdriveModes)
	mux.HandleFunc("GET /api/elements/", server.Cfg.HandleElements)
	mux.HandleFunc("GET /api/status-conditions/", server.Cfg.HandleStatusConditions)
	mux.HandleFunc("GET /api/modifiers/", server.Cfg.HandleModifiers)
	mux.HandleFunc("GET /api/agility-tiers/", server.Cfg.HandleAgilityTiers)

	mux.HandleFunc("GET /api/enums/", server.Cfg.HandleEnums)
	mux.HandleFunc("GET /api/endpoints/", server.Cfg.HandleEndpoints)

	mux.HandleFunc("GET /api/al-bhed", server.Cfg.HandleAlBhedGet)
	mux.HandleFunc("GET /api/al-bhed/", server.Cfg.HandleAlBhedGet)
	mux.HandleFunc("POST /api/al-bhed", server.Cfg.HandleAlBhedPost)
	mux.HandleFunc("POST /api/al-bhed/", server.Cfg.HandleAlBhedPost)

	mux.HandleFunc("GET /api/turn-order", server.Cfg.HandleTurnOrderGet)
	mux.HandleFunc("GET /api/turn-order/", server.Cfg.HandleTurnOrderGet)
	mux.HandleFunc("POST /api/turn-order", server.Cfg.HandleTurnOrderPost)
	mux.HandleFunc("POST /api/turn-order/", server.Cfg.HandleTurnOrderPost)

	mux.HandleFunc("GET /api/drop-chance", server.Cfg.HandleDropChanceGet)
	mux.HandleFunc("GET /api/drop-chance/", server.Cfg.HandleDropChanceGet)
	mux.HandleFunc("POST /api/drop-chance", server.Cfg.HandleDropChancePost)
	mux.HandleFunc("POST /api/drop-chance/", server.Cfg.HandleDropChancePost)

	mux.HandleFunc("GET /api/crit-chance", server.Cfg.HandleCritChanceGet)
	mux.HandleFunc("GET /api/crit-chance/", server.Cfg.HandleCritChanceGet)
	mux.HandleFunc("POST /api/crit-chance", server.Cfg.HandleCritChancePost)
	mux.HandleFunc("POST /api/crit-chance/", server.Cfg.HandleCritChancePost)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
