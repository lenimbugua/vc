package util

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func LogFile(config *Config) ( *os.File, error) {

	// Extract the directory path from the log file path
	dir := filepath.Dir(config.LogFile)

	// Create the directory and any necessary parent directories
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create directory")
		return nil, err
	}

	file, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create log file")
		return nil, err
	}

	return file, nil

}
