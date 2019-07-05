package srv

import (
	"io/ioutil"
	"encoding/json"
)

type GbdbConfig struct {
	DbDirPath        string `json:dbDirPath`
	LogDirPath       string `json:logDirPath`
	BufferPageMemory uint64 `json:bufferPageMemory`
}

type ServerStartConfig struct {
	GbdbConfig
	dbFiles  []string
	logFiles []string
}

func loadGbdbConfig() (c GbdbConfig) {
	data, err := ioutil.ReadFile("../../gbdbconfig.json")
	if err != nil {
		panic(err)
		return
	}
	var cnf = GbdbConfig{}
	err = json.Unmarshal(data, &cnf)
	if err != nil {
		panic(err)
		return
	}
	return cnf
}

func getSereStartConfig() (ssc ServerStartConfig) {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//	//if err != nil {
	//	//	log.Fatal(err)
	//	//}
	ssc = ServerStartConfig{
		GbdbConfig: loadGbdbConfig(),
	}
	ssc.dbFiles = scanDirFiles(ssc.DbDirPath)
	ssc.logFiles = scanDirFiles(ssc.LogDirPath)
	return
}

func scanDirFiles(dir string) (farray []string) {
	dir_list, e := ioutil.ReadDir(dir)
	if e != nil {
		panic(e)
		return
	}
	for _, v := range dir_list {
		farray = append(farray, dir+v.Name())
	}
	return
}
