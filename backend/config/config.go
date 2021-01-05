package config

import (
	"github.com/wailsapp/wails"
	"pomogo/backend/localstore"
	"time"
)

//App is the application specific config
type App struct {
	Width         int    `json:"appWidth"`
	Height        int    `json:"appHeight"`
	Title         string `json:"appTitle"`
	Colour        string `json:"appColour"`
	Notifications bool   `json:"appNotifications"`
	Sounds        bool   `json:"appSounds"`
}

// Pomodoro is the config for the pomodoro timer functionality
type Pomodoro struct {
	WorkTime       time.Duration `json:"workTime"`
	ShortBreakTime time.Duration `json:"shortBreakTime"`
	LongBreakTime  time.Duration `json:"longBreakTime"`
	WorkRounds     int           `json:"workRounds"`
}

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
