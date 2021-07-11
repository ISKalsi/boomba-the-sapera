package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

const PORT = "PORT"

func TestGetPortFromEnvVariable(t *testing.T) {
	if err := os.Setenv(PORT, "1234"); err != nil {
		log.Panic(err)
	}

	t.Cleanup(func() {
		if err := os.Unsetenv(PORT); err != nil {
			log.Panic(err)
		}
	})

	port := getPort()
	assert.Equal(t, "1234", port)
}

func TestGetDefaultPort(t *testing.T) {
	envPort, present := os.LookupEnv(PORT)
	if present {
		if err := os.Unsetenv(PORT); err != nil {
			log.Panic(err)
		}
	}

	t.Cleanup(func() {
		if err := os.Setenv(PORT, envPort); err != nil {
			log.Panic(err)
		}
	})

	port := getPort()
	assert.Equal(t, "8080", port)
}
