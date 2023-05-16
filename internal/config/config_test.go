package config

import (
	"os"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	testCases := []struct {
		env map[string]string
		cfg *Config
	}{
		// check default config loading
		{
			env: map[string]string{
				"SERVE_PORT": "8080",
			},
			cfg: &Config{
				LogLevel:  "info",
				ServeHost: "0.0.0.0",
				ServePort: "8080",
			},
		},
	}

	for _, tc := range testCases {
		for k, v := range tc.env {
			if err := os.Setenv(k, v); err != nil {
				t.Errorf("Failed to set environment variable %s, %v", k, err)
			}
		}
		cfg, err := New()
		if err != nil {
			t.Errorf("Failed to load config, %v", err)
		}
		if !reflect.DeepEqual(cfg, tc.cfg) {
			t.Errorf("Config loading was incorrect\nexpected: '%+v'\ngot: '%+v'", *tc.cfg, *cfg)
		}
	}
}
