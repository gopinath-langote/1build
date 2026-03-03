package fixtures

import (
	"os"
	"testing"

	"github.com/gopinath-langote/1build/testing/def"
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
)

func featureRenameTestData() []Test {
	feature := "rename"
	return []Test{
		shouldRenameCommand(feature),
		shouldRenameCommandInSpecifiedFile(feature),
		shouldFailRenameWhenOldCommandNotFound(feature),
		shouldFailRenameWhenNewNameAlreadyExists(feature),
		shouldFailRenameWhenConfigFileNotFound(feature),
	}
}

func shouldRenameCommand(feature string) Test {
	defaultFileContent := `
project: Sample Project
commands:
  - build:
        command: go build
  - test:
        command: go test ./...
`
	expectedContent := `project: Sample Project
commands:
  - compile:
        command: go build
  - test:
        command: go test ./...
`
	return Test{
		Feature: feature,
		Name:    "shouldRenameCommand",
		CmdArgs: Args("rename", "build", "compile"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/" + def.ConfigFileName
			assert.FileExists(t, filePath)
			content, _ := os.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedContent) &&
				assert.Contains(t, actualOutput, "Renamed command 'build' to 'compile'.")
		},
	}
}

func shouldRenameCommandInSpecifiedFile(feature string) Test {
	defaultFileContent := `
project: Sample Project
commands:
  - build:
        command: go build
`
	expectedContent := `project: Sample Project
commands:
  - compile:
        command: go build
`
	return Test{
		Feature: feature,
		Name:    "shouldRenameCommandInSpecifiedFile",
		CmdArgs: func(dir string) []string {
			return []string{"rename", "build", "compile", "-f", dir + "/some-dir/some-config.yaml"}
		},
		Setup: func(dir string) error {
			_ = os.MkdirAll(dir+"/some-dir", 0750)
			return utils.CreateConfigFileWithName(dir+"/some-dir", "some-config.yaml", defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			filePath := dir + "/some-dir/some-config.yaml"
			assert.FileExists(t, filePath)
			content, _ := os.ReadFile(filePath)
			return assert.Contains(t, string(content), expectedContent) &&
				assert.Contains(t, actualOutput, "Renamed command 'build' to 'compile'.")
		},
	}
}

func shouldFailRenameWhenOldCommandNotFound(feature string) Test {
	defaultFileContent := `
project: Sample Project
commands:
  - build:
        command: go build
`
	return Test{
		Feature:          feature,
		Name:             "shouldFailRenameWhenOldCommandNotFound",
		CmdArgs:          Args("rename", "nonexistent", "compile"),
		ExpectedExitCode: 1,
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Error: command 'nonexistent' not found in configuration.")
		},
	}
}

func shouldFailRenameWhenNewNameAlreadyExists(feature string) Test {
	defaultFileContent := `
project: Sample Project
commands:
  - build:
        command: go build
  - test:
        command: go test ./...
`
	return Test{
		Feature:          feature,
		Name:             "shouldFailRenameWhenNewNameAlreadyExists",
		CmdArgs:          Args("rename", "build", "test"),
		ExpectedExitCode: 1,
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, defaultFileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "Error: command 'test' already exists in configuration.")
		},
	}
}

func shouldFailRenameWhenConfigFileNotFound(feature string) Test {
	return Test{
		Feature:          feature,
		Name:             "shouldFailRenameWhenConfigFileNotFound",
		CmdArgs:          Args("rename", "build", "compile"),
		ExpectedExitCode: 1,
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, "no '"+def.ConfigFileName+"' file found")
		},
	}
}
