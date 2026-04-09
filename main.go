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

	apiCfg, err := api.ConfigInit()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/healthz", apiCfg.HandlerReadiness)
	mux.HandleFunc("POST /admin/reset", apiCfg.HandlerResetDatabase)

	mux.HandleFunc("GET /api/locations/", apiCfg.HandleLocations)
	mux.HandleFunc("GET /api/sublocations/", apiCfg.HandleSublocations)
	mux.HandleFunc("GET /api/areas/", apiCfg.HandleAreas)
	
	mux.HandleFunc("GET /api/monster-formations/", apiCfg.HandleMonsterFormations)
	mux.HandleFunc("GET /api/shops/", apiCfg.HandleShops)
	mux.HandleFunc("GET /api/treasures/", apiCfg.HandleTreasures)
	mux.HandleFunc("GET /api/quests/", apiCfg.HandleQuests)
	mux.HandleFunc("GET /api/sidequests/", apiCfg.HandleSidequests)
	mux.HandleFunc("GET /api/subquests/", apiCfg.HandleSubquests)
	mux.HandleFunc("GET /api/arena-creations/", apiCfg.HandleArenaCreations)
	mux.HandleFunc("GET /api/blitzball-prizes/", apiCfg.HandleBlitzballPrizes)
	mux.HandleFunc("GET /api/fmvs/", apiCfg.HandleFMVs)
	mux.HandleFunc("GET /api/songs/", apiCfg.HandleSongs)
	
	mux.HandleFunc("GET /api/player-units/", apiCfg.HandlePlayerUnits)
	mux.HandleFunc("GET /api/characters/", apiCfg.HandleCharacters)
	mux.HandleFunc("GET /api/aeons/", apiCfg.HandleAeons)
	mux.HandleFunc("GET /api/character-classes/", apiCfg.HandleCharacterClasses)
	mux.HandleFunc("GET /api/monsters/", apiCfg.HandleMonsters)
	
	mux.HandleFunc("GET /api/abilities/", apiCfg.HandleAbilities)
	mux.HandleFunc("GET /api/player-abilities/", apiCfg.HandlePlayerAbilities)
	mux.HandleFunc("GET /api/overdrive-abilities/", apiCfg.HandleOverdriveAbilities)
	mux.HandleFunc("GET /api/item-abilities/", apiCfg.HandleItemAbilities)
	mux.HandleFunc("GET /api/trigger-commands/", apiCfg.HandleTriggerCommands)
	mux.HandleFunc("GET /api/unspecified-abilities/", apiCfg.HandleUnspecifiedAbilities)
	mux.HandleFunc("GET /api/enemy-abilities/", apiCfg.HandleEnemyAbilities)

	mux.HandleFunc("GET /api/aeon-commands/", apiCfg.HandleAeonCommands)
	mux.HandleFunc("GET /api/overdrive-commands/", apiCfg.HandleOverdriveCommands)
	mux.HandleFunc("GET /api/overdrives/", apiCfg.HandleOverdrives)
	mux.HandleFunc("GET /api/ronso-rages/", apiCfg.HandleRonsoRages)
	mux.HandleFunc("GET /api/submenus/", apiCfg.HandleSubmenus)
	mux.HandleFunc("GET /api/topmenus/", apiCfg.HandleTopmenus)
	
	mux.HandleFunc("GET /api/all-items/", apiCfg.HandleAllItems)
	mux.HandleFunc("GET /api/items/", apiCfg.HandleItems)
	mux.HandleFunc("GET /api/key-items/", apiCfg.HandleKeyItems)
	mux.HandleFunc("GET /api/spheres/", apiCfg.HandleSpheres)
	mux.HandleFunc("GET /api/primers/", apiCfg.HandlePrimers)
	mux.HandleFunc("GET /api/mixes/", apiCfg.HandleMixes)
	
	mux.HandleFunc("GET /api/auto-abilities/", apiCfg.HandleAutoAbilities)
	mux.HandleFunc("GET /api/equipment-tables/", apiCfg.HandleEquipmentTables)

	mux.HandleFunc("GET /api/overdrive-modes/", apiCfg.HandleOverdriveModes)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
