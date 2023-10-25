package services

import "errors"

func validateEvent(event string) error {
	valid_events := []string{"ENTER", "LEAVE", "RING"}

	for _, valid_event := range valid_events {
		if event == valid_event {
			return nil
		}
	}
	return errors.New("invalid event")

}

func notify(mobile, event string) error {
	return validateEvent(event)
}
