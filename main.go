package main

import (
	"log"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
)



func main() {
	//const filepathRoot = "."
	const port = "8080"

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
	mux.HandleFunc("GET /api/monsters/", apiCfg.HandleMonsters)
	mux.HandleFunc("GET /api/overdrive-modes/", apiCfg.HandleOverdriveModes)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}