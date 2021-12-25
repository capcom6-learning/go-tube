package config

import (
	"testing"
)

func TestGetConfigEmpty(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("DBHOST", "")
	t.Setenv("DBNAME", "")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic on empty enviroment variables")
		}
	}()

	GetConfig("")

	t.TempDir()
}

func TestGetConfig(t *testing.T) {
	t.Setenv("PORT", "8080")
	t.Setenv("DBHOST", "host")
	t.Setenv("DBNAME", "name")

	cfg := GetConfig("")
	if cfg.Port != 8080 {
		t.Errorf("Got %v, want %v", cfg.Port, 8080)
	}

	if cfg.DbHost != "host" {
		t.Errorf("Got %v, want %v", cfg.DbHost, "host")
	}

	if cfg.DbName != "name" {
		t.Errorf("Got %v, want %v", cfg.DbName, "name")
	}
}
