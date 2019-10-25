package fixtures

import (
	"testing"

	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
)

func featureRootTestData() []Test {
	feature := "root"

	return []Test{
		shouldFailIfYamlFileIsNotPresent(feature),
		shouldFailIfYamlFileIsNotInCorrectYamlFormat(feature),
		shouldShowHelpMessageIfNoCommandSpecified(feature),
	}
}

func shouldFailIfYamlFileIsNotPresent(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldFailIfYamlFileIsNotPresent",
		CmdArgs: []string{},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "no '"+def.ConfigFileName+"' file found in current directory")
		},
	}
}

func shouldFailIfYamlFileIsNotInCorrectYamlFormat(feature string) Test {
	erroredFileMessage :=
		`Unable to parse '1build.yaml' config file. Make sure file is in correct format.
Sample format is:

------------------------------------------------------------------------
project: Sample Project
commands:
  - build: npm run build
  - lint: eslint
------------------------------------------------------------------------
`
	return Test{
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

func shouldShowHelpMessageIfNoCommandSpecified(feature string) Test {
	commandListMessage := `Please specify a command to 1build`
	defaultFileContent := `
project: Sample Project
commands:
  - build: npm run build
  - lint: eslint
`

	return Test{
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
