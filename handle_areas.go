package main

import (
	"fmt"
	"net/http"

	//"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (cfg *apiConfig) handleAreas(w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, "areas")
	
	// this whole thing can probably be generalized
	switch len(segments) {
	case 0:
		// /api/areas
		fmt.Println(segments)
		fmt.Println("this should trigger /api/areas")
		return
	case 1:
		// /api/areas/{id}
		fmt.Println(segments)
		fmt.Println("this should trigger /api/areas/{id}")
		return

	case 2:
		// /api/areas/{id}/{subSection}
		// areaID := segments[0]
		subSection := segments[1]
		switch subSection {
		case "treasures":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/treasures")
			return
		case "shops":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/shops")
			return
		case "monster-formations":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/monster-formations")
			return
		case "monsters":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/monsters")
			return
		default:
			fmt.Println(segments)
			fmt.Println("this should trigger an error: this sub section is not supported. Supported sub-sections: monsters, monster-formations, shops, treasures.")
			return
		}

	default:
		respondWithError(w, http.StatusBadRequest, `Wrong format. Usage: /api/areas/{id}, or /api/areas/{id}/{sub-section}. Supported sub-sections: monsters, monster-formations, shops, treasures.`, nil)
		return
	}
}
