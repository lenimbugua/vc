package api

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/lenimbugua/vc/db/sqlc"
	"github.com/lenimbugua/vc/util"
)

/* -----For Nare -----
 */
const (
	httpMethod = "POST"
	url        = "https://api.betnare.com/v1/nare-league/periods"
)

type NarePeriodsRequestBody struct {
	CompetitionID int `json:"competition_id"`
}

// prepare nare get periods request
func FormNareGetPeriodsRquest(competitionID int) ([]byte, error) {
	periods := NarePeriodsRequestBody{
		CompetitionID: 2,
	}

	// Marshal the struct to JSON
	requestBody, err := json.Marshal(periods)
	if err != nil {
		return nil, err
	}

	return requestBody, nil

}

func (server *Server) getPeriod(requestBody []byte) ([]byte, error) {
	logger := server.logger.CustomLogger()

	body, err := util.MakeHTTPRequest(server.httpClient, logger, httpMethod, url, bytes.NewReader(requestBody))
	if err != nil {
		logger.Error().Err(err).Msg("Failed")
		return nil, err
	}

	return body, nil
}

type nareCustomTime struct {
	time.Time
}

func (ct *nareCustomTime) UnmarshalJSON(b []byte) error {
	// Parse the time using the desired format
	t, err := time.Parse("\"2006-01-02 15:04:05\"", string(b))
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

type Period struct {
	CompetitionID int            `json:"competition_id"`
	EndTime       nareCustomTime `json:"end_time"`
	RoundID       int64          `json:"round_id"`
	RoundNumber   int            `json:"round_number"`
	StartTime     nareCustomTime `json:"start_time"`
}

func (server *Server) unmarshalNarePeriods(jsonData []byte) ([]Period, error) {
	logger := server.logger.CustomLogger()

	var periods []Period

	err := json.Unmarshal(jsonData, &periods)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal nare periods")
		return nil, err
	}
	return periods, nil
}

func (server *Server) saveNarePeriods(periods []Period) {
	logger := server.logger.CustomLogger()

	for _, period := range periods {

		_, err := server.dbStore.CreatePeriod(context.Background(), db.CreatePeriodParams{
			CompetitionID: int32(period.CompetitionID),
			EndTime:       period.EndTime.Time,
			RoundID:       period.RoundID,
			StartTime:     period.StartTime.Time,
		})

		if err != nil {
			logger.Error().Err(err).Msg("Could not save period to db")
			continue
		}
	}

}

func (server *Server) nareVirtualLeaguePeriods(ctx *gin.Context) {
	logger := server.logger.CustomLogger()

	req, err := FormNareGetPeriodsRquest(2)
	if err != nil {
		logger.Error().Err(err).Msg("Could not form nare request")
		return
	}
	jsonPeriods, err := server.getPeriod(req)
	if err != nil {
		logger.Error().Err(err).Msg("Could not get periods")
		return
	}
	periods, err := server.unmarshalNarePeriods(jsonPeriods)

	if err != nil {
		logger.Error().Err(err).Msg("Could not unmarshal periods")
		return
	}

	server.saveNarePeriods(periods)
}
