package fixtures

import (
	"testing"

	"github.com/gopinath-langote/1build/testing/utils"
)

func featureFlagTestData() []Test {
	feature := "flag"
	return []Test{
		shouldExecuteCommandWithquietFlag(feature),
		shouldExecuteBeforeCommandWithquietFlag(feature),
		shouldExecuteAfterCommandWithquietFlag(feature),
		shouldExecuteBeforeAndAfterCommandWithquietFlag(feature),
		shouldStopExecutionIfBeforeCommandFailedWithquietFlag(feature),
		shouldStopExecutionIfCommandFailedWithquietFlag(feature),
	}
}

func shouldExecuteCommandWithquietFlag(feature string) Test {
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


-------------------------------[ ` + "build" + ` ]--------------------------------

`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteCommandWithquietFlag",
		CmdArgs: []string{"build", "--quiet"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
		},
	}
}

func shouldExecuteBeforeCommandWithquietFlag(feature string) Test {
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


-------------------------------[ ` + "before" + ` ]-------------------------------
-------------------------------[ ` + "build" + ` ]--------------------------------

`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteBeforeCommandWithquietFlag",
		CmdArgs: []string{"build", "--quiet"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
		},
	}
}

func shouldExecuteAfterCommandWithquietFlag(feature string) Test {
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


-------------------------------[ ` + "build" + ` ]--------------------------------
-------------------------------[ ` + "after" + ` ]--------------------------------

`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteAfterCommandWithquietFlag",
		CmdArgs: []string{"build", "--quiet"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
		},
	}
}

func shouldExecuteBeforeAndAfterCommandWithquietFlag(feature string) Test {
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


-------------------------------[ ` + "before" + ` ]-------------------------------
-------------------------------[ ` + "build" + ` ]--------------------------------
-------------------------------[ ` + "after" + ` ]--------------------------------

`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteBeforeAndAfterCommandWithquietFlag",
		CmdArgs: []string{"build", "--quiet"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {

			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
		},
	}
}

func shouldStopExecutionIfBeforeCommandFailedWithquietFlag(feature string) Test {
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


-------------------------------[ ` + "before" + ` ]-------------------------------

`
	return Test{
		Feature: feature,
		Name:    "shouldStopExecutionIfBeforeCommandFailedWithquietFlag",
		CmdArgs: []string{"build", "--quiet"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertFailureMessageNone(t, actualOutput, "before", "10") &&
				assertFailureBanner(t, actualOutput)

		},
	}
}

func shouldStopExecutionIfCommandFailedWithquietFlag(feature string) Test {
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


-------------------------------[ ` + "before" + ` ]-------------------------------
-------------------------------[ ` + "build" + ` ]--------------------------------

`
	return Test{
		Feature: feature,
		Name:    "shouldStopExecutionIfCommandFailedWithquietFlag",
		CmdArgs: []string{"build", "--quiet"},
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertFailureMessageNone(t, actualOutput, "build", "127") &&
				assertFailureBanner(t, actualOutput)
		},
	}
}
