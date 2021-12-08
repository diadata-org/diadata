package diaApi

import (
	"context"
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia/helpers/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Hash password using the Bcrypt hashing algorithm
// and then return the hashed password as a
// base64 encoded string
func hashPassword(password string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	return string(hashedPasswordBytes), err
}

// Check if two passwords match using Bcrypt's CompareHashAndPassword
// which return nil on success and an error on failure.
func doPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}

type RestBasicAuth struct {
	username string
	password string
}

func BasicAuth(c *gin.Context) {
	username, password, hasAuth := c.Request.BasicAuth()
	if !hasAuth {
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	}

	postgres := db.PostgresDatabase()
	today := time.Now().Format("2006-01-02")

	query := fmt.Sprintf("SELECT username, password from rest_basicauth where username = %s AND is_active = true AND (active_until IS NULL OR active_until <= %s", username, today)

	rows, err := postgres.Query(context.Background(), query)
	if err != nil {
		log.Error("Run basicauth user search query:", err)
		return
	}
	for rows.Next() {
		var basicAuth RestBasicAuth
		err := rows.Scan(
			&basicAuth.username,
			&basicAuth.password,
		)
		if err != nil {
			log.Error(err)
			return
		}
		// Get the Basic Authentication credentials
		if doPasswordsMatch(basicAuth.password, password) {
			log.WithFields(log.Fields{
				"user": username,
				"endpoint": c.Request.URL.Path,
			}).Info("User authenticated")
		} else {
			c.Abort()
			c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			return
		}
	}
	defer rows.Close()
}
