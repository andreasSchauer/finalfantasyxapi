package api


type Primer struct {
	ID                 	int32                	`json:"id"`
	Name              	string               	`json:"name"`
	KeyItem				NamedAPIResource		`json:"key_item"`
	Description        	string               	`json:"description"`
	AlBhedLetter		string					`json:"al_bhed_letter"`
	EnglishLetter		string					`json:"english_letter"`
	Areas				[]AreaAPIResource		`json:"areas"`
	Treasures          	[]UnnamedAPIResource 	`json:"treasures"`
}