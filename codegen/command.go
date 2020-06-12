package codegen

import (
	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib"
)

const (
	apiRoot = "pkg/api"
)

var groups []model.Group

func Command() codegen.Command {
	for i, group := range groups {
		group.RenderClients = true
		group.RenderController = true
		group.MockgenDirective = true
		group.CustomTemplates = contrib.AllCustomTemplates
		groups[i] = group
	}

	return codegen.Command{
		Groups: groups,
	}
}
