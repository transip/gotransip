package kubernetes

type eventWrapper struct {
	Event Event `json:"event"`
}

type eventsWrapper struct {
	Events []Event `json:"events"`
}

// Event A kubernetes cluster event
type Event struct {
	Name               string `json:"name"`
	Namespace          string `json:"namespace"`
	Type               string `json:"type"`
	Message            string `json:"message"`
	Reason             string `json:"reason"`
	Count              int    `json:"count"`
	CreationTimestamp  int    `json:"creationTimestamp"`
	FirstTimestamp     int    `json:"firstTimestamp"`
	LastTimestamp      int    `json:"lastTimestamp"`
	InvolvedObjectKind string `json:"involvedObjectKind"`
	InvolvedObjectName string `json:"involvedObjectName"`
	SourceComponent    string `json:"sourceComponent"`
}
