package models

type OfficesWorkingDay struct {
	ID       uint  `gorm:"primaryKey" json:"id"`
	OfficeID uint  `json:"office_id"`
	Day      uint8 `json:"day"`
}

func GetDayIndex(day string) uint8 {
	switch day {
	case "Monday":
		return 1
	case "Tuesday":
		return 2
	case "Wednesday":
		return 3
	case "Thursday":
		return 4
	case "Friday":
		return 5
	case "Saturday":
		return 6
	case "Sunday":
		return 7
	}
	return 0
}
