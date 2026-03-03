package testing

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"

	utils2 "github.com/gopinath-langote/1build/cmd/utils"

	"github.com/gopinath-langote/1build/testing/fixtures"
	"github.com/gopinath-langote/1build/testing/utils"
)

var binaryName = "1build"
var binaryPath string
var testDirectory string

func TestAll(t *testing.T) {
	Tests := fixtures.GetFixtures()
	for _, tt := range Tests {
		testResourceDirectory := utils.RecreateTestResourceDirectory(testDirectory)

		t.Run(".:."+tt.Feature+".:."+tt.Name, func(t *testing.T) {
			if tt.Setup != nil {
				_ = tt.Setup(testResourceDirectory)
			}
			var args []string
			if tt.CmdArgs != nil {
				args = tt.CmdArgs(testResourceDirectory)
			}
			cmd := exec.Command(binaryPath, args...)
			cmd.Dir = testResourceDirectory
			out, err := cmd.CombinedOutput()
			actualExitCode := 0
			if exitErr, ok := err.(*exec.ExitError); ok {
				actualExitCode = exitErr.ExitCode()
			}
			if tt.ExpectedExitCode != 0 {
				assert.Equal(t, tt.ExpectedExitCode, actualExitCode, "exit code mismatch")
			}
			_ = tt.Assertion(testResourceDirectory, string(out), t)
			if tt.Teardown != nil {
				_ = tt.Teardown(testResourceDirectory)
			}
		})

	}

}

func TestMain(m *testing.M) {
	testDir, _ := utils.CreateTempDir()
	testDirectory = testDir

	binaryPath = testDir + "/" + binaryName
	buildBinary(binaryPath)

	fmt.Println(utils2.Dash() + "\nBinary Path:- '" + binaryPath + "'\n" + utils2.Dash())

	exitCode := m.Run()

	utils.RemoveAllFilesFromDir(testDir)
	os.Exit(exitCode)
}

func buildBinary(path string) {
	err := os.Chdir("..")
	if err != nil {
		fmt.Printf("could not make binary for %s: %v", binaryName, err.Error())
		os.Exit(1)
	}

	getDep := exec.Command("go", "build", "-o", path)
	if err := getDep.Run(); err != nil {
		fmt.Printf("could not make binary for %s: %v", binaryName, err.Error())
		os.Exit(1)
	}
}
