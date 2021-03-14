package gsd

import (
	"encoding/json"
	"strings"
	"time"
)

// Helper for parsing dates without times
type Date time.Time

// Implement Marshaler and Unmarshaler interface
func (j *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = Date(t)
	return nil
}

func (j Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

// Maybe a Format function for printing your date
func (j Date) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
