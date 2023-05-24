package graph

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/RoongJin/pokedex-graphql-sqlite/graph/model"
	"github.com/lib/pq"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Database struct {
	DBPointer *sql.DB
}

type Resolver struct {
	DB Database
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func AddPokemon(db *sql.DB, name string, description string, category string, typeOf []string, abilities []string) int64 {
	stmt := `INSERT INTO "pokedex"("name", "description", "category", "type", "abilities") values($1,$2,$3,$4,$5)`

	_, err := db.Exec(stmt, name, description, category, pq.Array(typeOf), pq.Array(abilities))
	checkErr(err)

	rows, err := db.Query(`select * from "pokedex" where "name"=$1`, name)
	checkErr(err)
	var d1 string
	var d2 string
	var d3 string
	var d4 string
	var d5 string
	var dummy string

	for rows.Next() {
		err = rows.Scan(&d1, &d2, &d3, &d4, &d5, &dummy)
		fmt.Println("Name: " + d1)
		fmt.Println("Description: " + d2)
		fmt.Println("Category: " + d3)
		fmt.Println("Type: " + d4)
		fmt.Println("Abilities: " + d5)
	}

	id_int64, _ := strconv.ParseInt(dummy, 10, 64)
	return id_int64
}

func UpdatePokemon(db *sql.DB, id int, name string, description string, category string, typeOf []string, abilities []string) {
	stmt := `update "pokedex" set "name"=$1, "description"=$2, "category"=$3, "type"=$4, "abilities"=$5 where "id"=$6`

	_, err := db.Exec(stmt, name, description, category, pq.Array(typeOf), pq.Array(abilities), id)
	checkErr(err)
}

func DeletePokemon(db *sql.DB, id int) {
	stmt := `delete from "pokedex" where "id"=$1`

	_, err := db.Exec(stmt, id)
	checkErr(err)
}

func GetAllPokemons(db *sql.DB) ([]*model.Pokemon, error) {
	rows, err := db.Query(`select * from "pokedex"`)
	checkErr(err)

	var pokeList []*model.Pokemon
	for rows.Next() {
		var name string
		var desc string
		var category string
		var types string
		var abilities string
		var dummy string
		err = rows.Scan(&name, &desc, &category, &types, &abilities, &dummy)
		fmt.Println("Name: " + name)
		fmt.Println("Description: " + desc)
		fmt.Println("Category: " + category)
		fmt.Println("Type: " + types)
		fmt.Println("Abilities: " + abilities)

		t := strings.Split(types, " ")
		a := strings.Split(abilities, " ")

		poke := model.Pokemon{
			ID:          dummy,
			Name:        name,
			Description: desc,
			Category:    category,
			Type:        t,
			Abilities:   a,
		}
		pokeList = append(pokeList, &poke)
	}

	defer rows.Close()

	return pokeList, nil
}

func FindPokemonById(db *sql.DB, id int64) (model.Pokemon, error) {
	rows, err := db.Query(`select * from "pokedex" where "id"=$1`, id)
	checkErr(err)
	var name string
	var desc string
	var category string
	var types string
	var abilities string
	var dummy string

	for rows.Next() {
		err = rows.Scan(&name, &desc, &category, &types, &abilities, &dummy)
		fmt.Println("Name: " + name)
		fmt.Println("Description: " + desc)
		fmt.Println("Category: " + category)
		fmt.Println("Type: " + types)
		fmt.Println("Abilities: " + abilities)
	}

	dm := model.Pokemon{}
	if name == "" {
		return dm, fmt.Errorf("Pokemon with this ID does not exist!")
	}

	t := strings.Split(types, " ")
	a := strings.Split(abilities, " ")

	poke := model.Pokemon{
		ID:          dummy,
		Name:        name,
		Description: desc,
		Category:    category,
		Type:        t,
		Abilities:   a,
	}

	defer rows.Close()
	return poke, nil
}
