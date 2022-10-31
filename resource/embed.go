package resource

import _ "embed"

var (
	//go:embed application.yaml
	ApplicationConfigIn string
)
