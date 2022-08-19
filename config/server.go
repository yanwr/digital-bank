package config

import (
	"log"
	"os"
	"yanwr/digital-bank/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func CreateServer() Server {
	return Server{
		port:   os.Getenv(SERVER_PORT),
		server: gin.Default(),
	}
}

func (s *Server) LoadRoutesAndRunServer() {
	router := routes.LoadRoutes(s.server)
	log.Fatal(router.Run(":" + s.port))
}
