package main

import (
	"context"
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia/helpers/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type StaticEndpoint struct {
	URI string
	content string
}

// AddEndpoints returns the static endpoints for the rest interface
func AddEndpoints(routeGroup *gin.RouterGroup)  {
	pgConn := db.PostgresDatabase()
	query := fmt.Sprintf("select endpoint, value from rest_static_endpoints where active = true")

	rows, err := pgConn.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	endpoints := routeGroup.Group("/")

	for rows.Next() {
		var endpoint StaticEndpoint
		err := rows.Scan(
			&endpoint.URI,
			&endpoint.content,
		)
		if err != nil {
			log.Error("error reading endpoint ", err)
			continue
		}

		log.Info("added endpoint ", endpoint.URI)
		endpoints.GET("/" + endpoint.URI, func(context *gin.Context) {
			context.String(http.StatusOK, endpoint.content)
		})
	}
}