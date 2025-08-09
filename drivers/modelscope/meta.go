package modelscope

import (
	"github.com/OpenListTeam/OpenList/v4/internal/driver"
	"github.com/OpenListTeam/OpenList/v4/internal/op"
)

const (
	DriverName = "modelscope"
)

var config = driver.Config{
	Name:      DriverName,
	OnlyProxy: true,
}

type Addition struct {
	driver.RootPath
	ModelID  string `json:"model_id" required:"true"`
	Revision string `json:"revision" required:"true" default:"master"`
}

func (a *Addition) GetRootPath() string {
	return a.ModelID
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &ModelScope{}
	})
}
