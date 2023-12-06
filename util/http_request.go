package util

import (
	"io"
	"net/http"

	"github.com/rs/zerolog"
)

func MakeHTTPRequest(client *http.Client, logger *zerolog.Logger, httpMethod string, url string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		logger.Error().Err(err).Msg("could not create request")
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		logger.Error().Err(err).Msg("could not send request")
		return nil, err
	}

	// Close response body
	defer func() {
		if response.Body != nil {
			err := response.Body.Close()
			if err != nil {
				logger.Error().Err(err).Msg("Could not close body stream")
			}
		}
	}()

	// Read the response body
	res, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Error().Err(err).Msg("could read response")
		return nil, err
	}
	return res, nil

}
