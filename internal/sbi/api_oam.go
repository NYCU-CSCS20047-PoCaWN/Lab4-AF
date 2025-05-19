package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getOAMRoute() []Route {
	return []Route{
		{
			Name:    "Get Warning User",
			Method:  http.MethodGet,
			Pattern: "/warning_user",
			APIFunc: s.HTTPGetWarningUser,
		},
	}
}

func (s *Server) HTTPGetWarningUser(c *gin.Context) {
	s.Processor().GetWarningUsers(c)
}
