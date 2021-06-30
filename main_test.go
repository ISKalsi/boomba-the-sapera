package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

const PORT = "PORT"

func TestSetupRoutes(t *testing.T) {
	tests := []struct {
		url    string
		method string
	}{
		{"/", "GET"},
		{"/start", "POST"},
		{"/move", "POST"},
		{"/end", "POST"},
	}

	game := Game{}
	router := SetupRoutes(&game)

	for _, test := range tests {
		name := "\"" + test.url + "\""
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			body := strings.NewReader("{}")
			req := httptest.NewRequest(test.method, test.url, body)

			router.ServeHTTP(w, req)
			assert.Equal(t, 200, w.Code)
		})
	}
}

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
