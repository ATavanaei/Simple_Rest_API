package Config

import (
	"encoding/json"
	"flag"
	"os"
	"time"
)

var ProjectPath = "/root/work/src/app/ServerSystemV2/UserManagement/"

type Configuration struct {
	Serve_port      string
	Read_Timeout    time.Duration
	Write_Timeout   time.Duration
	Postgres struct {
		Po_host     string
		Po_port     int
		Po_user     string
		Po_password string
		Po_dbname   string
	}
	Mongo struct{
		Mo_host     string
		Mo_port     int
	}
	Redis struct{
		Server      string
		Port        string
		Password    string
		Db          int
	}
	Kafka struct{
		Server      string
		Port        string
	}
}

func LoadConfig() (Configuration){
	path := ProjectPath + "Config/conf.json"
	c := flag.String("c", path, "configuration file")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}
	return Config
}