package theme0

import (
	"github.com/GoAdminGroup/components/login"
	"github.com/GoAdminGroup/go-admin/template"
)

type Theme0 struct {
	*template.BaseComponent
}

func (*Theme0) GetAssetList() []string {
	return AssetsList
}

func (*Theme0) GetAsset(name string) ([]byte, error) {
	return Asset(name)
}

func (*Theme0) GetHTML() string {
	return List["login"]
}

func init() {
	login.Register("theme0", new(Theme0))
}
