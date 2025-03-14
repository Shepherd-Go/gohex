package app

import (
	"os"

	"github.com/andresxlp/gohex/internal/enums"
	"github.com/andresxlp/gohex/internal/utils/templatesExec"
)

type Service struct{}

var (
	files = []string{enums.MainFile, enums.LaunchFile, enums.DiFile, enums.ConfigFile, enums.HealthFile, enums.RouterFile, enums.EnvFile, enums.GHAIntegrationFile, enums.GoModFile, enums.GoSumFile}
)

func (s *Service) CreateAllFiles(module string) {
	for _, file := range files[:8] {
		createFiles(file, module)
	}
}

func createFiles(fileName, module string) {
	queryTempl := map[string]enums.TemplateLabel{
		files[0]: enums.GetMainFile,
		files[1]: enums.GetLaunchFile,
		files[2]: enums.GetDiFile,
		files[3]: enums.GetConfigFile,
		files[4]: enums.GetHealthFile,
		files[5]: enums.GetRouterFile,
		files[6]: enums.GetEnvFile,
		files[7]: enums.GetGHAIntegration,
	}

	createFile(fileName, templatesExec.GetTemplateWhitValues(queryTempl[fileName], dataTemplate(module)))
}

func dataTemplate(module string) map[string]interface{} {
	dataMap := map[string]interface{}{
		enums.Module: module,
	}
	return dataMap
}

func createFile(fileName string, data string) {
	err := os.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}
