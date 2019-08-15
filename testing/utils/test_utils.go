package utils

import (
	"github.com/gopinath-langote/1buildgo/testing/def"
	"io/ioutil"
	"os"
)

func CreateConfigFile(dir string, content string) error {
	return ioutil.WriteFile(dir+"/"+def.ConfigFileName, []byte(content), 0777)
}

func CreateTempDir() (string, error) {
	return ioutil.TempDir("", "onebuild_test")
}

func RemoveAllFilesFromDir(dir string) {
	_ = os.RemoveAll(dir)
}

func RecreateTestResourceDirectory(dir string) string {
	restResourceDirectory := dir + "/resources"
	RemoveAllFilesFromDir(restResourceDirectory)
	_ = os.Mkdir(restResourceDirectory, 0777)
	return restResourceDirectory
}
