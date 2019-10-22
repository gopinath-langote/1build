package utils

import (
	"github.com/gopinath-langote/1build/testing/def"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var regex = regexp.MustCompile(ansi)

// CreateConfigFile creates a config file
func CreateConfigFile(dir string, content string) error {
	return ioutil.WriteFile(dir+"/"+def.ConfigFileName, []byte(content), 0777)
}

// CreateTempDir created temporary directory
func CreateTempDir() (string, error) {
	return ioutil.TempDir("", "onebuild_test")
}

// RemoveAllFilesFromDir Cleans up the directory
func RemoveAllFilesFromDir(dir string) {
	_ = os.RemoveAll(dir)
}

// RecreateTestResourceDirectory cleans up test resources and recreates it
func RecreateTestResourceDirectory(dir string) string {
	restResourceDirectory := dir + "/resources"
	RemoveAllFilesFromDir(restResourceDirectory)
	_ = os.Mkdir(restResourceDirectory, 0777)
	return restResourceDirectory
}

const (
	// MaxOutputWidth is the number of spaces to use on a console
	MaxOutputWidth = 72
)

// PlainBanner return dashes with fixed length - 72
func PlainBanner() string {
	return strings.Repeat("-", MaxOutputWidth)
}

// AssertContains checks two string without ANSI colors
func AssertContains(t *testing.T, actualOutput string, expectedOutput string) bool {
	actualOutput = regex.ReplaceAllString(actualOutput, "")
	return assert.Contains(t, actualOutput, expectedOutput)
}
