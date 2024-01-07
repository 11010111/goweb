package types_test

import (
	"testing"

	"github.com/11010111/goweb/types"
)

func TestHero(t *testing.T) {
	hero := types.NewHero()
	hero.SetHeadline("Hero Headline")
	hero.SetBodytext("Hero Bodytext")

	if hero.GetHeadline() != "Hero Headline" {
		t.Error("Headline test failed")
	}

	if hero.GetBodytext() != "Hero Bodytext" {
		t.Error("Bodytext test failed")
	}
}
