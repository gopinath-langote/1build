package utils

import (
	"github.com/gopinath-langote/1build/testing/def"
	"io/ioutil"
	"os"
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
