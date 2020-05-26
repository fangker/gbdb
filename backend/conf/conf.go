package conf

import (
	"encoding/json"
	"io/ioutil"
)

type GbDbConfig struct {
	DbDirPath        string `json:"dbDirPath"`
	LogDirPath       string `json:"logDirPath"`
	BufferPageMemory uint64 `json:"bufferPageMemory"`
}

type ServerStartConfig struct {
	GbDbConfig
	dbFiles  string
	logFiles string
}

var _ssc *ServerStartConfig

func loadGbDbConfig() (c GbDbConfig) {
	data, err := ioutil.ReadFile("../../gbDbConfig.json")
	if err != nil {
		panic(err)
	}
	var cnf = GbDbConfig{}
	err = json.Unmarshal(data, &cnf)
	if err != nil {
		panic(err)
	}
	return cnf
}

func GetServerStartConfig() (ssc *ServerStartConfig) {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//	//if err != nil {
	//	//	log.Fatal(err)
	//	//}
	if _ssc != nil {
		return _ssc
	}
	ssc = &ServerStartConfig{
		GbDbConfig: loadGbDbConfig(),
	}
	ssc.dbFiles = ssc.DbDirPath
	ssc.logFiles = ssc.LogDirPath
	return
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
