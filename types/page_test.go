package types_test

import (
	"testing"

	"github.com/11010111/goweb/types"
)

func TestPage(t *testing.T) {
	page := types.NewPage()
	page.SetTitle("Home")
	page.SetDescription("Home Description")
	page.SetIndex("index")
	page.SetFollow("follow")
	page.SetRoutes([]types.Route{})
	page.SetContent([]types.Content[types.Block]{})

	if page.GetTitle() != "Home" {
		t.Error("Title test failed")
	}

	if page.GetDescription() != "Home Description" {
		t.Error("Description test failed")
	}

	if page.GetIndex() != "index" && page.GetIndex() != "noindex" {
		t.Error("Index test failed")
	}

	if page.GetFollow() != "follow" && page.GetFollow() != "nofollow" {
		t.Error("Follow test failed")
	}

	if page.GetRoutes() == nil {
		t.Error("Routea test failed")
	}

	if page.GetContent() == nil {
		t.Error("Content test failed")
	}

	page.AppendRoute(types.Route{})
	page.AppendContent(types.Content[types.Block]{})

	if len(page.GetRoutes()) == 0 {
		t.Error("Appen Route test failed")
	}

	if len(page.GetContent()) == 0 {
		t.Error("Append Content test failed")
	}
}
