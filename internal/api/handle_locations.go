package api

import (
	"fmt"
	"net/http"
	//"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (cfg *Config) HandleLocations(w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, "locations")

	// this whole thing can probably be generalized
	switch len(segments) {
	case 0:
		// /api/locations
		fmt.Println(segments)
		fmt.Println("this should trigger /api/locations")
		return
	case 1:
		// /api/locations/{name or id}
		fmt.Println(segments)
		fmt.Println("this should trigger /api/locations/{name or id}")
		return

	case 2:
		// /api/locations/{id}/{subSection}
		// locationID := segments[0]
		subSection := segments[1]
		switch subSection {
		case "sublocations":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/locations/{name or id}/sublocations")
			return
		case "areas":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/locations/{name or id}/areas")
			return
		case "monsters":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/locations/{name or id}/monsters")
			return
		case "monster-formations":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/locations/{name or id}/monster-formations")
			return
		case "shops":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/locations/{name or id}/shops")
			return
		case "treasures":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/locations/{name or id}/treasures")
			return
		default:
			fmt.Println(segments)
			fmt.Println("this should trigger an error: this sub section is not supported. Supported sub-sections: sublocations, areas, monsters, monster-formations, shops, treasures.")
			return
		}

	default:
		respondWithError(w, http.StatusBadRequest, `Wrong format. Usage: /api/sublocations/{name or id}, or /api/sublocations/{name or id}/{sub-section}. Supported sub-sections: sublocations, areas, monsters, monster-formations, shops, treasures.`, nil)
		return
	}
}
