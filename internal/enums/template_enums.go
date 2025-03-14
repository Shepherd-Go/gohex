package enums

type TemplateLabel int

const (
	GetMainFile TemplateLabel = iota
	GetLaunchFile
	GetDiFile
	GetConfigFile
	GetHealthFile
	GetRouterFile
	GetEnvFile
	GetGHAIntegration
)

var templateLabel = [...]string{
	"GetMainFile",
	"GetLaunchFile",
	"GetDiFile",
	"GetConfigFile",
	"GetHealthFile",
	"GetRouterFile",
	"GetEnvFile",
	"GetGHAIntegration",
}

func (value TemplateLabel) String() string {
	return templateLabel[value]
}
