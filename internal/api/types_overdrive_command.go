package api

type OverdriveCommand struct {
	ID          int32				`json:"id"`
	Name        string 				`json:"name"`
	Description string 				`json:"description"`
	Rank        int32  				`json:"rank"`
	User        NamedAPIResource 	`json:"user"`
	Topmenu     *NamedAPIResource 	`json:"topmenu"`
	OpenSubmenu NamedAPIResource 	`json:"open_submenu"`
	Overdrives	[]NamedAPIResource	`json:"overdrives"`
}