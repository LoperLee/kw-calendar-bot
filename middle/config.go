package middle

import (
	"encoding/json"
	"os"
)

/*
	@file	config.go
	@author	LoperLee
	@date	2019-11-25
*/

type ConnectInfo struct {
	User   string
	Pass   string
	Addr   string
	Port   int
	DbName string
}

/*
	@author	LoperLee
	@date	2019-11-25
	@brief	Get connection info from config.json
	@return	ConnectInfo
*/
func GetConnectionInfo() ConnectInfo {
	conn := ConnectInfo{}
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		panic("Can`t find config file")
	}
	json.NewDecoder(file).Decode(&conn)
	return conn
}
