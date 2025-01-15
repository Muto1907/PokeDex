package internal

func Get_pokemon_names_from_location_area(location_area Location_area) []string {
	res := []string{}
	for _, pokemon := range location_area.PokemonEncounters {
		res = append(res, pokemon.Pokemon.Name)
	}
	return res
}
