package types_test

import (
	"testing"

	"github.com/11010111/goweb/types"
)

func TestRoute(t *testing.T) {
	route := types.NewRoute()
	route.SetTitle("Link")
	route.SetHref("/")
	route.SetTarget(types.TARGET_SELF.String())
	route.SetRelation(types.REL_NONE.String())

	if route.GetTitle() != "Link" {
		t.Error("Title test failed")
	}

	if route.GetHref() != "/" {
		t.Error("Href test failed")
	}

	if route.GetTarget() != "_self" {
		t.Error("Target test failed")
	}

	if route.GetRelation() != "" {
		t.Error("Relation test failed")
	}
}
