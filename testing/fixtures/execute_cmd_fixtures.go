package fixtures

import (
	"os"
	"testing"

	"github.com/gopinath-langote/1build/testing/utils"
)

func featureExecuteCmdTestData() []Test {
	feature := "exec"

	return []Test{
		shouldExecuteAvailableCommand(feature),
		shouldExecuteAvailableCommandFromSpecifiedFile(feature),
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
-------------------------------[ build ]--------------------------------
building project

`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteAvailableCommand",
		CmdArgs: Args("build"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
		},
	}
}

func shouldExecuteAvailableCommandFromSpecifiedFile(feature string) Test {
	fileContent := `
project: Sample Project
commands:
  - build: echo building project
`
	expectedOutput := `
-------------------------------[ build ]--------------------------------
building project

`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteAvailableCommandFromSpecifiedFile",
		CmdArgs: func(dir string) []string {
			return []string{"build", "-f", dir + "/some-dir/some-config.yaml"}
		},
		Setup: func(dir string) error {
			_ = os.MkdirAll(dir+"/some-dir/", 0750)
			return utils.CreateConfigFileWithName(dir+"/some-dir", "some-config.yaml", fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
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

	expectedOutput := `
Error: Command "random" not found.
------------------------------------------------------------------------
project: Sample Project
commands:
build | echo building project
------------------------------------------------------------------------
`
	return Test{
		Feature: feature,
		Name:    "shouldShowErrorIfCommandNotFound",
		CmdArgs: Args("random"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput)
		},
	}
}

func shouldExecuteBeforeCommand(feature string) Test {
	fileContent := `
project: Sample Project
beforeAll: echo running pre-command
commands:
  - build: echo building project
`
	expectedOutput := `beforeAll: echo running pre-command
-----------------------------[ beforeAll ]------------------------------
running pre-command
Executing command: build
  command: echo building project
-------------------------------[ build ]--------------------------------
building project

------------------------------------------------------------------------
SUCCESS - Total Time: 00s
------------------------------------------------------------------------
`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteBeforeCommand",
		CmdArgs: Args("build"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
		},
	}
}

func shouldExecuteAfterCommand(feature string) Test {
	fileContent := `
project: Sample Project
afterAll: echo running post-command
commands:
  - build: echo building project
`
	expectedOutput := `Executing command: build
  command: echo building project
-------------------------------[ build ]--------------------------------
building project
afterAll: echo running post-command
------------------------------[ afterAll ]------------------------------
running post-command

------------------------------------------------------------------------
SUCCESS - Total Time: 00s
------------------------------------------------------------------------
`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteAfterCommand",
		CmdArgs: Args("build"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
		},
	}
}

func shouldExecuteBeforeAndAfterCommand(feature string) Test {
	fileContent := `
project: Sample Project
beforeAll: echo running pre-command
afterAll: echo running post-command
commands:
  - build: echo building project
`
	expectedOutput := `beforeAll: echo running pre-command
-----------------------------[ beforeAll ]------------------------------
running pre-command
Executing command: build
  command: echo building project
-------------------------------[ build ]--------------------------------
building project
afterAll: echo running post-command
------------------------------[ afterAll ]------------------------------
running post-command

------------------------------------------------------------------------
SUCCESS - Total Time: 00s
------------------------------------------------------------------------
`
	return Test{
		Feature: feature,
		Name:    "shouldExecuteBeforeAndAfterCommand",
		CmdArgs: Args("build"),
		Setup: func(dir string) error {
			return utils.CreateConfigFile(dir, fileContent)
		},
		Assertion: func(dir string, actualOutput string, t *testing.T) bool {
			return utils.AssertContains(t, actualOutput, expectedOutput) &&
				assertSuccessBanner(t, actualOutput)
		},
	}
}

func shouldStopExecutionIfBeforeCommandFailed(feature string) Test {
	fileContent := `
project: Sample Project
beforeAll: exit 10
afterAll: echo running post-command
commands:
  - build: echo building project
`
	expectedOutput := `beforeAll: exit 10
-----------------------------[ beforeAll ]------------------------------

`
	return Test{
		Feature: feature,
		Name:    "shouldStopExecutionIfBeforeCommandFailed",
		CmdArgs: Args("build"),
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

func shouldStopExecutionIfCommandFailed(feature string) Test {
	fileContent := `
project: Sample Project
beforeAll: echo running pre-command
afterAll: echo running post-command
commands:
  - build: invalid_command
`
	expectedOutput := `beforeAll: echo running pre-command
-----------------------------[ beforeAll ]------------------------------
running pre-command
Executing command: build
  command: invalid_command
-------------------------------[ build ]--------------------------------

`
	return Test{
		Feature: feature,
		Name:    "shouldStopExecutionIfCommandFailed",
		CmdArgs: Args("build"),
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

func assertSuccessBanner(t *testing.T, actualOutput string) bool {
	return utils.AssertContains(t, actualOutput, "SUCCESS - Total Time")
}

func assertFailureMessage(t *testing.T, actualOutput string, phase string, exitCode string) bool {
	errorText := "\nExecution failed in phase '" + phase + "' – exit code: " + exitCode
	return utils.AssertContains(t, actualOutput, errorText)
}

func assertFailureMessageNone(t *testing.T, actualOutput string, phase string, exitCode string) bool {
	errorText := "\nExecution failed in phase '" + phase + "' – exit code: " + exitCode
	return utils.AssertNotContains(t, actualOutput, errorText)
}

func assertFailureBanner(t *testing.T, actualOutput string) bool {
	return utils.AssertContains(t, actualOutput, "FAILURE - Total Time")
}
