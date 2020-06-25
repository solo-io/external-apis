package codegen

import (
	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib"
)

const (
	apiRoot = "pkg/api"
)

var Groups []model.Group

func Command() codegen.Command {
	for i, group := range Groups {
		group.RenderClients = true
		group.RenderController = true
		group.MockgenDirective = true
		group.CustomTemplates = contrib.AllCustomTemplates
		Groups[i] = group
	}

	return codegen.Command{
		Groups: Groups,
	}
}
