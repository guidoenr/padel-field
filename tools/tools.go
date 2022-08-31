package tools

import (
	"github.com/guidoenr/padel-field/logger"
	"time"
)

// GetTodayDate returns the current local datetime in America/Buenos_Aires
func GetTodayDate() (time.Time, *time.Location) {
	location, err := time.LoadLocation("America/Buenos_Aires")
	if err != nil {
		logger.Logerror.Printf("error loading the location: %v", err)
	}
	todayDate, err := time.ParseInLocation("02-01-2006", time.Now().Format("02-01-2006"), location)
	if err != nil {
		logger.Logerror.Printf("error getting localTime: %v", err)
	}
	return todayDate, location
}
