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
		shouldShowListOfCommandsIfNoCommandSpecified(feature),
	}
}

func shouldFailIfYamlFileIsNotPresent(feature string) Test {
	return Test{
		Feature:          feature,
		Name:             "shouldFailIfYamlFileIsNotPresent",
		ExpectedExitCode: 1,
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "no '"+def.ConfigFileName+"' file found")
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
  - build:
      before: echo "before"
      command: npm run build
      after: echo "after"
------------------------------------------------------------------------
`
	return Test{
		Feature:          feature,
		Name:             "shouldFailIfYamlFileIsNotInCorrectYamlFormat",
		ExpectedExitCode: 1,
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "invalid yaml file")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, erroredFileMessage)
		},
	}
}

func shouldShowListOfCommandsIfNoCommandSpecified(feature string) Test {
	defaultFileContent := `project: Sample Project
commands:
  - build: npm run build
  - lint: eslint
`

	return Test{
		Feature: feature,
		Name:    "shouldShowListOfCommandsIfNoCommandSpecified",
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "project: Sample Project") &&
				assert.Contains(t, actualOutput, "build | npm run build") &&
				assert.Contains(t, actualOutput, "lint | eslint")
		},
	}
}
