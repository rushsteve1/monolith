package shared

import (
	"encoding/json"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type DatabaseCfg struct {
	UseSqlite bool   `json:"sqlite"`
	Uri       string `json:"uri"`
	User      string `json:"user"`
	Password  string `json:"password"`
}

func (db DatabaseCfg) String() string {
	if db.UseSqlite {
		return db.Uri
	}
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
	Addr       string `json:"addr"`
	CgiPath    string `json:"cgipath"`
	StaticPath string `json:"staticpath"`
}

type Config struct {
	UseCaddy     bool            `json:"caddy"`
	Database     DatabaseCfg     `json:"database"`
	Overseer     OverseerCfg     `json:"overseer"`
	SwissArmyBot SwissArmyBotCfg `json:"swissarmybot"`
	WebServer    WebServerCfg    `json:"webserver"`
}

func LoadConfig(path string) Config {
	var c Config

	data, err := os.ReadFile(path)
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
