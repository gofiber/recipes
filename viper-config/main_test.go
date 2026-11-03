package main

import "testing"

// go test runs in the recipe directory, so config.yaml is read as well.
func TestLoadConfigLayers(t *testing.T) {
	t.Setenv("APP_PORT", "8080")
	cfg, err := loadConfig()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Service.Version != "1.0.0" || cfg.Env != "staging" || cfg.Port != 8080 {
		t.Fatalf("got version %q (default), env %q (file), port %d (env)", cfg.Service.Version, cfg.Env, cfg.Port)
	}
}
