package pokeapi

type PokemonActual struct {
	ID                     int              `json:"id"`
	Name                   string           `json:"name"`
	BaseExperience         int              `json:"base_experience"`
	Height                 int              `json:"height"`
	IsDefault              bool             `json:"is_default"`
	Order                  int              `json:"order"`
	Weight                 int              `json:"weight"`
	Abilities              []Ability        `json:"abilities"`
	Forms                  []NamedAPIResource `json:"forms"`
	GameIndices            []GameIndex      `json:"game_indices"`
	HeldItems              []HeldItem       `json:"held_items"`
	LocationAreaEncounters string           `json:"location_area_encounters"`
	Moves                  []Move           `json:"moves"`
	Species                NamedAPIResource `json:"species"`
	Sprites                Sprites          `json:"sprites"`
	Cries                  Cries            `json:"cries"`
	Stats                  []Stat           `json:"stats"`
	Types                  []TypeSlot       `json:"types"`
	PastTypes              []PastType       `json:"past_types"`
}

// Ability holds ability data.
type Ability struct {
	IsHidden bool              `json:"is_hidden"`
	Slot     int               `json:"slot"`
	Ability  NamedAPIResource  `json:"ability"`
}

// NamedAPIResource is a common structure with a name and URL.
type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// GameIndex holds game index data.
type GameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

// HeldItem represents an item the Pokémon may hold.
type HeldItem struct {
	Item           NamedAPIResource `json:"item"`
	VersionDetails []VersionDetail  `json:"version_details"`
}

// VersionDetail holds details for a held item in a given version.
type VersionDetail struct {
	Rarity  int              `json:"rarity"`
	Version NamedAPIResource `json:"version"`
}

// Move represents a move the Pokémon can learn.
type Move struct {
	Move                NamedAPIResource     `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

// VersionGroupDetail holds how and when a move is learned.
type VersionGroupDetail struct {
	LevelLearnedAt  int              `json:"level_learned_at"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
}

// Sprites holds all sprite information.
type Sprites struct {
	BackDefault      string       `json:"back_default"`
	BackFemale       *string      `json:"back_female"`
	BackShiny        string       `json:"back_shiny"`
	BackShinyFemale  *string      `json:"back_shiny_female"`
	FrontDefault     string       `json:"front_default"`
	FrontFemale      *string      `json:"front_female"`
	FrontShiny       string       `json:"front_shiny"`
	FrontShinyFemale *string      `json:"front_shiny_female"`
	Other            OtherSprites `json:"other"`
	Versions         Versions     `json:"versions"`
}

// OtherSprites holds alternative sprite sets.
type OtherSprites struct {
	DreamWorld      DreamWorldSprites   `json:"dream_world"`
	Home            HomeSprites         `json:"home"`
	OfficialArtwork OfficialArtwork     `json:"official-artwork"`
	Showdown        ShowdownSprites     `json:"showdown"`
}

// DreamWorldSprites holds sprites from the dream world.
type DreamWorldSprites struct {
	FrontDefault string  `json:"front_default"`
	FrontFemale  *string `json:"front_female"`
}

// HomeSprites holds sprites from the home set.
type HomeSprites struct {
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// OfficialArtwork holds the official artwork sprites.
type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

// ShowdownSprites holds showdown sprites.
type ShowdownSprites struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// Versions holds sprite versions for various generations.
type Versions struct {
	GenerationI    GenerationISprites    `json:"generation-i"`
	GenerationII   GenerationIISprites   `json:"generation-ii"`
	GenerationIII  GenerationIIISprites  `json:"generation-iii"`
	GenerationIV   GenerationIVSprites   `json:"generation-iv"`
	GenerationV    GenerationVSprites    `json:"generation-v"`
	GenerationVI   GenerationVISprites   `json:"generation-vi"`
	GenerationVII  GenerationVIISprites  `json:"generation-vii"`
	GenerationVIII GenerationVIIISprites `json:"generation-viii"`
}

// GenerationISprites holds Generation I sprite sets.
type GenerationISprites struct {
	RedBlue GenerationIRedBlueSprites `json:"red-blue"`
	Yellow  GenerationIYellowSprites    `json:"yellow"`
}

type GenerationIRedBlueSprites struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}

type GenerationIYellowSprites struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}

// GenerationIISprites holds Generation II sprite sets.
type GenerationIISprites struct {
	Crystal GenerationIICrystalSprites `json:"crystal"`
	Gold    GenerationIIGoldSprites    `json:"gold"`
	Silver  GenerationIISilverSprites  `json:"silver"`
}

type GenerationIICrystalSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type GenerationIIGoldSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type GenerationIISilverSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

// GenerationIIISprites holds Generation III sprite sets.
type GenerationIIISprites struct {
	Emerald          GenerationIIIEmeraldSprites          `json:"emerald"`
	FireredLeafgreen GenerationIIIFireredLeafgreenSprites   `json:"firered-leafgreen"`
	RubySapphire     GenerationIIIRubySapphireSprites      `json:"ruby-sapphire"`
}

type GenerationIIIEmeraldSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type GenerationIIIFireredLeafgreenSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type GenerationIIIRubySapphireSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

// GenerationIVSprites holds Generation IV sprite sets.
type GenerationIVSprites struct {
	DiamondPearl        GenerationIVDPSprites       `json:"diamond-pearl"`
	HeartgoldSoulsilver GenerationIVHSSSprites      `json:"heartgold-soulsilver"`
	Platinum            GenerationIVPlatinumSprites `json:"platinum"`
}

type GenerationIVDPSprites struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type GenerationIVHSSSprites struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type GenerationIVPlatinumSprites struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// GenerationVSprites holds Generation V sprite sets.
type GenerationVSprites struct {
	BlackWhite GenerationVBlackWhiteSprites `json:"black-white"`
}

type GenerationVBlackWhiteSprites struct {
	Animated         GenerationVAnimatedSprites `json:"animated"`
	BackDefault      string                     `json:"back_default"`
	BackFemale       *string                    `json:"back_female"`
	BackShiny        string                     `json:"back_shiny"`
	BackShinyFemale  *string                    `json:"back_shiny_female"`
	FrontDefault     string                     `json:"front_default"`
	FrontFemale      *string                    `json:"front_female"`
	FrontShiny       string                     `json:"front_shiny"`
	FrontShinyFemale *string                    `json:"front_shiny_female"`
}

type GenerationVAnimatedSprites struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// GenerationVISprites holds Generation VI sprite sets.
type GenerationVISprites struct {
	OmegarubyAlphasapphire GenerationVIOASprites `json:"omegaruby-alphasapphire"`
	XY                     GenerationVIXYSprites `json:"x-y"`
}

type GenerationVIOASprites struct {
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type GenerationVIXYSprites struct {
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// GenerationVIISprites holds Generation VII sprite sets.
type GenerationVIISprites struct {
	Icons             GenerationVIIIcons      `json:"icons"`
	UltraSunUltraMoon GenerationVIIUSUMSprites `json:"ultra-sun-ultra-moon"`
}

type GenerationVIIIcons struct {
	FrontDefault string  `json:"front_default"`
	FrontFemale  *string `json:"front_female"`
}

type GenerationVIIUSUMSprites struct {
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// GenerationVIIISprites holds Generation VIII sprite sets.
type GenerationVIIISprites struct {
	Icons GenerationVIIIIconsDetail `json:"icons"`
}

type GenerationVIIIIconsDetail struct {
	FrontDefault string  `json:"front_default"`
	FrontFemale  *string `json:"front_female"`
}

// Cries holds the cry audio URLs.
type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

// Stat represents a stat of the Pokémon.
type Stat struct {
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
	Stat     NamedAPIResource `json:"stat"`
}

// TypeSlot represents a Pokémon type slot.
type TypeSlot struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

// PastType holds types for previous generations.
type PastType struct {
	Generation NamedAPIResource `json:"generation"`
	Types      []TypeSlot       `json:"types"`
}