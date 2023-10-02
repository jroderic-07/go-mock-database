package main

import (
	"flag"
	"runtime"
	"net/http"
	"go-mock-database/internal/handlers"
	"go-mock-database/internal/service"
	"go-mock-database/internal/storage"
)

func main() {
	var portNum = flag.String("portNum", "8080", "port to run the api on")
	var cpuNum = flag.Int("cpuNum", 0, "num of cpus to use")

	flag.Parse()

	if *cpuNum != 0 {
		runtime.GOMAXPROCS(*cpuNum)
	}

	footballDatabaseService := service.NewFootballDatabaseService(
		storage.NewPlayerMemRepo(), storage.NewTeamMemRepo(),
	)

	ctx := handlers.HandlerContext{
		DB: *footballDatabaseService,
	}

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/createPlayer", handlers.CreatePlayerHandler(&ctx))
	http.HandleFunc("/getPlayer", handlers.GetPlayerByIDHandler(&ctx))
	http.HandleFunc("/createTeam", handlers.CreateTeamHandler(&ctx))
	http.HandleFunc("/getTeam", handlers.GetTeamByIDHandler(&ctx))
	http.ListenAndServe(":" + *portNum, nil)
}
