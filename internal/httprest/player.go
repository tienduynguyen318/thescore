package httprest

import (
	"encoding/json"
	"thescore/internal/domain"
)

type player struct {
	Player                        string  `json:"Player"`
	Team                          string  `json:"Team"`
	Position                      string  `json:"Pos"`
	RushingAttempts               int     `json:"Att"`
	RushingAttemptsPerGame        float64 `json:"Att/G"`
	TotalRushingYards             int     `json:"Yds"`
	RushingAverageYardsPerAttempt float64 `json:"Avg"`
	RushingYardsPerGame           float64 `json:"Yds/G"`
	TotalRushingTouchdown         int     `json:"TD"`
	LongestRush                   string  `json:"Lng"`
	RushingFirstDown              int     `json:"1st"`
	RushingFirstDownPercentage    float64 `json:"1st%"`
	Rushing20YardsEach            int     `json:"20+"`
	Rushing40YardsEach            int     `json:"40+"`
	RushingFumbles                int     `json:"FUM"`
}

func newPlayer(p domain.Player) player {
	return player{
		Player:                        p.Player(),
		Team:                          p.Team(),
		Position:                      p.Position(),
		RushingAttempts:               p.RushingAttempts(),
		RushingAttemptsPerGame:        p.RushingAttemptsPerGame(),
		TotalRushingYards:             p.TotalRushingYards(),
		RushingAverageYardsPerAttempt: p.RushingAverageYardsPerAttempt(),
		RushingYardsPerGame:           p.RushingYardsPerGame(),
		TotalRushingTouchdown:         p.TotalRushingTouchdown(),
		LongestRush:                   p.LongestRush(),
		RushingFirstDown:              p.RushingFirstDown(),
		RushingFirstDownPercentage:    p.RushingFirstDownPercentage(),
		Rushing20YardsEach:            p.Rushing20YardsEach(),
		Rushing40YardsEach:            p.Rushing40YardsEach(),
		RushingFumbles:                p.RushingFumbles(),
	}
}

type players []player

func newPlayers(ps []domain.Player) players {
	result := make([]player, 0)
	for _, player := range ps {
		result = append(result, newPlayer(player))
	}
	return result
}

func (p players) MarshalJSON() ([]byte, error) {
	type Alias players
	return json.Marshal((Alias)(p))
}
