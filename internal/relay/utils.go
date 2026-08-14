package relay

import "encoding/json"

// processEventPayload - converts an event payload to its target payload type
func processEventPayload(payload any, target any) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, target)
}
