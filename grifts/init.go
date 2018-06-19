package grifts

import (
	"budayaid/budayaid_backend_sql/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
