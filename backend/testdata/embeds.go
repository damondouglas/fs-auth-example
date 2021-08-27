package testdata

import (
	_ "embed"
)

//go:embed .env
var DotEnv []byte
