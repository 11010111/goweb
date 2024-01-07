package types_test

import (
	"testing"

	"github.com/11010111/goweb/types"
)

func TestConfirmation(t *testing.T) {
	confirm := types.NewConfirmation()
	confirm.SetHeadline("Thank you")
	confirm.SetBodytext("Confirmation Bodytext")
	confirm.SetFirstname("John")
	confirm.SetLastname("Doe")
	confirm.SetEmail("john.doe@example.com")

	if confirm.GetHeadline() != "Thank you" {
		t.Error("Headline test failed")
	}

	if confirm.GetBodytext() != "Confirmation Bodytext" {
		t.Error("Bodytext test failed")
	}

	if confirm.GetFirstname() != "John" {
		t.Error("Firstname test failed")
	}

	if confirm.GetLastname() != "Doe" {
		t.Error("Lastname test failed")
	}

	if confirm.GetEmail() != "john.doe@example.com" {
		t.Error("Email test failed")
	}
}
