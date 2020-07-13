package common

import (
	util "github.com/Floor-Gang/utilpkg/config"
	"github.com/google/uuid"
	"log"
	"os"
	"strings"
)

type Config struct {
	DB        DBConfig        `yaml:"database"`
	Discord   DiscordConfig   `yaml:"discord_bot"`
	WebServer WebServerConfig `yaml:"webserver"`
}

type DiscordConfig struct {
	Token  string `yaml:"token"`
	Prefix string `yaml:"prefix"`
	// Guild to serve
	Guild string `yaml:"guild"`
	// Roles to check for in a user
	Whitelist []string `yaml:"whitelisted_roles"`
	// Administrators that can manage the bot (this also counts as a whitelist)
	AdminRoles []string `yaml:"admin_roles"`
}

type DBConfig struct {
	Location string `yaml:"location"`
}

type WebServerConfig struct {
	Port  int    `yaml:"port"`
	Token string `yaml:"token"`
}

func GetConfig(configPath string) (config Config) {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalln("Failed to get current working directory?\n" + err.Error())
	}

	// default configuration, all the other attributes are blank
	config = Config{
		DB: DBConfig{
			Location: cwd + "/mcauth.db",
		},
		Discord: DiscordConfig{
			Prefix:     ".mc",
			Whitelist:  []string{"1", "2", "3"},
			AdminRoles: []string{"1", "2", "3"},
		},
		WebServer: WebServerConfig{
			Port:  5353,
			Token: strings.Replace(uuid.New().String(), "-", "", 4),
		},
	}

	err = util.GetConfig(configPath, &config)

	if err != nil {
		log.Fatalln(err)
	}

	return config
}
