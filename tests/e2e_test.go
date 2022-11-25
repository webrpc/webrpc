package tests

import (
	"testing"
	"os"
	"net/http"
	"log"
)

func TestMain(m *testing.M) {
	code := m.Run()

	client := http.Client{}
	client.Get("http://localhost:9999/shutdown")
	log.Printf("server has shutdown")

	os.Exit(code)
}
