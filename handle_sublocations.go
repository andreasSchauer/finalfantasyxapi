package main

import (
	"fmt"
	"net/http"
	
	//"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// a lot of fields can be embedded from areas (make a shared struct)
// maybe make sidequest a slice, even for areas to keep the structure consistent, since they're so similar
// for booleans, if one of the areas is true, the boolean is true

func (cfg *Config) HandleSublocations(w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, "sublocations")

	// this whole thing can probably be generalized
	switch len(segments) {
	case 0:
		// /api/sublocations
		fmt.Println(segments)
		fmt.Println("this should trigger /api/sublocations")
		return
	case 1:
		// /api/sublocations/{name or id}
		fmt.Println(segments)
		fmt.Println("this should trigger /api/sublocations/{name or id}")
		return

	case 2:
		// /api/sublocations/{id}/{subSection}

		// sublocationID := segments[0]
		subSection := segments[1]
		switch subSection {
		case "areas":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/sublocations/{name or id}/areas")
			return
		case "monsters":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/sublocations/{name or id}/monsters")
			return
		case "monster-formations":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/sublocations/{name or id}/monster-formations")
			return
		case "shops":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/sublocations/{name or id}/shops")
			return
		case "treasures":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/sublocations/{name or id}/treasures")
			return
		default:
			fmt.Println(segments)
			fmt.Println("this should trigger an error: this sub section is not supported. Supported sub-sections: areas, monsters, monster-formations, shops, treasures.")
			return
		}

	default:
		respondWithError(w, http.StatusBadRequest, `wrong format. usage: '/api/sublocations/{name or id}', or '/api/sublocations/{name or id}/{sub-section}'. supported sub-sections: areas, monsters, monster-formations, shops, treasures.`, nil)
		return
	}
}
