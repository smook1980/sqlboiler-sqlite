package override

import "embed"

//go:embed templates_test
var fs embed.FS

func AssetNames() []string {
	return []string{"templates_test/singleton/sqlite_main_test.go.tpl"}
}

func Asset(name string) ([]byte, error) {
	return fs.ReadFile(name)
}
