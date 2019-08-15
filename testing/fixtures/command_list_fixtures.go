package fixtures

import (
	"github.com/gopinath-langote/1buildgo/testing/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func FeatureListTestData() []Test {
	var feature = "list"
	return []Test{
		shouldShowListOfCommands(feature),
		shouldNotShowAnyCommandsIfNoCommandsFound(feature),
		shouldShowCommandsWithBeforeAndAfterIfPresent(feature),
	}
}

func shouldShowListOfCommands(feature string) Test {
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

	return Test{
		Feature: feature,
		Name:    "shouldShowListOfCommands",
		CmdArgs: []string{"list"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, commandListMessage)
		},
	}
}

func shouldNotShowAnyCommandsIfNoCommandsFound(feature string) Test {
	emptyCommandListMessage := `--------------------------------------------------
project: Sample Project
commands:
--------------------------------------------------
`
	return Test{
		Feature: feature,
		Name:    "shouldNotShowAnyCommandsIfNoCommandsFound",
		CmdArgs: []string{},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "project: Sample Project\ncommands:\n")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, emptyCommandListMessage)
		},
	}
}

func shouldShowCommandsWithBeforeAndAfterIfPresent(feature string) Test {
	expectedOutput := `--------------------------------------------------
project: Sample Project
before: pre_command
after: post_command
commands:
build | npm run build
--------------------------------------------------
`
	fileContent := `
project: Sample Project
before: pre_command
after: post_command
commands:
  - build: npm run build
`

	return Test{
		Feature: feature,
		Name:    "shouldShowCommandsWithBeforeAndAfterIfPresent",
		CmdArgs: []string{"list"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}
