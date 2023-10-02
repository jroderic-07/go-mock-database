package storage

import (
	"go-mock-database/internal/model"
	"sync"
	"fmt"
)

type PlayerMemRepo struct {
	players map[int]model.Player
	sync.Mutex
}

type TeamMemRepo struct {
	teams map[int]model.Team
	sync.Mutex
}

func NewPlayerMemRepo() *PlayerMemRepo {
	return &PlayerMemRepo {
		players: make(map[int]model.Player),
	}
}

func NewTeamMemRepo() *TeamMemRepo {
	return &TeamMemRepo {
		teams: make(map[int]model.Team),
	}
}

func (p *PlayerMemRepo) CreatePlayer(id int, player *model.Player) error {
	p.Lock()
	defer p.Unlock()

	_, exists := p.players[id]
	if exists {
		return fmt.Errorf("player %d already exists", id)
	} 
	
	p.players[id] = *player
	return nil
}

func (p *PlayerMemRepo) GetPlayerByID(id int) (*model.Player, error) {
	p.Lock()
	defer p.Unlock()

	player, exists := p.players[id]
	if !exists {
		return &model.Player{}, fmt.Errorf("player %d does not exist", id)
	}

	return &player, nil
}

func (t *TeamMemRepo) CreateTeam(id int, team *model.Team) error {
	t.Lock()
	defer t.Unlock()

	_, exists := t.teams[id]
	if exists {
		return fmt.Errorf("team %d already exists", id)
	}

	t.teams[id] = *team
	return nil
}

func (t *TeamMemRepo) GetTeamByID (id int) (*model.Team, error) {
	t.Lock()
	defer t.Unlock()

	team, exists := t.teams[id]
	if !exists {
		return &model.Team{}, fmt.Errorf("team %d does not exist", id)
	}

	return &team, nil
}