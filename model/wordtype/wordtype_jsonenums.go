// generated by jsonenums -type=WordType; DO NOT EDIT

package wordtype

import (
	"encoding/json"
	"fmt"
)

var (
	_WordTypeNameToValue = map[string]WordType{
		"Noun":      Noun,
		"Adjective": Adjective,
		"Verb":      Verb,
		"Unknown":   Unknown,
	}

	_WordTypeValueToName = map[WordType]string{
		Noun:      "Noun",
		Adjective: "Adjective",
		Verb:      "Verb",
		Unknown:   "Unknown",
	}
)

func init() {
	var v WordType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_WordTypeNameToValue = map[string]WordType{
			interface{}(Noun).(fmt.Stringer).String():      Noun,
			interface{}(Adjective).(fmt.Stringer).String(): Adjective,
			interface{}(Verb).(fmt.Stringer).String():      Verb,
			interface{}(Unknown).(fmt.Stringer).String():   Unknown,
		}
	}
}

// MarshalJSON is generated so WordType satisfies json.Marshaler.
func (r WordType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _WordTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid WordType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so WordType satisfies json.Unmarshaler.
func (r *WordType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("WordType should be a string, got %s", data)
	}
	v, ok := _WordTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid WordType %q", s)
	}
	*r = v
	return nil
}