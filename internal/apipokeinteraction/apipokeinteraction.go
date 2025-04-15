package apipokeinteraction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Ciobi0212/pokedex/internal/models"
	"github.com/Ciobi0212/pokedex/internal/pokecache"
)

func FetchAndUnmarshal[T any](url string, cache *pokecache.Cache) (*T, error) {
	cacheKey := url
	cacheEntry, ok := cache.Get(cacheKey)

	var jsonBytes []byte

	if !ok {
		fmt.Println("---- FETCHING DATA ----")
		res, err := http.Get(url)

		if err != nil {
			return nil, fmt.Errorf("error getting calling endpoint: %w", err)
		}

		resBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading bytes from response %w", err)
		}

		cache.Add(cacheKey, resBytes)

		jsonBytes = resBytes
	} else {
		fmt.Println("---- USING CACHE ----")
		jsonBytes = cacheEntry.Entry
	}

	var decodedJson T
	err := json.Unmarshal(jsonBytes, &decodedJson)

	if err != nil {
		return nil, fmt.Errorf("error unmarshaling the json")
	}

	return &decodedJson, nil
}

func GetLocationAreas(offset int, cache *pokecache.Cache) ([]string, error) {
	url := "https://pokeapi.co/api/v2/location-area?offset=" + strconv.Itoa(offset)

	decodedJson, err := FetchAndUnmarshal[models.LocationAreaResponse](url, cache)

	if err != nil {
		return nil, fmt.Errorf("error fetch and unmarshal: %w", err)
	}

	results := decodedJson.Results

	var mapNames []string
	for _, result := range results {
		name := result.Name
		mapNames = append(mapNames, name)
	}

	return mapNames, nil
}

func GetPokemonsFromArea(areaName string, cache *pokecache.Cache) ([]string, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + areaName

	areaInfo, err := FetchAndUnmarshal[models.AreaPokemonInfo](url, cache)

	if err != nil {
		return nil, fmt.Errorf("error fetch and unmarshal: %w", err)
	}

	var pokemonNames []string

	for _, encounter := range areaInfo.Encounters {
		pokemonNames = append(pokemonNames, encounter.Pokemon.Name)
	}

	return pokemonNames, nil
}

func GetPokemonInfo(pokemonName string, cache *pokecache.Cache) (*models.Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	pokemonInfo, err := FetchAndUnmarshal[models.Pokemon](url, cache)

	if err != nil {
		return nil, fmt.Errorf("error fetch and unmarshal: %w", err)
	}

	return pokemonInfo, nil
}
