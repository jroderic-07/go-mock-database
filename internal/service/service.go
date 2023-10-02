package service

import (
	"go-mock-database/internal/model"
)

type PlayerRepository interface {
	CreatePlayer(int, *model.Player) error
	GetPlayerByID(int) (*model.Player, error)
}

type TeamRepository interface {
	CreateTeam(int, *model.Team) error
	GetTeamByID(int) (*model.Team, error)
}

type FootballDatabaseService struct {
	PlayerRepo PlayerRepository
	TeamRepo TeamRepository
}

func NewFootballDatabaseService(playerRepo PlayerRepository, teamRepo TeamRepository) *FootballDatabaseService {
	return &FootballDatabaseService{
		PlayerRepo: playerRepo,
		TeamRepo: teamRepo,
	}
}
