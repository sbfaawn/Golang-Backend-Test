package helper

import "time"

const (
	DateFormatLayout = "2006-01-02"
)

func IsValidDate(date string) error {
	_, err := time.Parse(DateFormatLayout, date)
	if err != nil {
		return err
	}
	return nil
}

func ConvertToTimeObject(date string) time.Time {
	time, _ := time.Parse(DateFormatLayout, date)
	return time
}
