package domain

type Player struct {
	player                        string
	team                          string
	position                      string
	rushingAttempts               int
	rushingAttemptsPerGame        float64
	totalRushingYards             int
	rushingAverageYardsPerAttempt float64
	rushingYardsPerGame           float64
	totalRushingTouchdown         int
	longestRush                   string
	rushingFirstDown              int
	rushingFirstDownPercentage    float64
	rushing20YardsEach            int
	rushing40YardsEach            int
	rushingFumbles                int
}

type PlayerAttribute struct {
	Player                        string
	Team                          string
	Position                      string
	RushingAttempts               int
	RushingAttemptsPerGame        float64
	TotalRushingYards             int
	RushingAverageYardsPerAttempt float64
	RushingYardsPerGame           float64
	TotalRushingTouchdown         int
	LongestRush                   string
	RushingFirstDown              int
	RushingFirstDownPercentage    float64
	Rushing20YardsEach            int
	Rushing40YardsEach            int
	RushingFumbles                int
}

func NewPlayer(attr PlayerAttribute) (Player, error) {
	return Player{
		player:                        attr.Player,
		team:                          attr.Team,
		position:                      attr.Position,
		rushingAttempts:               attr.RushingAttempts,
		rushingAttemptsPerGame:        attr.RushingAttemptsPerGame,
		totalRushingYards:             attr.TotalRushingYards,
		rushingAverageYardsPerAttempt: attr.RushingAverageYardsPerAttempt,
		rushingYardsPerGame:           attr.RushingYardsPerGame,
		totalRushingTouchdown:         attr.TotalRushingTouchdown,
		longestRush:                   attr.LongestRush,
		rushingFirstDown:              attr.RushingFirstDown,
		rushingFirstDownPercentage:    attr.RushingFirstDownPercentage,
		rushing20YardsEach:            attr.Rushing20YardsEach,
		rushing40YardsEach:            attr.Rushing40YardsEach,
		rushingFumbles:                attr.RushingFumbles,
	}, nil
}

func (p *Player) Player() string {
	return p.player
}

func (p *Player) Team() string {
	return p.team
}

func (p *Player) Position() string {
	return p.position
}

func (p *Player) RushingAttempts() int {
	return p.rushingAttempts
}

func (p *Player) RushingAttemptsPerGame() float64 {
	return p.rushingAttemptsPerGame
}

func (p *Player) TotalRushingYards() int {
	return p.totalRushingYards
}

func (p *Player) RushingAverageYardsPerAttempt() float64 {
	return p.rushingAverageYardsPerAttempt
}

func (p *Player) RushingYardsPerGame() float64 {
	return p.rushingYardsPerGame
}

func (p *Player) TotalRushingTouchdown() int {
	return p.totalRushingTouchdown
}

func (p *Player) LongestRush() string {
	return p.longestRush
}

func (p *Player) RushingFirstDown() int {
	return p.rushingFirstDown
}

func (p *Player) RushingFirstDownPercentage() float64 {
	return p.rushingFirstDownPercentage
}

func (p *Player) Rushing20YardsEach() int {
	return p.rushing20YardsEach
}

func (p *Player) Rushing40YardsEach() int {
	return p.rushing40YardsEach
}

func (p *Player) RushingFumbles() int {
	return p.rushingFumbles
}
