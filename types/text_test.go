package types_test

import (
	"strings"
	"testing"

	"github.com/11010111/goweb/types"
)

func TestText(t *testing.T) {
	text := types.NewText()
	text.SetHeadline("Text Headline")
	text.SetBodytext("Text Bodytext")

	if strings.Compare(text.GetHeadline(), "Text Headline") != 0 {
		t.Error("Headline test failed")
	}

	if strings.Compare(text.GetBodytext(), "Text Bodytext") != 0 {
		t.Error("Bodytext test failed")
	}
}
