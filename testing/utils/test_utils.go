package utils

import (
	"github.com/gopinath-langote/1build/testing/def"
	"io/ioutil"
	"os"
)

func createConfigFile(dir string, content string) error {
	return ioutil.WriteFile(dir+"/"+def.ConfigFileName, []byte(content), 0777)
}

func createTempDir() (string, error) {
	return ioutil.TempDir("", "onebuild_test")
}

func removeAllFilesFromDir(dir string) {
	_ = os.RemoveAll(dir)
}

func recreateTestResourceDirectory(dir string) string {
	restResourceDirectory := dir + "/resources"
	removeAllFilesFromDir(restResourceDirectory)
	_ = os.Mkdir(restResourceDirectory, 0777)
	return restResourceDirectory
}
