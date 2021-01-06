package config

import (
	"github.com/wailsapp/wails"
	"pomogo/backend/localstore"
	"time"
)

//Config is the collection of app settings
type Config struct {
	App        *App
	Pomodoro   *Pomodoro
	Runtime    *wails.Runtime
	Logger     *wails.CustomLogger
	LocalStore *localstore.LocalStore
}

// WailsInit performs setup when Wails is ready
func (c *Config) WailsInit(r *wails.Runtime) error {
	c.Runtime = r
	c.Logger = c.Runtime.Log.New("Config")
	c.Logger.Info("Initialised config.")
	return nil
}

// NewConfig returns a new instance of Config.
func NewConfig() *Config {
	c := &Config{}
	c.LocalStore = localstore.NewLocalStore()

	return c
}
