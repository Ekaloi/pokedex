package pokeapi

import "github.com/Ekaloi/pokedex/poketypes"

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterMethodRate struct {
	EncounterMethod EncounterMethod `json:"encounter_method"`
	VersionDetails  []poketypes.VersionDetail `json:"version_details"`
}


type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Name struct {
	Name     string   `json:"name"`
	Language Language `json:"language"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterDetail struct {
	MinLevel        int        `json:"min_level"`
	MaxLevel        int        `json:"max_level"`
	ConditionValues []struct{} `json:"condition_values"`
	Chance          int        `json:"chance"`
	Method          EncounterMethod `json:"method"`
}

type PokemonEncounter struct {
	Pokemon        Pokemon         `json:"pokemon"`
	VersionDetails []VersionChance `json:"version_details"`
}

type VersionChance struct {
	Version         Version          `json:"version"`
	MaxChance       int              `json:"max_chance"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
}

type AreaData struct {
	ID                  int                   `json:"id"`
	Name                string                `json:"name"`
	GameIndex           int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location            Location              `json:"location"`
	Names               []Name                `json:"names"`
	PokemonEncounters   []PokemonEncounter    `json:"pokemon_encounters"`
}