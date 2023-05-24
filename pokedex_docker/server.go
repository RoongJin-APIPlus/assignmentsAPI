package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/RoongJin/pokedex-graphql-sqlite/graph"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

const (
	host     = "pokedex-postgres"
	port     = 5432
	user     = "postgres"
	password = "test"
	dbname   = "myPostgres"
	defPort  = "8080"
)

func main() {
	r := chi.NewRouter()

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	stmt1 := `CREATE DATABASE pokedex;`
	_, err = db.Exec(stmt1)
	if err != nil {
		log.Fatal(err)
	}

	stmt2 := `CREATE TABLE pokedex(Name VARCHAR(50) NOT NULL, Description VARCHAR(200) NOT NULL, Category VARCHAR(50) NOT NULL, Type VARCHAR(50) ARRAY NOT NULL, Abilities VARCHAR(50) ARRAY NOT NULL, ID SERIAL NOT NULL PRIMARY KEY);`
	_, err = db.Exec(stmt2)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: graph.Database{DBPointer: db},
	}}))

	//DeletePokemon(db, 1, "Bulbasaur", "There is a plant seed on its back right from the day this Pokémon is born. The seed slowly grows larger.", "Seed", "Grass Poison", "Overgrow")
	//FindPokemonById(db, 1)

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defPort)
	log.Fatal(http.ListenAndServe(":"+defPort, r))
}
