package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kiambogo/coffeeworks/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
