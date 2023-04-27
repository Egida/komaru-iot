package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadConfig() {
	writeConfig()
	file, err := ioutil.ReadFile("serverConfig.json")
	if err != nil {
		Logger.Fatal("Decoding 'serverConfig.json' failed.")
		return
	}
	err = json.NewDecoder(bytes.NewReader(file)).Decode(&Server)
	if err != nil {
		Logger.Fatal("Decoding 'serverConfig.json' failed.")
		return
	}
	Logger.Info("Loaded 'serverConfig.json' successfully")
}

func writeConfig() {
	if _, err := os.Stat("serverConfig.json"); os.IsNotExist(err) {
		cfg := serverConfig{
			MasterPort:    1337,
			SlavePort:     69,
			SshKey:        "id_rsa",
			MysqlAddress:  "localhost:3306",
			MysqlUser:     "root",
			MysqlPassword: "",
			MysqlDatabase: "cat",
		}
		file, _ := json.MarshalIndent(cfg, "", " ")
		_ = ioutil.WriteFile("serverConfig.json", file, 0644)
	}
}
