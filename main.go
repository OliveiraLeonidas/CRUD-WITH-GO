package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OliveiraLeonidas/simple-crud/config"
	"github.com/OliveiraLeonidas/simple-crud/models"
	"github.com/gorilla/mux"
)

// entry point application
func main() {
	//Chama a função database
	db := config.SetupDatabase()

	//chama a funçã close para fechar a conexão ao final da operação no banco de dados
	defer db.Close()

	// crie um novo roteador com mux router
	router := mux.NewRouter()

	// migration
	_, err := db.Exec(models.CreateTableSQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	// definindo uma nova rota
	router.HandleFunc("/tasks", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Servidor rodando com sucesso!")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Hello go project"))
	}).Methods("GET")

	// rota que verifica conexão com banco de dados [using docker]
	router.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		err := db.Ping()
		if err != nil {
			http.Error(res, "Erro ao conectar ao banco de dados", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(res, "Banco de dados conectado com sucesso!")
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
