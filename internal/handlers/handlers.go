package handlers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"strconv"
	"go-mock-database/internal/model"
	"go-mock-database/internal/service"
)

type HandlerContext struct {
	DB service.FootballDatabaseService 
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Football Database REST API")
}

func CreatePlayerHandler(ctx *HandlerContext) http.HandlerFunc { 
	return func (w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		var player model.Player
		err = json.Unmarshal(body, &player)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}
		
		err = ctx.DB.PlayerRepo.CreatePlayer(id, &player)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		playerJsonBytes, err := json.Marshal(player)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("player created"))
		log.Printf("Created %s with id %d", string(playerJsonBytes), id)
	}
}

func GetPlayerByIDHandler(ctx *HandlerContext) http.HandlerFunc { 
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		player, err := ctx.DB.PlayerRepo.GetPlayerByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		playerJsonBytes, err := json.Marshal(player)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(playerJsonBytes)
		log.Printf("Retrieved %s with id %d", string(playerJsonBytes), id)
	}
}

func CreateTeamHandler(ctx *HandlerContext) http.HandlerFunc { 
	return func (w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		var team model.Team
		err = json.Unmarshal(body, &team)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}
		
		err = ctx.DB.TeamRepo.CreateTeam(id, &team)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		teamJsonBytes, err := json.Marshal(team)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("team created"))
		log.Printf("Created %s with id %d", string(teamJsonBytes), id)
	}
}

func GetTeamByIDHandler(ctx *HandlerContext) http.HandlerFunc { 
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		player, err := ctx.DB.TeamRepo.GetTeamByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		teamJsonBytes, err := json.Marshal(player)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(teamJsonBytes)
		log.Printf("Retrieved %s with id %d", string(teamJsonBytes), id)
	}
}
