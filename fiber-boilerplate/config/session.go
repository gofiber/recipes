package config

import (
	"github.com/gofiber/session"
	"github.com/gofiber/session/provider/redis"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"time"
)

type SessionConfiguration struct {
	Session_DSN    string
	Session_DB     int
	Session_Lookup string
}

var SessionConfig *SessionConfiguration //nolint:gochecknoglobals

func LoadSessionConfig() {
	loadDefaultSessionConfig()
	ViperConfig.Unmarshal(&SessionConfig)
}

func loadDefaultSessionConfig() {
	ViperConfig.SetDefault("SESSION_DSN", "127.0.0.1:6379")
	ViperConfig.SetDefault("SESSION_DB", 1)
	ViperConfig.SetDefault("SESSION_LOOKUP", "cookie:fiber-boilerplate-session")
}

func LoadSession() {
	LoadSessionConfig()
	provider := redis.New(redis.Config{
		KeyPrefix:   "fiber_boilerplate",
		Addr:        SessionConfig.Session_DSN,
		PoolSize:    8,                //nolint:gomnd
		IdleTimeout: 30 * time.Second, //nolint:gomnd
		DB:          SessionConfig.Session_DB,
	})
	Session = session.New(session.Config{
		Provider: provider,
	})
}
