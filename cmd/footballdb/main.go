package main

import (
	"flag"
	"log"
	"runtime"
	"go-mock-database/internal/service"
	"go-mock-database/internal/model"
	"go-mock-database/internal/storage"
)

func main() {
	var cpuNum = flag.Int("cpuNum", 0, "num of cpus to use")

	flag.Parse()

	if *cpuNum != 0 {
		runtime.GOMAXPROCS(*cpuNum)
	}

	footballDatabaseService := service.NewFootballDatabaseService(
                storage.NewPlayerMemRepo(), storage.NewTeamMemRepo(),
    )

    log.Println(footballDatabaseService.PlayerRepo)

	_ = footballDatabaseService.PlayerRepo.CreatePlayer(0, &model.Player{
		Age: 10,
		TeamId: 0,
		FirstName: "joe",
		LastName: "smith",
		Position: "forward",
	})

	player, _ := footballDatabaseService.PlayerRepo.GetPlayerByID(0)
	log.Println(player)

	log.Println(footballDatabaseService.TeamRepo)

	_ = footballDatabaseService.TeamRepo.CreateTeam(0, &model.Team{
		Name: "Reading FC",
	})

	team, _ := footballDatabaseService.TeamRepo.GetTeamByID(0)
	log.Println(team)

}
