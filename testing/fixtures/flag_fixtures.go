package fixtures

import (
	"testing"

	"github.com/gopinath-langote/1build/testing/utils"
)

func featureFlagTestData() []Test {
	feature := "flag"
	return []Test{
		shouldExecuteCommandWithquietFlag(feature),
		shouldExecuteBeforeAllCommandWithquietFlag(feature),
		shouldExecuteAfterAllCommandWithquietFlag(feature),
		shouldExecuteBeforeAllAndAfterAllCommandWithquietFlag(feature),
		shouldStopExecutionIfBeforeAllCommandFailedWithquietFlag(feature),
		shouldStopExecutionIfCommandFailedWithquietFlag(feature),
	}
}

func shouldExecuteCommandWithquietFlag(feature string) Test {
	fileContent := `
project: Sample Project
commands:
  - build: echo building project
`
	expectedOutput := `Executing command: build
  command: echo building project
-------------------------------[ build ]--------------------------------

------------------------------------------------------------------------
SUCCESS - Total Time: 00s
------------------------------------------------------------------------
`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteCommandWithquietFlag",
		CmdArgs: Args("build", "--quiet"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
		},
	}
}

func shouldExecuteBeforeAllCommandWithquietFlag(feature string) Test {
	fileContent := `
project: Sample Project
beforeAll: echo running pre-command
commands:
  - build: echo building project
`
	expectedOutput := `beforeAll: echo running pre-command
-----------------------------[ beforeAll ]------------------------------
Executing command: build
  command: echo building project
-------------------------------[ build ]--------------------------------

------------------------------------------------------------------------
SUCCESS - Total Time: 00s
------------------------------------------------------------------------
`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteBeforeCommandWithquietFlag",
		CmdArgs: Args("build", "--quiet"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
		},
	}
}

func shouldExecuteAfterAllCommandWithquietFlag(feature string) Test {
	fileContent := `
project: Sample Project
afterAll: echo running post-command
commands:
  - build: echo building project
`
	expectedOutput := `Executing command: build
  command: echo building project
-------------------------------[ build ]--------------------------------
afterAll: echo running post-command
------------------------------[ afterAll ]------------------------------

------------------------------------------------------------------------
SUCCESS - Total Time: 00s
------------------------------------------------------------------------
`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteAfterCommandWithquietFlag",
		CmdArgs: Args("build", "--quiet"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
		},
	}
}

func shouldExecuteBeforeAllAndAfterAllCommandWithquietFlag(feature string) Test {
	fileContent := `
project: Sample Project
beforeAll: echo running pre-command
afterAll: echo running post-command
commands:
  - build: echo building project
`
	expectedOutput := `beforeAll: echo running pre-command
-----------------------------[ beforeAll ]------------------------------
Executing command: build
  command: echo building project
-------------------------------[ build ]--------------------------------
afterAll: echo running post-command
------------------------------[ afterAll ]------------------------------

------------------------------------------------------------------------
SUCCESS - Total Time: 00s
------------------------------------------------------------------------
`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteBeforeAndAfterCommandWithquietFlag",
		CmdArgs: Args("build", "--quiet"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
		},
	}
}

func shouldStopExecutionIfBeforeAllCommandFailedWithquietFlag(feature string) Test {
	fileContent := `
project: Sample Project
beforeAll: exit 10
afterAll: echo running post-command
commands:
  - build: echo building project
`
	expectedOutput := `beforeAll: exit 10
-----------------------------[ beforeAll ]------------------------------

Execution failed in phase 'beforeAll' – exit code: 10

------------------------------------------------------------------------
FAILURE - Total Time: 00s
------------------------------------------------------------------------
`
	return Test{
		Feature:          feature,
		Name:             "shouldStopExecutionIfBeforeCommandFailedWithquietFlag",
		CmdArgs:          Args("build", "--quiet"),
		ExpectedExitCode: 10,
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertFailureMessage(t, actualOutput, "beforeAll", "10") &&
				assertFailureBanner(t, actualOutput)

		},
	}
}

func shouldStopExecutionIfCommandFailedWithquietFlag(feature string) Test {
	fileContent := `
project: Sample Project
beforeAll: echo running pre-command
afterAll: echo running post-command
commands:
  - build: invalid_command
`
	expectedOutput := `beforeAll: echo running pre-command
-----------------------------[ beforeAll ]------------------------------
Executing command: build
  command: invalid_command
-------------------------------[ build ]--------------------------------

Execution failed in phase 'build' – exit code: 127

------------------------------------------------------------------------
FAILURE - Total Time: 00s
------------------------------------------------------------------------
`
	return Test{
		Feature:          feature,
		Name:             "shouldStopExecutionIfCommandFailedWithquietFlag",
		CmdArgs:          Args("build", "--quiet"),
		ExpectedExitCode: 127,
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertFailureMessage(t, actualOutput, "build", "127") &&
				assertFailureBanner(t, actualOutput)
		},
	}
}
