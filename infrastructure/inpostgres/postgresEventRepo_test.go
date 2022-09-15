package inpostgres

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPostgresEventRepo(t *testing.T) {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "qawsed345rf"
	dbname := "calendar"

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	_, err := NewPostgresEventRepo(connectionString)
	assert.NoError(t, err)
}
