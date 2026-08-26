package pokeapi

type Pokemon struct {
	Name            string `json:"name"`
	Base_Experience int    `json:"base_experience"`
}

type RespLocationAreasStruct struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type RespPokemonInLocationAreas struct {
	Pokemon_Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
