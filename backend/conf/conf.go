package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type GbDbConfig struct {
	BasePath         string `json:"basePath"`
	DbDirPath        string `json:"dbDirPath"`
	LogDirPath       string `json:"logDirPath"`
	BufferPageMemory uint64 `json:"bufferPageMemory"`
}

type ServerStartConfig struct {
	confFilePath string
	GbDbConfig
}

var _ssc *ServerStartConfig

func loadGbDbConfig(ssc *ServerStartConfig) {
	data, err := ioutil.ReadFile(ssc.confFilePath)
	if err != nil {
		panic(err)
	}
	var cnf = GbDbConfig{}
	err = json.Unmarshal(data, &cnf)
	if !filepath.IsAbs(cnf.DbDirPath) {
		cnf.DbDirPath = filepath.Join(cnf.BasePath, cnf.DbDirPath)
	}
	if !filepath.IsAbs(cnf.LogDirPath) {
		cnf.LogDirPath = filepath.Join(cnf.BasePath, cnf.LogDirPath)
	}
	if err != nil {
		panic(err)
	}
	ssc.GbDbConfig = cnf
}

func GetServerStartConfig() *ServerStartConfig {
	ssc := &ServerStartConfig{}
	ssc.confFilePath = os.Getenv("GBDB_CONFIG")
	if ssc.confFilePath == "" {
		ssc.confFilePath = os.Getenv("GBDB_DEFAULT_CONFIG")
	}
	loadGbDbConfig(ssc)
	_ssc = ssc
	return ssc
}

func GetConfig() (ssc *ServerStartConfig) {
	if _ssc != nil {
		return _ssc
	}
	return GetServerStartConfig()
}

func scanDirFiles(dir string) (fArray []string) {
	dirList, e := ioutil.ReadDir(dir)
	if e != nil {
		panic(e)
		return
	}
	for _, v := range dirList {
		fArray = append(fArray, dir+v.Name())
	}
	return
}
