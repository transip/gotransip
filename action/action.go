package action

import "encoding/json"

// Action struct
type Action struct {
	UUID             string          `json:"uuid,omitempty"`
	Name             string          `json:"name"`
	ActionStartTime  string          `json:"actionStartTime"`
	Status           string          `json:"status"`
	Metadata         json.RawMessage `json:"metadata"`
	ParentActionUUID string          `json:"parentActionUuid"`
}

// actionWrapper struct contains a single Action in it,
// this is solely used for unmarshalling/marshalling
type actionWrapper struct {
	Action Action `json:"action"`
}

// actionsWrapper struct contains a list of Actions in it,
// this is solely used for unmarshalling/marshalling
type actionsWrapper struct {
	Actions []Action `json:"actions"`
}
