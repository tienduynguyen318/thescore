package domain

type ByTotalRushingYard []Player

func (p ByTotalRushingYard) Len() int {
	return len(p)
}
func (p ByTotalRushingYard) Less(i, j int) bool {
	return p[i].totalRushingYards < p[j].totalRushingYards
}
func (p ByTotalRushingYard) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// I dont watch American Football so I am not sure how to sort this
type ByLongestRush []Player

func (p ByLongestRush) Len() int {
	return len(p)
}
func (p ByLongestRush) Less(i, j int) bool {
	return p[i].longestRush < p[j].longestRush
}
func (p ByLongestRush) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ByTotalRushingTouchdowns []Player

func (p ByTotalRushingTouchdowns) Len() int {
	return len(p)
}
func (p ByTotalRushingTouchdowns) Less(i, j int) bool {
	return p[i].totalRushingTouchdown < p[j].totalRushingTouchdown
}
func (p ByTotalRushingTouchdowns) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
