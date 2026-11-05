package main

import (
	"log"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
)

func main() {
	const port = "8080"

	err := config.LoadEnvFromRoot()
	if err != nil {
		log.Fatal(err)
	}

	var server api.Server
	err = server.Init()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/healthz", server.HandlerReadiness)
	mux.HandleFunc("POST /admin/reset", server.HandlerResetDatabase)

	mux.HandleFunc("GET /api/locations/", server.HandleLocations)
	mux.HandleFunc("GET /api/sublocations/", server.HandleSublocations)
	mux.HandleFunc("GET /api/areas/", server.HandleAreas)

	mux.HandleFunc("GET /api/monster-formations/", server.HandleMonsterFormations)
	mux.HandleFunc("GET /api/shops/", server.HandleShops)
	mux.HandleFunc("GET /api/treasures/", server.HandleTreasures)
	mux.HandleFunc("GET /api/quests/", server.HandleQuests)
	mux.HandleFunc("GET /api/sidequests/", server.HandleSidequests)
	mux.HandleFunc("GET /api/subquests/", server.HandleSubquests)
	mux.HandleFunc("GET /api/arena-creations/", server.HandleArenaCreations)
	mux.HandleFunc("GET /api/blitzball-prizes/", server.HandleBlitzballPrizes)
	mux.HandleFunc("GET /api/fmvs/", server.HandleFMVs)
	mux.HandleFunc("GET /api/songs/", server.HandleSongs)

	mux.HandleFunc("GET /api/player-units/", server.HandlePlayerUnits)
	mux.HandleFunc("GET /api/characters/", server.HandleCharacters)
	mux.HandleFunc("GET /api/aeons/", server.HandleAeons)
	mux.HandleFunc("GET /api/character-classes/", server.HandleCharacterClasses)
	mux.HandleFunc("GET /api/monsters/", server.HandleMonsters)

	mux.HandleFunc("GET /api/abilities/", server.HandleAbilities)
	mux.HandleFunc("GET /api/player-abilities/", server.HandlePlayerAbilities)
	mux.HandleFunc("GET /api/overdrive-abilities/", server.HandleOverdriveAbilities)
	mux.HandleFunc("GET /api/item-abilities/", server.HandleItemAbilities)
	mux.HandleFunc("GET /api/trigger-commands/", server.HandleTriggerCommands)
	mux.HandleFunc("GET /api/misc-abilities/", server.HandleMiscAbilities)
	mux.HandleFunc("GET /api/enemy-abilities/", server.HandleEnemyAbilities)

	mux.HandleFunc("GET /api/aeon-commands/", server.HandleAeonCommands)
	mux.HandleFunc("GET /api/overdrive-commands/", server.HandleOverdriveCommands)
	mux.HandleFunc("GET /api/overdrives/", server.HandleOverdrives)
	mux.HandleFunc("GET /api/ronso-rages/", server.HandleRonsoRages)
	mux.HandleFunc("GET /api/submenus/", server.HandleSubmenus)
	mux.HandleFunc("GET /api/topmenus/", server.HandleTopmenus)

	mux.HandleFunc("GET /api/all-items/", server.HandleAllItems)
	mux.HandleFunc("GET /api/items/", server.HandleItems)
	mux.HandleFunc("GET /api/key-items/", server.HandleKeyItems)
	mux.HandleFunc("GET /api/spheres/", server.HandleSpheres)
	mux.HandleFunc("GET /api/primers/", server.HandlePrimers)
	mux.HandleFunc("GET /api/mixes/", server.HandleMixes)

	mux.HandleFunc("GET /api/auto-abilities/", server.HandleAutoAbilities)
	mux.HandleFunc("GET /api/equipment-tables/", server.HandleEquipmentTables)
	mux.HandleFunc("GET /api/equipment/", server.HandleEquipment)
	mux.HandleFunc("GET /api/celestial-weapons/", server.HandleCelestialWeapons)

	mux.HandleFunc("GET /api/stats/", server.HandleStats)
	mux.HandleFunc("GET /api/properties/", server.HandleProperties)
	mux.HandleFunc("GET /api/overdrive-modes/", server.HandleOverdriveModes)
	mux.HandleFunc("GET /api/elements/", server.HandleElements)
	mux.HandleFunc("GET /api/status-conditions/", server.HandleStatusConditions)
	mux.HandleFunc("GET /api/modifiers/", server.HandleModifiers)
	mux.HandleFunc("GET /api/agility-tiers/", server.HandleAgilityTiers)

	mux.HandleFunc("GET /api/enums/", server.HandleEnums)
	mux.HandleFunc("GET /api/endpoints/", server.HandleEndpoints)

	mux.HandleFunc("GET /api/al-bhed", server.HandleAlBhedGet)
	mux.HandleFunc("GET /api/al-bhed/", server.HandleAlBhedGet)
	mux.HandleFunc("POST /api/al-bhed", server.HandleAlBhedPost)
	mux.HandleFunc("POST /api/al-bhed/", server.HandleAlBhedPost)

	mux.HandleFunc("GET /api/turn-order", server.HandleTurnOrderGet)
	mux.HandleFunc("GET /api/turn-order/", server.HandleTurnOrderGet)
	mux.HandleFunc("POST /api/turn-order", server.HandleTurnOrderPost)
	mux.HandleFunc("POST /api/turn-order/", server.HandleTurnOrderPost)

	mux.HandleFunc("GET /api/drop-chance", server.HandleDropChanceGet)
	mux.HandleFunc("GET /api/drop-chance/", server.HandleDropChanceGet)
	mux.HandleFunc("POST /api/drop-chance", server.HandleDropChancePost)
	mux.HandleFunc("POST /api/drop-chance/", server.HandleDropChancePost)

	mux.HandleFunc("GET /api/crit-chance", server.HandleCritChanceGet)
	mux.HandleFunc("GET /api/crit-chance/", server.HandleCritChanceGet)
	mux.HandleFunc("POST /api/crit-chance", server.HandleCritChancePost)
	mux.HandleFunc("POST /api/crit-chance/", server.HandleCritChancePost)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
