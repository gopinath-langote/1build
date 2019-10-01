package fixtures

import (
	"strings"
	"testing"

	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
)

func successBanner() string {
	return strings.Repeat("-", 72) + "\n" + utils.Colored("SUCCESS", utils.CYAN)
}

func featureExecuteCmdTestData() []Test {
	feature := "exec"

	return []Test{
		shouldExecuteAvailableCommand(feature),
		shouldShowErrorIfCommandNotFound(feature),
		shouldExecuteBeforeCommand(feature),
		shouldExecuteAfterCommand(feature),
		shouldExecuteBeforeAndAfterCommand(feature),
		shouldStopExecutionIfBeforeCommandFailed(feature),
		shouldStopExecutionIfCommandFailed(feature),
	}
}

func shouldExecuteAvailableCommand(feature string) Test {
	fileContent := `
project: Sample Project
commands:
  - build: echo building project
`
	expectedOutput := `
-----    ---------------------
Phase    Command
-----    ---------------------
build    echo building project


-------------------------------[ ` + utils.Colored("build", utils.CYAN) + ` ]--------------------------------
building project

` + successBanner()
	return Test{
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

func shouldShowErrorIfCommandNotFound(feature string) Test {
	fileContent := `
project: Sample Project
commands:
  - build: echo building project
`

	expectedOutput := utils.ColoredB("\nError building execution plan. Command \"random\" not found.", utils.RED) + `
------------------------------------------------------------------------
project: Sample Project
commands:
build | echo building project
------------------------------------------------------------------------
`
	return Test{
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

func shouldExecuteBeforeCommand(feature string) Test {
	fileContent := `
project: Sample Project
before: echo running pre-command
commands:
  - build: echo building project
`
	expectedOutput := `
------    ------------------------
Phase     Command
------    ------------------------
before    echo running pre-command
build     echo building project


-------------------------------[ ` + utils.Colored("before", utils.CYAN) + ` ]-------------------------------
running pre-command
-------------------------------[ ` + utils.Colored("build", utils.CYAN) + ` ]--------------------------------
building project

` + successBanner()
	return Test{
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

func shouldExecuteAfterCommand(feature string) Test {
	fileContent := `
project: Sample Project
after: echo running post-command
commands:
  - build: echo building project
`
	expectedOutput := `
-----    -------------------------
Phase    Command
-----    -------------------------
build    echo building project
after    echo running post-command


-------------------------------[ ` + utils.Colored("build", utils.CYAN) + ` ]--------------------------------
building project
-------------------------------[ ` + utils.Colored("after", utils.CYAN) + ` ]--------------------------------
running post-command

` + successBanner()
	return Test{
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

func shouldExecuteBeforeAndAfterCommand(feature string) Test {
	fileContent := `
project: Sample Project
before: echo running pre-command
after: echo running post-command
commands:
  - build: echo building project
`
	expectedOutput := `
------    -------------------------
Phase     Command
------    -------------------------
before    echo running pre-command
build     echo building project
after     echo running post-command


-------------------------------[ ` + utils.Colored("before", utils.CYAN) + ` ]-------------------------------
running pre-command
-------------------------------[ ` + utils.Colored("build", utils.CYAN) + ` ]--------------------------------
building project
-------------------------------[ ` + utils.Colored("after", utils.CYAN) + ` ]--------------------------------
running post-command

` + successBanner()
	return Test{
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

func shouldStopExecutionIfBeforeCommandFailed(feature string) Test {
	fileContent := `
project: Sample Project
before: exit 10
after: echo running post-command
commands:
  - build: echo building project
`
	expectedOutput := `
------    -------------------------
Phase     Command
------    -------------------------
before    exit 10
build     echo building project
after     echo running post-command


-------------------------------[ ` + utils.Colored("before", utils.CYAN) + ` ]-------------------------------

` + utils.ColoredB("Execution failed during phase \"before\" - Execution of the script \"exit 10\" returned non-zero exit code : 10", utils.RED) + `
`
	return Test{
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

func shouldStopExecutionIfCommandFailed(feature string) Test {
	fileContent := `
project: Sample Project
before: echo running pre-command
after: echo running post-command
commands:
  - build: invalid_command
`
	expectedOutput := `
------    -------------------------
Phase     Command
------    -------------------------
before    echo running pre-command
build     invalid_command
after     echo running post-command


-------------------------------[ ` + utils.Colored("before", utils.CYAN) + ` ]-------------------------------
running pre-command
-------------------------------[ ` + utils.Colored("build", utils.CYAN) + ` ]--------------------------------

` + utils.ColoredB("Execution failed during phase \"build\" - Execution of the script \"invalid_command\" returned non-zero exit code : 127", utils.RED) + `
`
	return Test{
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
