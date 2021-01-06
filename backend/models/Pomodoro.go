package models

import "time"

// Pomodoro is the config for the pomodoro timer functionality
type Pomodoro struct {
	ID             uint `gorm:"AUTO_INCREMENT"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	WorkTime       time.Duration `json:"workTime"`
	ShortBreakTime time.Duration `json:"shortBreakTime"`
	LongBreakTime  time.Duration `json:"longBreakTime"`
	WorkRounds     int           `json:"workRounds"`
}
