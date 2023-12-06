package api

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lenimbugua/vc/util"
)

const (
	httpMethod        = "POST"
	url               = "https://api.betnare.com/v1/nare-league/periods"
)

func (server *Server) createPeriod(ctx *gin.Context) {

	logger := server.logger.CustomLogger(ctx)

	// Example POST request with JSON payload
	requestBody := []byte(`{
		"competition_id": 2
	  }`)

	body, err := util.MakeHTTPRequest(server.httpClient, logger, httpMethod, url, bytes.NewReader(requestBody))
	if err != nil {
		logger.Error().Err(err).Msg("Failed")
	}

	// Print the response
	logger.Info().Msgf(string(body))

	ctx.JSON(http.StatusOK, nil)
}
