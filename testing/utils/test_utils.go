package utils

import (
	"github.com/gopinath-langote/1build/testing/def"
	"github.com/logrusorgru/aurora"
	"io/ioutil"
	"os"
	"strings"
)

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
	MaxOutputWidth = 72
)

// OneBuildColor represents type for color enum
type OneBuildColor int

const (
	// CYAN is 1build's default color standard
	CYAN OneBuildColor = 0

	// RED is used in failure messages
	RED  OneBuildColor = 1
)

// PlainBanner return dashes with fixed length - 72
func PlainBanner() string {
	return strings.Repeat("-", MaxOutputWidth)
}

// ColoredB return text in color with bold format
func ColoredB(text string, color OneBuildColor) string {
	return colorize(text, color).Bold().String()
}

// Colored return text in color
func Colored(text string, color OneBuildColor) string {
	return colorize(text, color).String()
}

func colorize(text string, color OneBuildColor) aurora.Value {
	var coloredText aurora.Value
	if color == CYAN {
		coloredText = aurora.BrightCyan(text)
	} else {
		coloredText = aurora.Red(text)
	}
	return coloredText
}
