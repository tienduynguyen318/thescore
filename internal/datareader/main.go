package datareader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"thescore/internal/domain"
	"thescore/pkg/logger"
)

type dataReader struct {
	logger logger.Logger
}

type Config struct {
	Logger logger.Logger
}

func NewDataReader(config Config) *dataReader {
	return &dataReader{
		logger: config.Logger,
	}
}

func (dr *dataReader) ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

func (dr *dataReader) ParseFile(data []byte) ([]Player, error) {
	players := make([]Player, 0)
	err := json.Unmarshal(data, &players)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (dr *dataReader) ToDomainPlayers(players []Player) ([]domain.Player, error) {
	res := make([]domain.Player, 0)
	for _, player := range players {
		domainPlayer, err := toDomainPlayer(player)
		if err != nil {
			dr.logger.Errorf("error convert to domain model %v", err)
			continue
		}
		res = append(res, domainPlayer)
	}
	return res, nil
}

func toDomainPlayer(player Player) (domain.Player, error) {
	totalRushingYards, err := player.sanitizeTotalRushingYards()
	if err != nil {
		return domain.Player{}, err
	}
	return domain.NewPlayer(domain.PlayerAttribute{
		Player:                        player.Player,
		Team:                          player.Team,
		Position:                      player.Position,
		RushingAttempts:               player.RushingAttempts,
		RushingAttemptsPerGame:        player.RushingAttemptsPerGame,
		TotalRushingYards:             totalRushingYards,
		RushingAverageYardsPerAttempt: player.RushingAverageYardsPerAttempt,
		RushingYardsPerGame:           player.RushingYardsPerGame,
		TotalRushingTouchdown:         player.TotalRushingTouchdown,
		LongestRush:                   fmt.Sprintf("%v", player.LongestRush),
		RushingFirstDown:              player.RushingFirstDown,
		RushingFirstDownPercentage:    player.RushingFirstDownPercentage,
		Rushing20YardsEach:            player.Rushing20YardsEach,
		Rushing40YardsEach:            player.Rushing40YardsEach,
		RushingFumbles:                player.RushingFumbles,
	})
}

func (p *Player) sanitizeTotalRushingYards() (int, error) {
	switch val := p.TotalRushingYards.(type) {
	case int:
		return val, nil
	case string:
		val = strings.Replace(val, ",", "", -1)
		num64, err := strconv.ParseInt(val, 10, 64)
		return int(num64), err
	case float64:
		return int(val), nil
	default:
		return 0, errors.New("Cannot parse total rushing yards")
	}
}

func (p *Player) sanitizeLongestRushingYards() string {
	return fmt.Sprintf("%v", p.LongestRush)
}
