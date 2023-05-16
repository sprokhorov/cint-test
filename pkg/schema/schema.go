/*
The schema package share the schemas between the packages.
*/
package schema

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Reminder is single reminder object.
type Reminder struct {
	Message string `json:"message" validate:"required"`
	Time    string `json:"time" validate:"required"`
}

// IsValid returns error if Reminder is invalid.
func (r *Reminder) IsValid() error {
	val := validator.New()
	if err := val.Struct(*r); err != nil {
		return err
	}
	pattern := `^([0-9]{2}):([0-9]{2})$`
	if ok, _ := regexp.Match(pattern, []byte(r.Time)); !ok {
		return fmt.Errorf("reminder.time mismatch the regexp pattern '%s'", pattern)
	}
	return nil
}
