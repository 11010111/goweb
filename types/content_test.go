package types_test

import (
	"testing"

	"github.com/11010111/goweb/types"
)

func TestContent(t *testing.T) {
	content := types.NewContent()
	content.SetContentType("hero")
	content.SetBlock(types.Hero{})

	if content.GetContentType() != "hero" {
		t.Error("ContentType test failed")
	}

	if content.GetBlock() == nil {
		t.Error("Block test failed")
	}
}
