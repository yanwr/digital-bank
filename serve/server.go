package server

import (
	"log"
	"os"
	"yanwr/digital-bank/env"
	"yanwr/digital-bank/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IServe interface {
	RunServerAndLoadRoutes()
}

type Server struct {
	port       string
	server     *gin.Engine
	mainRoutes routes.IRoutes
}

func CreateServer(conDB *gorm.DB) Server {
	return Server{
		port:       os.Getenv(env.SERVER_PORT),
		server:     gin.Default(),
		mainRoutes: routes.NewRoutes(conDB),
	}
}

func (s *Server) RunServerAndLoadRoutes() {
	router := s.mainRoutes.LoadRoutes(s.server)
	log.Fatal(router.Run(":" + s.port))
}
