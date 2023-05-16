package schema

import (
	"testing"
)

func TestReminderValidation(t *testing.T) {
	testCases := []struct {
		reminder Reminder
		valid    bool
	}{
		{
			reminder: Reminder{Message: "test 1"},
			valid:    false,
		},
		{
			reminder: Reminder{Message: "test 2", Time: "15:31"},
			valid:    true,
		},
		{
			reminder: Reminder{Message: "test 2", Time: "15-31"},
			valid:    false,
		},
	}

	for _, tc := range testCases {
		err := tc.reminder.IsValid()
		valid := err == nil
		if valid != tc.valid {
			t.Errorf("Test case is failed\nReminder: '%v'\nvalid: '%v'\nexpected: '%v'\nerror: '%v'", tc.reminder, valid, tc.valid, err)
		}
	}
}
