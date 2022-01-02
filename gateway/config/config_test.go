package config

import (
	"testing"
)

func TestGetConfigEmpty(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("METADATA_HOST", "")

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
	t.Setenv("METADATA_HOST", "host")

	cfg := GetConfig("")
	if cfg.Port != 8080 {
		t.Errorf("Got %v, want %v", cfg.Port, 8080)
	}

	if cfg.MetadataHost != "host" {
		t.Errorf("Got %v, want %v", cfg.MetadataHost, "host")
	}
}
