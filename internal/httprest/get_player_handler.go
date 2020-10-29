package httprest

import (
	"net/http"
	"thescore/internal/domain"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) getPlayer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	names := s.parseNameParam(r)
	sort := s.parseSortParam(r)
	var err error
	var players = make([]domain.Player, 0)
	if len(names) == 0 {
		players, err = s.service.FetchAllPlayers()
	} else {
		players, err = s.service.FetchPlayersByNames(names)
	}
	if err != nil {
		s.logger.Errorf("error fetch player")
		s.JSONErrorResponse(w, r, newHTTPErrorFromDomain(err))
		return
	}
	if sort != "" {
		switch sort {
		case "td", "touchdown":
			players = s.service.SortPlayerByTotalRushingTouchdowns(players)
		case "lng", "longestrush":
			players = s.service.SortPlayerByTotalRushingTouchdowns(players)
		case "yds", "rushingyard":
			players = s.service.SortPlayerByTotalRushingTouchdowns(players)
		default:
			s.logger.Errorf("Unknown sort")
			s.JSONErrorResponse(w, r, newHTTPErrorFromDomain(err))
			return
		}
	}
	if err = s.MarshalPayload(w, http.StatusOK, newPlayers(players)); err != nil {
		s.logger.Errorf("error marshalling players: %+v\n", err)
		s.JSONErrorResponse(w, r, newInternalServerError(err))
		return
	}
}
