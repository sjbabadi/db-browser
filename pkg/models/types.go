package models

import "errors"

type NullString string

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	strVal, ok := value.(string)
	if !ok {
		return errors.New("Invalid string in database")
	}

	*s = NullString(strVal)
	return nil
}

/*
{"id": 3}

{ {"Int": 3, "Valid": true} }
*/
