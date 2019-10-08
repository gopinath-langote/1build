package fixtures

import (
	"testing"

	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
)

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

`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteAvailableCommand",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
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

`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteBeforeCommand",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
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

`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteAfterCommand",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
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

`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteBeforeAndAfterCommand",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
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
`
	return Test{
		Feature: feature,
		Name:    "shouldStopExecutionIfBeforeCommandFailed",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput) &&
				assertFailureMessage(t, actualOutput, "before", "10") &&
				assertFailureBanner(t, actualOutput)

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
`
	return Test{
		Feature: feature,
		Name:    "shouldStopExecutionIfCommandFailed",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedOutput) &&
				assertFailureMessage(t, actualOutput, "build", "127") &&
				assertFailureBanner(t, actualOutput)

		},
	}
}

func assertSuccessBanner(t *testing.T, actualOutput string) bool {
	bannerOutput := utils.ColoredB("SUCCESS", utils.CYAN) + " - Total Time"
	return assert.Contains(t, actualOutput, bannerOutput)
}

func assertFailureMessage(t *testing.T, actualOutput string, phase string, exitCode string) bool {
	errorMessage := utils.Colored("\nExecution failed in phase '"+phase+"' – exit code: "+exitCode, utils.RED)
	return assert.Contains(t, actualOutput, errorMessage)
}

func assertFailureBanner(t *testing.T, actualOutput string) bool {
	bannerOutput := utils.ColoredB("FAILURE", utils.RED) + " - Total Time"
	return assert.Contains(t, actualOutput, bannerOutput)
}
