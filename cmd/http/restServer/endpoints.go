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
	URI     string
	content string
}

// AddEndpoints returns the static endpoints for the rest interface
func AddEndpoints(engine *gin.Engine) {
	pgConn := db.PostgresDatabase()
	query := fmt.Sprintf("select endpoint, value from rest_static_endpoints where active = true")

	log.Info("reading static endpoints")
	rows, err := pgConn.Query(context.Background(), query)
	if err != nil {
		log.Error("error reading endpoint from postgres", err)
		return
	}

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
		engine.GET("/"+endpoint.URI, func(context *gin.Context) {
			context.String(http.StatusOK, endpoint.content)
		})
	}
	defer rows.Close()
	log.Info("finished reading static endpoints")
}
