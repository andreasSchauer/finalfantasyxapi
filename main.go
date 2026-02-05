package main

import (
	"log"
	"net/http"
)



func main() {
	//const filepathRoot = "."
	const port = "8080"

	apiCfg, err := ConfigInit()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/healthz", apiCfg.HandlerReadiness)
	mux.HandleFunc("POST /admin/reset", apiCfg.HandlerResetDatabase)

	mux.HandleFunc("GET /api/areas/", apiCfg.HandleAreas)
	mux.HandleFunc("GET /api/sublocations/", apiCfg.HandleSublocations)
	mux.HandleFunc("GET /api/locations/", apiCfg.HandleLocations)
	
	mux.HandleFunc("GET /api/arena-creations/", apiCfg.HandleArenaCreations)
	mux.HandleFunc("GET /api/fmvs/", apiCfg.HandleFMVs)
	mux.HandleFunc("GET /api/monsters/", apiCfg.HandleMonsters)
	mux.HandleFunc("GET /api/monster-formations/", apiCfg.HandleMonsterFormations)
	mux.HandleFunc("GET /api/overdrive-modes/", apiCfg.HandleOverdriveModes)
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