package domain

type TestFactory struct{}

var testFactory TestFactory

func (t *TestFactory) NewTestPlayers() []Player {
	return []Player{
		{player: "Joe Banyard", totalRushingYards: 7, longestRush: "7", totalRushingTouchdown: 6},
		{player: "Shaun Hill", totalRushingYards: 5, longestRush: "9", totalRushingTouchdown: 2},
		{player: "Breshad Perriman", totalRushingYards: 2, longestRush: "2", totalRushingTouchdown: 3},
		{player: "Lance Dunbar", totalRushingYards: 31, longestRush: "10", totalRushingTouchdown: 4},
		{player: "Tyreek Hill", totalRushingYards: 267, longestRush: "70T", totalRushingTouchdown: 5},
	}
}
