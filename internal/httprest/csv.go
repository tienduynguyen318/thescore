package httprest

import (
	"fmt"
	"thescore/internal/domain"
)

func convertDomainToPlayerToCSVData(player domain.Player) []string{
	var res = make([]string, 0)
	res = append(res, player.Player())
	res = append(res, player.Team())
	res = append(res, player.Position())
	res = append(res, fmt.Sprintf("%v", player.RushingAttempts()))
	res = append(res, fmt.Sprintf("%v", player.RushingAttemptsPerGame()))
	res = append(res, fmt.Sprintf("%v", player.TotalRushingYards()))
	res = append(res, fmt.Sprintf("%v", player.RushingAverageYardsPerAttempt()))
	res = append(res, fmt.Sprintf("%v", player.RushingYardsPerGame()))
	res = append(res, fmt.Sprintf("%v", player.TotalRushingTouchdown()))
	res = append(res, fmt.Sprintf("%v", player.LongestRush()))
	res = append(res, fmt.Sprintf("%v", player.RushingFirstDown()))
	res = append(res, fmt.Sprintf("%v", player.RushingFirstDownPercentage()))
	res = append(res, fmt.Sprintf("%v", player.Rushing20YardsEach()))
	res = append(res, fmt.Sprintf("%v", player.Rushing40YardsEach()))
	res = append(res, fmt.Sprintf("%v", player.RushingFumbles()))
	return res
}
