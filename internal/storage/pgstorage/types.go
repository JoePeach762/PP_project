package pgstorage

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// StringSlice — вспомогательный тип для корректного сканирования JSONB в []string
type StringSlice []string

func (s StringSlice) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return json.Marshal(s)
}

func (s *StringSlice) Scan(value interface{}) error {
	if value == nil {
		*s = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, s)
	case string:
		return json.Unmarshal([]byte(v), s)
	default:
		return errors.New("failed to scan StringSlice: unsupported type")
	}
}
