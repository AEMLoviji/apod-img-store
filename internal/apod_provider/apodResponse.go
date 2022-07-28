package apodprovider

import (
	"encoding/json"
	"strings"
	"time"
)

type ApodJsonDate time.Time

func (j *ApodJsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = ApodJsonDate(t)
	return nil
}

func (j ApodJsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

type ApodResponse struct {
	Title     string       `json:"title"`
	Copyright string       `json:"copyright"`
	Url       string       `json:"url"`
	Date      ApodJsonDate `json:"date"`
}
