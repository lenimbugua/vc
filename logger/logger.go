package logger

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lenimbugua/vc/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Log interface {
	CustomLogger(c *gin.Context) zerolog.Logger
}

type LogStore struct {
	config *util.Config
	file   *os.File
}

func NewLogger(config *util.Config, file *os.File) Log {
	return &LogStore{
		config: config,
		file:   file,
	}
}

func (logStore *LogStore) CustomLogger(c *gin.Context) zerolog.Logger {
	// before request
	t := time.Now()

	// Set example variable
	c.Set("example", "12345")

	c.Next()

	// after request
	latency := time.Since(t)

	// access the status we are sending
	status := c.Writer.Status()

	rec := &ResponseRecorder{
		StatusCode: http.StatusOK,
	}

	var writer = io.MultiWriter(logStore.file)

	/* If env is prod write to file else write to std output */
	if logStore.config.Env == "dev" {
		writer = io.MultiWriter(os.Stderr)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: writer, TimeFormat: "2006-01-02 15:04:05"}).With().Timestamp().Caller().Logger()
	logger := log.Info()
	if rec.StatusCode != http.StatusOK {
		logger = log.Error().Bytes("body", rec.Body)
	}

	// Custom format
	logger.Str("origin", c.GetHeader("Referer")).
		Str("client_ip", c.ClientIP()).
		Time("time", time.Now()).
		Str("method", c.Request.Method).
		Str("path", c.FullPath()).
		Str("protocol", c.Request.Proto).
		Int("status_code", status).
		Dur("duration", latency).
		Str("user_agent", c.Request.UserAgent()).Msg(c.Request.Host)

	return log.Logger
}

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (rec *ResponseRecorder) WriteHeader(statusCode int) {
	rec.StatusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func (rec *ResponseRecorder) Write(body []byte) (int, error) {
	rec.Body = body
	return rec.ResponseWriter.Write(body)
}
