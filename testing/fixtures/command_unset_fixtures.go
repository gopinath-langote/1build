package fixtures

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
)

func featureUnsetTestsData() []Test {
	feature := "unset"

	return []Test{
		shouldUnsetTheExistingCommand(feature),
		shouldUnsetTheExistingCommandFromSpecifiedFile(feature),
		unsetShouldFailWhenConfigurationFileIsNotFound(feature),
		unsetShouldFailWhenConfigurationFileIsInInvalidFormat(feature),
		unsetShouldFailWhenCommandIsNotFound(feature),
		shouldUnsetMultipleCommands(feature),
		shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound1(feature),
		shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound2(feature),
		shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound3(feature),
		shouldUnsetTheBeforeCommand(feature),
		shouldUnsetTheAfterCommand(feature),
		unsetBeforeShouldFailWhenBeforeIsNotFound(feature),
		unsetAfterShouldFailWhenAfterIsNotFound(feature),
		unsetSingleCommandShouldFailWhenInvalidCommandNameIsEntered(feature),
		unsetMultipleCommandShouldFailWhenInvalidCommandNameIsEntered(feature),
	}
}

func shouldUnsetTheExistingCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetTheExistingCommand",
		CmdArgs: Args("unset", "build"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Exactly(t, expectedOutput, string(content))
		},
	}
}

func shouldUnsetTheExistingCommandFromSpecifiedFile(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetTheExistingCommandFromSpecifiedFile",
		CmdArgs: func(dir string) []string {
			return []string{"unset", "build", "-f", dir + "/some-dir/some-config.yaml"}
		},
		Setup: func(dir string) error {
			_ = os.MkdirAll(dir+"/some-dir", 0750)
			return utils.CreateConfigFileWithName(dir+"/some-dir", "some-config.yaml", defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/some-dir/some-config.yaml"
			assert.FileExists(t, filePath)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Exactly(t, expectedOutput, string(content))
		},
	}
}

func unsetShouldFailWhenCommandIsNotFound(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	return Test{
		Feature: feature,
		Name:    "unsetShouldFailWhenCommandIsNotFound",
		CmdArgs: Args("unset", "Test"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Following command(s) not found: Test")
		},
	}
}

func unsetShouldFailWhenConfigurationFileIsNotFound(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "unsetShouldFailWhenConfigurationFileIsNotFound",
		CmdArgs: Args("unset", "build"),
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "no '"+def.ConfigFileName+"' file found")
		},
	}
}

func unsetShouldFailWhenConfigurationFileIsInInvalidFormat(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "unsetShouldFailWhenConfigurationFileIsInInvalidFormat",
		CmdArgs: Args("unset", "build"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "invalid config content")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Unable to parse '"+def.ConfigFileName+"' config file. Make sure file is in correct format.")
		},
	}
}

func shouldUnsetMultipleCommands(feature string) Test {

	defaultFileContent := `
project: Sample Project
before: go before
after: go after
commands:
  - build: go build
  - test: go test
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetMultipleCommands",
		CmdArgs: Args("unset", "build", "test", "before", "after"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Exactly(t, expectedOutput, string(content))
		},
	}
}

func shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound1(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
  - test: go test
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound1",
		CmdArgs: Args("unset", "build", "test", "missingCmd"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)

			testResult := assert.Contains(t, actualOutput, "Following command(s) not found: missingCmd") &&
				assert.Exactly(t, expectedOutput, string(content))

			return testResult
		},
	}
}

func shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound2(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
  - test: go test
  - lint: go lint
`

	expectedOutput := `project: Sample Project
commands:
  - lint:
        command: go lint
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound2",
		CmdArgs: Args("unset", "build", "missingCmd", "test", "missingCmd2"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)

			testResult := assert.Contains(t, actualOutput, "Following command(s) not found: missingCmd, missingCmd2") &&
				assert.Exactly(t, expectedOutput, string(content))

			return testResult
		},
	}
}

func shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound3(feature string) Test {

	defaultFileContent := `
project: Sample Project
beforeAll: go before
afterAll: go after
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetMultipleCommandsEvenWhenCommandIsNotFound3",
		CmdArgs: Args("unset", "--beforeAll", "--afterAll", "missingCmd"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)

			testResult := assert.Contains(t, actualOutput, "Following command(s) not found: missingCmd") &&
				assert.Exactly(t, expectedOutput, string(content))

			return testResult
		},
	}
}

func shouldUnsetTheBeforeCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
beforeAll: yo
commands: []
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetBeforeAllCommand",
		CmdArgs: Args("unset", "--beforeAll"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Exactly(t, expectedOutput, string(content))
		},
	}
}

func shouldUnsetTheAfterCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
afterAll: yo
commands: []
`

	expectedOutput := `project: Sample Project
commands: []
`

	return Test{
		Feature: feature,
		Name:    "shouldUnsetTheAfterAllCommand",
		CmdArgs: Args("unset", "--afterAll"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := ioutil.ReadFile(filePath)
			return assert.Exactly(t, expectedOutput, string(content))
		},
	}
}

func unsetBeforeShouldFailWhenBeforeIsNotFound(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	return Test{
		Feature: feature,
		Name:    "unsetBeforeShouldFailWhenBeforeIsNotFound",
		CmdArgs: Args("unset", "before"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Following command(s) not found: before")
		},
	}
}

func unsetAfterShouldFailWhenAfterIsNotFound(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`

	return Test{
		Feature: feature,
		Name:    "unsetAfterShouldFailWhenAfterIsNotFound",
		CmdArgs: Args("unset", "after"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Following command(s) not found: after")
		},
	}
}

func unsetSingleCommandShouldFailWhenInvalidCommandNameIsEntered(feature string) Test {
	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`
	invalidName := "!nv@lid-name"
	expectedOutput := "1build unset: '" + invalidName + "' is not a valid command name. See '1build unset --help'."

	return Test{
		Feature: feature,
		Name:    "unsetSingleCommandShouldFailWhenInvalidCommandNameIsEntered",
		CmdArgs: Args("unset", invalidName),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}

func unsetMultipleCommandShouldFailWhenInvalidCommandNameIsEntered(feature string) Test {
	defaultFileContent := `
project: Sample Project
commands:
  - build: go build
`
	invalidName := "!nv@lid-name"
	expectedOutput := "1build unset: '" + invalidName + "' is not a valid command name. See '1build unset --help'."

	return Test{
		Feature: feature,
		Name:    "unsetSingleCommandShouldFailWhenInvalidCommandNameIsEntered",
		CmdArgs: Args("unset", "build", invalidName),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}
