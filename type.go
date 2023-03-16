package thailandposttracking

import (
	"encoding/json"
	"strings"
	"time"
)

type TPDateTime time.Time // ThailandPost Datetime

// Implement Marshaler and Unmarshaler interface
func (j *TPDateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05+07:00", s)
	if err != nil {
		return err
	}
	*j = TPDateTime(t)
	return nil
}

func (j TPDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Maybe a Format function for printing your date
func (j TPDateTime) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
