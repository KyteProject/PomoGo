package storage

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Service interface {
}

type Pomodoro struct {
	gorm.Model
	UUID           string        `gorm:"type:varchar(255);PRIMARY_KEY;NOT NULL;"`
	WorkTime       time.Duration `json:"workTime"`
	ShortBreakTime time.Duration `json:"shortBreakTime"`
	LongBreakTime  time.Duration `json:"longBreakTime"`
	WorkRounds     int32         `gorm:"type:integer;DEFAULT=0;"`
}
