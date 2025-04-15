package models

type LocationArea struct {
	Name string `json:"name"`
}

type LocationAreaResponse struct {
	Results []LocationArea `json:"results"`
}

type Stat struct {
	Name string `json:"name"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Pokemon struct {
	Name    string  `json:"name"`
	BaseExp int     `json:"base_experience"`
	Stats   []Stats `json:"stats"`
}

type Encounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type AreaPokemonInfo struct {
	Encounters []Encounter `json:"pokemon_encounters"`
}
