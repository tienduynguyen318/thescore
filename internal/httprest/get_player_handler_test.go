package httprest

import (
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"thescore/internal/domain"
	"thescore/pkg/testutils/assert"
)

type serviceDouble struct {
	domain.TestFactory
}

func (sd *serviceDouble) StorePlayers([]domain.Player) error {
	return nil
}

func (sd *serviceDouble) FetchAllPlayers() ([]domain.Player, error) {
	return sd.NewTestPlayers(), nil
}

func (sd *serviceDouble) FetchPlayersByNames(names []string) ([]domain.Player, error) {
	players := sd.NewTestPlayers()
	for _, name := range names {
		if name == "Joe Banyard" {
			return []domain.Player{players[0]}, nil
		}
	}
	return players, nil
}

func (sd *serviceDouble) SortPlayerByTotalRushingYard([]domain.Player) []domain.Player {
	players := sd.NewTestPlayers()
	sort.Sort(domain.ByTotalRushingTouchdowns(players))
	return players
}

func (sd *serviceDouble) SortPlayerByLongestRush([]domain.Player) []domain.Player {
	players := sd.NewTestPlayers()
	sort.Sort(domain.ByLongestRush(players))
	return players
}

func (sd *serviceDouble) SortPlayerByTotalRushingTouchdowns([]domain.Player) []domain.Player {
	players := sd.NewTestPlayers()
	sort.Sort(domain.ByTotalRushingTouchdowns(players))
	return players
}

func newServiceDouble() *serviceDouble {
	return &serviceDouble{}
}

func TestGetPlayer(t *testing.T) {
	t.Parallel()
	t.Run("Valid Request - Success", func(t *testing.T) {
		t.Parallel()
		reqURL := "/player"
		resp := runServerTest(t, httptest.NewRequest(http.MethodGet, reqURL, nil))
		assert.VerifyGoldenResponse(t, resp, "testdata/all_players.golden")
	})
	t.Run("Valid Request - Success with a specific player", func(t *testing.T) {
		t.Parallel()
		reqURL := "/player?name=Joe%20Banyard"
		resp := runServerTest(t, httptest.NewRequest(http.MethodGet, reqURL, nil))
		assert.VerifyGoldenResponse(t, resp, "testdata/one_player.golden")
	})
	t.Run("Valid Request - Success with 2 specific players sort by touchdowns", func(t *testing.T) {
		t.Parallel()
		reqURL := "/player?sort=td"
		resp := runServerTest(t, httptest.NewRequest(http.MethodGet, reqURL, nil))
		assert.VerifyGoldenResponse(t, resp, "testdata/sort.golden")
	})
}
