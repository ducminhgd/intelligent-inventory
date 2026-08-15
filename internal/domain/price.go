package domain

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Price is a monetary value, stored as DECIMAL in the database. It serializes
// to JSON as a string (rather than a number) so clients never see a float64's
// decimal representation.
type Price float64

// MarshalJSON marshals p as a JSON string.
func (p Price) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatFloat(float64(p), 'f', -1, 64))
}

// UnmarshalJSON unmarshals p from either a JSON number or a JSON string.
func (p *Price) UnmarshalJSON(data []byte) error {
	s := strings.TrimSpace(string(data))
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return err
		}
		s = str
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid price %s: %w", string(data), err)
	}
	*p = Price(f)
	return nil
}
