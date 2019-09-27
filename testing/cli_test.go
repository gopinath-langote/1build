package testing

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

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
			cmd := exec.Command(binaryPath, tt.CmdArgs...)
			cmd.Dir = testResourceDirectory
			out, _ := cmd.Output()
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

	fmt.Println(utils2.BANNER() + "\nBinary Path:- '" + binaryPath + "'\n" + utils2.BANNER())

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
