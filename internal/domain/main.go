package domain

import (
	"sort"
	"thescore/pkg/logger"
)

type service struct {
	logger logger.Logger
}

var Cache []Player

type Service interface {
	StorePlayers([]Player) error
	FetchAllPlayers() ([]Player, error)
	FetchPlayersByNames([]string) ([]Player, error)
	SortPlayerByTotalRushingYard([]Player) []Player
	SortPlayerByLongestRush([]Player) []Player
	SortPlayerByTotalRushingTouchdowns([]Player) []Player
}

type Config struct {
	Logger logger.Logger
}

func New(config Config) *service {
	return &service{
		logger: config.Logger,
	}
}

var playerTable = map[string]Player{}

func (s *service) StorePlayers(players []Player) error {
	var err error
	Cache = players
	err = s.storePlayerInTable(players)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) storePlayerInTable(players []Player) error {
	for _, player := range players {
		playerTable[player.player] = player
	}
	return nil
}

func (s *service) FetchAllPlayers() ([]Player, error) {
	players := make([]Player, 0)
	for _, player := range playerTable {
		players = append(players, player)
	}
	Cache = players
	return players, nil
}

func (s *service) FetchPlayersByNames(playerNames []string) ([]Player, error) {
	players := make([]Player, 0)
	for _, playerName := range playerNames {
		player, err := s.fetchPlayerByName(playerName)
		if err != nil {
			continue
		}
		players = append(players, player)
	}
	Cache = players
	return players, nil
}

func (s *service) fetchPlayerByName(playerName string) (Player, error) {
	player, ok := playerTable[playerName]
	if !ok {
		return Player{}, NewNotFoundError("Player", playerName)
	}
	return player, nil
}

func (s *service) SortPlayerByTotalRushingYard(players []Player) []Player {
	sort.Sort(ByTotalRushingYard(players))
	return players
}

func (s *service) SortPlayerByLongestRush(players []Player) []Player {
	sort.Sort(ByLongestRush(players))
	return players
}

func (s *service) SortPlayerByTotalRushingTouchdowns(players []Player) []Player {
	sort.Sort(ByTotalRushingTouchdowns(players))
	return players
}
