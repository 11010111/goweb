package types_test

import (
	"testing"

	"github.com/11010111/goweb/types"
)

func TestForm(t *testing.T) {
	form := types.NewForm()
	form.SetHeadline("Form Headline")
	form.SetMessage("Form Message")

	if form.GetHeadline() != "Form Headline" {
		t.Error("Headline test failed")
	}

	if form.GetMessage() != "Form Message" {
		t.Error("Message test failed")
	}
}
