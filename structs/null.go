package structs

import (
	"database/sql"
	"encoding/json"
	"time"
)

type NullTime struct {
	sql.NullTime
}

func (response NullTime) MarshalJSON() ([]byte, error) {
	if response.Valid {
		return json.Marshal(response.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (response *NullTime) UnmarshalJSON(data []byte) error {
	var t *time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	if t != nil {
		response.Valid = true
		response.Time = *t
	} else {
		response.Valid = false
	}
	return nil
}
