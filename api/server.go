package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/lenimbugua/vc/db/sqlc"
	"github.com/lenimbugua/vc/logger"
	"github.com/lenimbugua/vc/util"
)

// server serves HTTP  requests for hazini
type Server struct {
	dbStore    *db.Store
	logger     logger.Log
	router     *gin.Engine
	config     *util.Config
	httpClient *http.Client
}

// Newserver creates a new HTTP server and sets up routing
func NewServer(dbStore *db.Store, logger logger.Log, config *util.Config, httpClient *http.Client) (*Server, error) {

	server := &Server{
		dbStore:    dbStore,
		logger:     logger,
		config:     config,
		httpClient: httpClient,
	}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("/periods", server.createPeriod)

	server.router = router
}

// start runs the server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
