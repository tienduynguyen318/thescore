package datareader

type TestFactory struct{}

var testFactory TestFactory

func (t *TestFactory) NewTestPlayer() Player {
	return Player{Player: "Joe Banyard", TotalRushingYards: "1,034"}
}
