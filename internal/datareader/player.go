package datareader

type Player struct {
	Player                        string      `json:"Player"`
	Team                          string      `json:"Team"`
	Position                      string      `json:"Pos"`
	RushingAttempts               int         `json:"Att"`
	RushingAttemptsPerGame        float64     `json:"Att/G"`
	TotalRushingYards             interface{} `json:"Yds"`
	RushingAverageYardsPerAttempt float64     `json:"Avg"`
	RushingYardsPerGame           float64     `json:"Yds/G"`
	TotalRushingTouchdown         int         `json:"TD"`
	LongestRush                   interface{} `json:"Lng"`
	RushingFirstDown              int         `json:"1st"`
	RushingFirstDownPercentage    float64     `json:"1st%"`
	Rushing20YardsEach            int         `json:"20+"`
	Rushing40YardsEach            int         `json:"40+"`
	RushingFumbles                int         `json:"FUM"`
}
