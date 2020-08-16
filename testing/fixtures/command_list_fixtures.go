package fixtures

import (
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func featureListTestData() []Test {
	var feature = "list"
	return []Test{
		shouldShowListOfCommands(feature),
		shouldShowListOfCommandsFromSpecifiedConfigFile(feature),
		shouldShowListOfCommandsFromSpecifiedConfigFileWithFullFlagName(feature),
		shouldNotShowAnyCommandsIfNoCommandsFound(feature),
		shouldShowCommandsWithBeforeAndAfterIfPresent(feature),
	}
}

func shouldShowListOfCommands(feature string) Test {
	commandListMessage := utils.PlainBanner() + `
project: Sample Project
commands:
build | npm run build
lint | eslint
` + utils.PlainBanner()
	defaultFileContent := `
project: Sample Project
commands:
  - build: npm run build
  - lint: eslint
`

	return Test{
		Feature: feature,
		Name:    "shouldShowListOfCommands",
		CmdArgs: Args("list"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, commandListMessage)
		},
	}
}

func shouldShowListOfCommandsFromSpecifiedConfigFile(feature string) Test {
	commandListMessage := utils.PlainBanner() + `
project: Sample Project
commands:
build | npm run build
` + utils.PlainBanner()
	defaultFileContent := `
project: Sample Project
commands:
  - build: npm run build
`
	return Test{
		Feature: feature,
		Name:    "shouldShowListOfCommandsFromSpecifiedConfigFile",
		CmdArgs: func(dir string) []string {
			strings := []string{"list", "-f", dir + "/custom-directory/some-config.yaml"}
			return strings
		},
		Setup: func(dir string) error {
			_ = os.MkdirAll(dir+"/custom-directory", 0777)
			err := utils.CreateConfigFileWithName(dir+"/custom-directory", "some-config.yaml", defaultFileContent)
			return err
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, commandListMessage)
		},
	}
}

func shouldShowListOfCommandsFromSpecifiedConfigFileWithFullFlagName(feature string) Test {
	commandListMessage := utils.PlainBanner() + `
project: Sample Project
commands:
build | npm run build
` + utils.PlainBanner()
	defaultFileContent := `
project: Sample Project
commands:
  - build: npm run build
`
	return Test{
		Feature: feature,
		Name:    "shouldShowListOfCommandsFromSpecifiedConfigFileWithFullFlagName",
		CmdArgs: func(dir string) []string {
			strings := []string{"list", "--file", dir + "/custom-directory/some-config.yaml"}
			return strings
		},
		Setup: func(dir string) error {
			_ = os.MkdirAll(dir+"/custom-directory", 0777)
			err := utils.CreateConfigFileWithName(dir+"/custom-directory", "some-config.yaml", defaultFileContent)
			return err
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, commandListMessage)
		},
	}
}

func shouldNotShowAnyCommandsIfNoCommandsFound(feature string) Test {
	emptyCommandListMessage := utils.PlainBanner() + `
project: Sample Project
commands:
` + utils.PlainBanner()
	return Test{
		Feature: feature,
		Name:    "shouldNotShowAnyCommandsIfNoCommandsFound",
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "project: Sample Project\ncommands:\n")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, emptyCommandListMessage)
		},
	}
}

func shouldShowCommandsWithBeforeAndAfterIfPresent(feature string) Test {
	expectedOutput := utils.PlainBanner() + `
project: Sample Project
before: pre_command
after: post_command
commands:
build | npm run build
` + utils.PlainBanner()
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
		CmdArgs: Args("list"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}
