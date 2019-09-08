package fixtures

import (
	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func featureRootTestData() []test {
	feature := "root"

	return []test{
		shouldFailIfYamlFileIsNotPresent(feature),
		shouldFailIfYamlFileIsNotInCorrectYamlFormat(feature),
		shouldShowListOfCommandsIfNoCommandSpecified(feature),
	}
}

func shouldFailIfYamlFileIsNotPresent(feature string) test {
	return test{
		Feature: feature,
		Name:    "shouldFailIfYamlFileIsNotPresent",
		CmdArgs: []string{},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "no '"+def.ConfigFileName+"' file found in current directory")
		},
	}
}

func shouldFailIfYamlFileIsNotInCorrectYamlFormat(feature string) test {
	erroredFileMessage :=
		`Unable to parse '1build.yaml' config file. Make sure file is in correct format.
Sample format is:

--------------------------------------------------
project: Sample Project
commands:
  - build: npm run build
  - lint: eslint
--------------------------------------------------
`
	return test{
		Feature: feature,
		Name:    "shouldFailIfYamlFileIsNotInCorrectYamlFormat",
		CmdArgs: []string{},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "invalid yaml file")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, erroredFileMessage)
		},
	}
}

func shouldShowListOfCommandsIfNoCommandSpecified(feature string) test {
	commandListMessage := `--------------------------------------------------
project: Sample Project
commands:
build | npm run build
lint | eslint
--------------------------------------------------
`
	defaultFileContent := `
project: Sample Project
commands:
  - build: npm run build
  - lint: eslint
`

	return test{
		Feature: feature,
		Name:    "shouldShowListOfCommandsIfNoCommandSpecified",
		CmdArgs: []string{},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, commandListMessage)
		},
	}
}
