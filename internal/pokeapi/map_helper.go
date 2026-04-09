package pokeapi

type ResourceReference struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterMethodVersionDetail struct {
	Rate    int               `json:"rate"`
	Version ResourceReference `json:"version"`
}

type EncounterMethodRate struct {
	EncounterMethod ResourceReference              `json:"encounter_method"`
	VersionDetails  []EncounterMethodVersionDetail `json:"version_details"`
}

type NameEntry struct {
	Name     string            `json:"name"`
	Language ResourceReference `json:"language"`
}

type EncounterDetail struct {
	MinLevel        int               `json:"min_level"`
	MaxLevel        int               `json:"max_level"`
	ConditionValues []any             `json:"condition_values"`
	Chance          int               `json:"chance"`
	Method          ResourceReference `json:"method"`
}

type PokemonEncounterVersionDetail struct {
	Version          ResourceReference `json:"version"`
	MaxChance        int               `json:"max_chance"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
}

type PokemonEncounter struct {
	Pokemon        ResourceReference               `json:"pokemon"`
	VersionDetails []PokemonEncounterVersionDetail `json:"version_details"`
}

type LocationArea struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             ResourceReference     `json:"location"`
	Names                []NameEntry           `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type ShallowLocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
