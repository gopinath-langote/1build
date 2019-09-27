package fixtures

import (
	"os"
	"testing"

	"github.com/gopinath-langote/1build/testing/utils"
	"github.com/stretchr/testify/assert"
)

func featureExecuteCmdTestData() []Test {
	feature := "exec"
	os.Setenv("DISABLE_1B_COLORS", "false")
	return []Test{
		shouldExecuteAvailableCommand(feature),
		shouldShowErrorIfCommandNotFound(feature),
		shouldExecuteBeforeCommand(feature),
		shouldExecuteAfterCommand(feature),
		shouldExecuteBeforeAndAfterCommand(feature),
		shouldStopExecutionIfBeforeCommandFailed(feature),
		shouldStopExecutionIfCommandFailed(feature),
		shouldPrintBannerOnEachPhase(feature),
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

--------------------------------[ build ]-------------------------------
building project

SUCCESS`
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

	expectedOutput := "\nError building execution plan. Command \"random\" not found." + `
--------------------------------------------------
project: Sample Project
commands:
build | echo building project
--------------------------------------------------
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

-------------------------------[ before ]-------------------------------
running pre-command

--------------------------------[ build ]-------------------------------
building project

SUCCESS
`
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

--------------------------------[ build ]-------------------------------
building project

--------------------------------[ after ]-------------------------------
running post-command

SUCCESS
`
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

-------------------------------[ before ]-------------------------------
running pre-command

--------------------------------[ build ]-------------------------------
building project

--------------------------------[ after ]-------------------------------
running post-command

SUCCESS
`
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

-------------------------------[ before ]-------------------------------

-----------------------------------------------------------------------------------------------------------
Execution failed during phase "before" - Execution of the script "exit 10" returned non-zero exit code : 10
-----------------------------------------------------------------------------------------------------------
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

-------------------------------[ before ]-------------------------------
running pre-command

--------------------------------[ build ]-------------------------------

-------------------------------------------------------------------------------------------------------------------
Execution failed during phase "build" - Execution of the script "invalid_command" returned non-zero exit code : 127
-------------------------------------------------------------------------------------------------------------------
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

func shouldPrintBannerOnEachPhase(feature string) Test {
	fileContent := `
project: Sample Project
before: echo running pre-command
after: echo running post-command
commands:
  - build: echo build
`
	expectedBeforeBanner := `----------------------------[ before ]---------------------------`
	expectedBuildBanner := `----------------------------[ build ]----------------------------`
	expectedAfterBanner := `----------------------------[ after ]---------------------------`
	return Test{
		Feature: feature,
		Name:    "shouldPrintBannerOnAllPhases",
		CmdArgs: []string{"build"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return assert.Contains(t, actualOutput, expectedBeforeBanner) &&
				assert.Contains(t, actualOutput, expectedBuildBanner) &&
				assert.Contains(t, actualOutput, expectedAfterBanner)
		},
	}
}
