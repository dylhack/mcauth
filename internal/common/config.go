package common

import (
	util "github.com/Floor-Gang/utilpkg/config"
	"github.com/dhghf/mcauth/internal/common/db"
	googleUUID "github.com/google/uuid"
	"log"
	"strings"
	"time"
)

// Config is all the configurations below combined.
type Config struct {
	DB        db.Config       `yaml:"database"`
	Discord   DiscordConfig   `yaml:"discord_bot"`
	WebServer WebServerConfig `yaml:"webserver"`
}

// DiscordConfig is the Discord bot config.
type DiscordConfig struct {
	Help       string   `yaml:"help_message"`
	Token      string   `yaml:"token"`
	Prefix     string   `yaml:"prefix"`
	Guild      string   `yaml:"guild"`
	Whitelist  []string `yaml:"whitelisted_roles"`
	AdminRoles []string `yaml:"admin_roles"`
}

// WebServerConfig is the configuration attributes for the webserver talking to the plugin.
type WebServerConfig struct {
	Port  int    `yaml:"port"`
	Token string `yaml:"token"`
}

// GetConfig reads and returns the config.yml. A new one will be created if the current one can't
// be found in the CWD.
func GetConfig(configPath string) (config Config) {
	// default configuration, all the other attributes are blank
	config = Config{
		DB: db.Config{
			Host:               "localhost",
			Port:               5432,
			User:               "postgres",
			Password:           "",
			Database:           "",
			MaxConnections:     50,
			MaxIdleConnections: 50,
			ConnLifespan:       time.Hour * 1,
		},
		Discord: DiscordConfig{
			Prefix:     ".mc",
			Whitelist:  []string{"1", "2", "3"},
			AdminRoles: []string{"1", "2", "3"},
		},
		WebServer: WebServerConfig{
			Port:  5353,
			Token: strings.Replace(googleUUID.New().String(), "-", "", 4),
		},
	}

	err := util.GetConfig(configPath, &config)

	if err != nil {
		log.Fatalln(err)
	}

	return config
}
