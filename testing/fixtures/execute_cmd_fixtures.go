package fixtures

import (
	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func featureExecuteCmdTestData() []test {
	feature := "exec"

	return []test{
		shouldExecuteAvailableCommand(feature),
		shouldShowErrorIfCommandNotFound(feature),
		shouldExecuteBeforeCommand(feature),
		shouldExecuteAfterCommand(feature),
		shouldExecuteBeforeAndAfterCommand(feature),
		shouldStopExecutionIfBeforeCommandFailed(feature),
		shouldStopExecutionIfCommandFailed(feature),
	}
}

func shouldExecuteAvailableCommand(feature string) test {
	fileContent := `
project: Sample Project
commands:
  - build: echo building project
`
	expectedOutput := `--------------------------------------------------
build : echo building project
--------------------------------------------------
building project
`
	return test{
		Feature: feature,
		Name:    "shouldExecuteAvailableCommand",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}

func shouldShowErrorIfCommandNotFound(feature string) test {
	fileContent := `
project: Sample Project
commands:
  - build: echo building project
`

	expectedOutput := `No command 'random' found in config file

--------------------------------------------------
project: Sample Project
commands:
build | echo building project
--------------------------------------------------
`
	return test{
		Feature: feature,
		Name:    "shouldShowErrorIfCommandNotFound",
		CmdArgs: []string{"random"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}

func shouldExecuteBeforeCommand(feature string) test {
	fileContent := `
project: Sample Project
before: echo running pre-command
commands:
  - build: echo building project
`
	expectedOutput := `--------------------------------------------------
Before: echo running pre-command

build : echo building project
--------------------------------------------------
running pre-command
building project
`
	return test{
		Feature: feature,
		Name:    "shouldExecuteBeforeCommand",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}

func shouldExecuteAfterCommand(feature string) test {
	fileContent := `
project: Sample Project
after: echo running post-command
commands:
  - build: echo building project
`
	expectedOutput := `--------------------------------------------------
build : echo building project

After: echo running post-command
--------------------------------------------------
building project
running post-command
`
	return test{
		Feature: feature,
		Name:    "shouldExecuteAfterCommand",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}

func shouldExecuteBeforeAndAfterCommand(feature string) test {
	fileContent := `
project: Sample Project
before: echo running pre-command
after: echo running post-command
commands:
  - build: echo building project
`
	expectedOutput := `--------------------------------------------------
Before: echo running pre-command

build : echo building project

After: echo running post-command
--------------------------------------------------
running pre-command
building project
running post-command
`
	return test{
		Feature: feature,
		Name:    "shouldExecuteBeforeAndAfterCommand",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}

func shouldStopExecutionIfBeforeCommandFailed(feature string) test {
	fileContent := `
project: Sample Project
before: invalid_command
after: echo running post-command
commands:
  - build: echo building project
`
	expectedOutput := `--------------------------------------------------
Before: invalid_command

build : echo building project

After: echo running post-command
--------------------------------------------------

Failed to execute 'invalid_command'
`
	return test{
		Feature: feature,
		Name:    "shouldStopExecutionIfBeforeCommandFailed",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}

func shouldStopExecutionIfCommandFailed(feature string) test {
	fileContent := `
project: Sample Project
before: echo running pre-command
after: echo running post-command
commands:
  - build: invalid_command
`
	expectedOutput := `--------------------------------------------------
Before: echo running pre-command

build : invalid_command

After: echo running post-command
--------------------------------------------------
running pre-command

Failed to execute 'invalid_command'
`
	return test{
		Feature: feature,
		Name:    "shouldStopExecutionIfCommandFailed",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput)
		},
	}
}
