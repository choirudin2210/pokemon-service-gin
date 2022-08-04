package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/pokemons", getPokemons)
    router.GET("/pokemons/:id", getPokemonByID)
    router.POST("/pokemons", postPokemons)

    router.Run("localhost:8080")
}

type Pokemon struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Desc  string  `json:"desc"`
    Category string  `json:"category"`
    Weight  float64 `json:"weight"`
    Height  float64 `json:"height"`
    Abilities  string `json:"abilities"`
}

var pokemons = []Pokemon{
    {ID: "1", Title: "Bulbasaur", Desc: "There is a plant seed on its back right from the day this Pokémon is born. The seed slowly grows larger.", Category: "Seed", Weight: 15.2, Height: 2.04, Abilities: "Overgrow"},
    {ID: "2", Title: "Ivysaur", Desc: "When the bulb on its back grows large, it appears to lose the ability to stand on its hind legs.",Category: "Seed", Weight: 28.7,  Height: 3.03, Abilities: "Overgrow"},
    {ID: "3", Title: "Venusaur", Desc: "Its plant blooms when it is absorbing solar energy. It stays on the move to seek sunlight.",Category: "Seed", Weight: 220.5, Height: 6.07, Abilities: "Overgrow"},
    {ID: "4", Title: "Charmander", Desc: "It has a preference for hot things. When it rains, steam is said to spout from the tip of its tail.",Category: "Lizard", Weight: 18.7, Height: 2.00, Abilities: "Blaze"},
}

func getPokemons(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, pokemons)
}

func postPokemons(c *gin.Context) {
    var newPokemon Pokemon

    // Call BindJSON to bind the received JSON to
    // newPokemon.
    if err := c.BindJSON(&newPokemon); err != nil {
        return
    }

    // Add the new Pokemon to the slice.
    pokemons = append(pokemons, newPokemon)
    c.IndentedJSON(http.StatusCreated, newPokemon)
}

func getPokemonByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of Pokemons, looking for
    // an Pokemon whose ID value matches the parameter.
    for _, a := range pokemons {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pokemon not found"})
}
