package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/sprokhorov/cint-test/internal/storage"
	"github.com/sprokhorov/cint-test/pkg/schema"
)

type Server struct {
	r       *gin.Engine
	storage storage.Storage
	log     *zerolog.Logger
}

func New(storage storage.Storage, log *zerolog.Logger) *Server {
	srv := &Server{r: gin.Default(), storage: storage, log: log}
	// setup routes
	srv.r.GET("/reminders", srv.remindersList)
	srv.r.POST("/reminders", srv.reminderCreate)
	srv.r.GET("/live", srv.livenessProbe)
	srv.r.GET("/ready", srv.readinessProbe)

	return srv
}

func (s *Server) Listen(host, port string) error {
	addr := fmt.Sprintf("%s:%s", host, port)
	return s.r.Run(addr)
}

func (s *Server) remindersList(c *gin.Context) {
	rmds, err := s.storage.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, rmds)
	}
}

func (s *Server) reminderCreate(c *gin.Context) {
	s.log.Debug().Msg("(s *Server).reminderCreate handler was called")
	var rmd schema.Reminder
	err := c.Bind(&rmd)
	if err == nil {
		if err := rmd.IsValid(); err != nil {
			s.log.Err(err).Msg("Failed to create reminder")
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid reminder", "error": err.Error()})
			return
		}
		err := s.storage.Create(c.Request.Context(), &rmd)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusCreated, rmd)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid reminder", "error": err.Error()})
		return
	}
}

func (s *Server) livenessProbe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (s *Server) readinessProbe(c *gin.Context) {
	if err := s.storage.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to ping storage"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
