package main

import (
	"log"
	"net/http"
)



func main() {
	//const filepathRoot = "."
	const port = "8080"

	apiCfg, err := apiConfigInit()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/healthz", apiCfg.handlerReadiness)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerResetDatabase)

	mux.HandleFunc("GET /api/locations/", apiCfg.handleLocations)
	mux.HandleFunc("GET /api/sublocations/", apiCfg.handleSublocations)
	mux.HandleFunc("GET /api/areas/", apiCfg.handleAreas)
	mux.HandleFunc("GET /api/monsters/", apiCfg.handleMonsters)
	mux.HandleFunc("GET /api/overdrive-modes/", apiCfg.handleOverdriveModes)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}