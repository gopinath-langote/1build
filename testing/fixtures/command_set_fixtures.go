package fixtures

import (
	"os"
	"testing"

	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
)

func featureSetTestsData() []Test {
	feature := "set"

	return []Test{
		shouldSetNewCommand(feature),
		shouldSetNewCommandInSpecifiedFile(feature),
		shouldUpdateExistingCommand(feature),
		shouldFailWhenConfigurationFileIsNotFound(feature),
		shouldFailWhenConfigurationFileIsInInvalidFormat(feature),
		shouldSetBeforeAllCommand(feature),
		shouldSetAfterAllCommand(feature),
		shouldDryRunSet(feature),
	}
}

func shouldSetNewCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build:
        command: go build
`

	expectedOutput := `project: Sample Project
commands:
  - build:
        command: go build
  - Test:
        command: go Test
`

	return Test{
		Feature: feature,
		Name:    "shouldSetNewCommand",
		CmdArgs: Args("set", "Test", "go Test"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := os.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldSetNewCommandInSpecifiedFile(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build:
        command: go build
`

	expectedOutput := `project: Sample Project
commands:
  - build:
        command: go build
  - Test:
        command: go Test
`

	return Test{
		Feature: feature,
		Name:    "shouldSetNewCommandInSpecifiedFile",
		CmdArgs: func(dir string) []string {
			return []string{"set", "Test", "go Test", "-f", dir + "/some-dir/some-config.yaml"}
		},
		Setup: func(dir string) error {
			_ = os.MkdirAll(dir+"/some-dir", 0750)
			return utils.CreateConfigFileWithName(dir+"/some-dir", "some-config.yaml", defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/some-dir/some-config.yaml"
			assert.FileExists(t, filePath)
			content, _ := os.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldUpdateExistingCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build:
        command: go build
`

	expectedOutput := `project: Sample Project
commands:
  - build:
        command: go build -o
`

	return Test{
		Feature: feature,
		Name:    "shouldUpdateExistingCommand",
		CmdArgs: Args("set", "build", "go build -o"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := os.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldFailWhenConfigurationFileIsNotFound(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldFailWhenConfigurationFileIsNotFound",
		CmdArgs: Args("set", "build", "go build -o"),
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "no '"+def.ConfigFileName+"' file found")
		},
	}
}

func shouldFailWhenConfigurationFileIsInInvalidFormat(feature string) Test {
	return Test{
		Feature: feature,
		Name:    "shouldFailWhenConfigurationFileIsInInvalidFormat",
		CmdArgs: Args("set", "build", "go build"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, "invalid config content")
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Unable to parse '"+def.ConfigFileName+"' config file. Make sure file is in correct format.")
		},
	}
}

func shouldSetBeforeAllCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build:
        command: go build
`

	expectedOutput := `project: Sample Project
before-all: yo
commands:
  - build:
        command: go build
`

	return Test{
		Feature: feature,
		Name:    "shouldSetBeforeAllCommand",
		CmdArgs: Args("set", "--before-all", "yo"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := os.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldSetAfterAllCommand(feature string) Test {

	defaultFileContent := `
project: Sample Project
commands:
  - build:
        command: go build
`

	expectedOutput := `project: Sample Project
after-all: yo
commands:
  - build:
        command: go build
`

	return Test{
		Feature: feature,
		Name:    "shouldSetAfterCommand",
		CmdArgs: Args("set", "--after-all", "yo"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, dir+"/"+def.ConfigFileName)
			content, _ := os.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedOutput)
		},
	}
}

func shouldDryRunSet(feature string) Test {
	defaultFileContent := `
project: Sample Project
commands:
  - build:
        command: go build
`

	return Test{
		Feature: feature,
		Name:    "shouldDryRunSet",
		CmdArgs: Args("set", "release", "go build -o dist/app", "--dry-run"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			// Output should indicate dry-run
			if !assert.Contains(t, actualOutput, "[dry-run]") {
				return false
			}
			// The config file must NOT have been modified
			filePath := dir + "/" + def.ConfigFileName
			content, _ := os.ReadFile(filePath)
			return assert.NotContains(t, string(content), "release")
		},
	}
}
