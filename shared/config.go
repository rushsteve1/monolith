package shared

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

type DatabaseCfg struct {
	Uri      string `json:"uri"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (db DatabaseCfg) String() string {
	return fmt.Sprintf("host=%s user=%s password=%s sslmode=disable", db.Uri, db.User, db.Password)
}

type OverseerCfg struct {
	Addr string `json:"addr"`
	Fcgi bool   `json:"fcgi"`
	Rpc  bool   `json:"rpc"`
}

type SwissArmyBotCfg struct {
	Addr  string `json:"addr"`
	Token string `json:"token"`
}

type WebServerCfg struct {
	Addr string `json:"addr"`
}

type Config struct {
	Database     DatabaseCfg     `json:"database"`
	Overseer     OverseerCfg     `json:"overseer"`
	SwissArmyBot SwissArmyBotCfg `json:"swissarmybot"`
	WebServer    WebServerCfg    `json:"webserver"`
}

func LoadConfig(path string) Config {
	var c Config

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	json.Unmarshal(data, &c)
	return c
}

func ConfigFromArgs() Config {
	if len(os.Args) < 2 {
		log.Fatal("Path to config file required")
	}
	return LoadConfig(os.Args[1])
}
