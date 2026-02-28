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

	mux.HandleFunc("GET /api/areas/", apiCfg.HandleAreas)
	mux.HandleFunc("GET /api/sublocations/", apiCfg.HandleSublocations)
	mux.HandleFunc("GET /api/locations/", apiCfg.HandleLocations)
	
	mux.HandleFunc("GET /api/aeons/", apiCfg.HandleAeons)
	mux.HandleFunc("GET /api/arena-creations/", apiCfg.HandleArenaCreations)
	mux.HandleFunc("GET /api/blitzball-prizes/", apiCfg.HandleBlitzballPrizes)
	mux.HandleFunc("GET /api/character-classes/", apiCfg.HandleCharacterClasses)
	mux.HandleFunc("GET /api/characters/", apiCfg.HandleCharacters)
	mux.HandleFunc("GET /api/fmvs/", apiCfg.HandleFMVs)
	mux.HandleFunc("GET /api/monsters/", apiCfg.HandleMonsters)
	mux.HandleFunc("GET /api/monster-formations/", apiCfg.HandleMonsterFormations)
	mux.HandleFunc("GET /api/overdrive-modes/", apiCfg.HandleOverdriveModes)
	mux.HandleFunc("GET /api/enemy-abilities/", apiCfg.HandleEnemyAbilities)
	mux.HandleFunc("GET /api/item-abilities/", apiCfg.HandleItemAbilities)
	mux.HandleFunc("GET /api/other-abilities/", apiCfg.HandleOtherAbilities)
	mux.HandleFunc("GET /api/overdrive-abilities/", apiCfg.HandleOverdriveAbilities)
	mux.HandleFunc("GET /api/player-abilities/", apiCfg.HandlePlayerAbilities)
	mux.HandleFunc("GET /api/trigger-commands/", apiCfg.HandleTriggerCommands)
	mux.HandleFunc("GET /api/shops/", apiCfg.HandleShops)
	mux.HandleFunc("GET /api/songs/", apiCfg.HandleSongs)
	mux.HandleFunc("GET /api/sidequests/", apiCfg.HandleSidequests)
	mux.HandleFunc("GET /api/subquests/", apiCfg.HandleSubquests)
	mux.HandleFunc("GET /api/treasures/", apiCfg.HandleTreasures)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}