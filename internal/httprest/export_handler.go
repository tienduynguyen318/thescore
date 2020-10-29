package httprest

import (
	"encoding/csv"
	"net/http"
	"os"
	"thescore/internal/domain"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) exportData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	file, err := os.Create("result.csv")
	if err != nil {
		s.logger.Errorf("error create csv file")
		s.JSONErrorResponse(w, r, newHTTPErrorFromDomain(err))
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	var data [][]string
	for _, player := range domain.Cache {
		data = append(data, convertDomainToPlayerToCSVData(player))
	}
	for _, value := range data {
		if err = writer.Write(value); err != nil {
			s.logger.Errorf("CSV writer error", err)
			continue
		}
	}
	w.WriteHeader(200)
}
