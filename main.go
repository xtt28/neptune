package main

import (
	"fmt"
	"os"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"github.com/xtt28/neptune/commands"
	"github.com/xtt28/neptune/config"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/handlers"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	db, err := database.ConnectSQLite3("instancedata/data.sqlite3")
	if err != nil {
		log.Fatalf("could not connect to persistent data storage: %s", err.Error())
	}

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	conf, err := readConfig(log)
	if err != nil {
		log.Fatalln(err)
	}
	config.LoadNeptuneConfig()

	srv := conf.New()
	srv.CloseOnProgramEnd()

	srv.World().SetTime(18_000)
	srv.World().StopTime()

	srv.World().StopWeatherCycle()
	srv.World().StopRaining()
	srv.World().StopThundering()

	commands.RegisterCommands()

	srv.Listen()
	for srv.Accept(handlers.PlayerHandler(db)) {
	}
}

// readConfig reads the configuration from the config.toml file, or creates the
// file if it does not yet exist.
func readConfig(log server.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return c.Config(log)
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c.Config(log)
}
