package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	t.Setenv("PORT", "3333")
	cfg, err := New()
	if err != nil {
		t.Fatal("error")
	}
	fmt.Println(cfg)
	if cfg.Port != 3333 {
		t.Fatal("Port was not 333")
	}
}
