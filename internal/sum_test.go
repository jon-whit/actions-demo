package internal

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func TestSum(t *testing.T) {

	s := Sum(1, 2)
	if s != 3 {
		t.Errorf("Expected 3, but got %d", s)
	}
}

func TestAvg(t *testing.T) {

	avg := Avg(2, 2)
	if avg != 2 {
		t.Errorf("Expected 2, but got %d", avg)
	}
}

func TestDatabase(t *testing.T) {
	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Failed to connect to docker pool: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_USER=" + "postgres",
			"POSTGRES_PASSWORD=" + "password",
			"POSTGRES_DB=" + "postgres",
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: "5433"},
			},
		},
	}

	resource, err := dockerPool.RunWithOptions(&opts, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Failed to start docker container: %s", err.Error())
	}

	dsn := fmt.Sprintf("%s://%s:%s@localhost:%s/%s?sslmode=disable", "postgres", "postgres", "password", "5433", "postgres")
	if err = dockerPool.Retry(func() error {
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		log.Fatalf("Failed to establish a conn to docker Postgres database: %v", err)
	}

	if err := dockerPool.Purge(resource); err != nil {
		log.Fatalf("Failed to purge docker resource: %s", err)
	}
}
