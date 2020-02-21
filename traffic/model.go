package traffic

// Traffic struct for Traffic
type Traffic struct {
	Traffic TrafficInformation `json:"trafficInformation"`
}

// TrafficInformation struct for TrafficInformation
type TrafficInformation struct {
	// The end date in 'Y-m-d' format
	EndDate string `json:"endDate"`
	// The maximum amount of bytes that can be used in this period
	MaxInBytes float32 `json:"maxInBytes"`
	// The start date in 'Y-m-d' format
	StartDate string `json:"startDate"`
	// The usage in bytes for this period
	UsedInBytes float32 `json:"usedInBytes"`
	// The usage in bytes
	UsedTotalBytes float32 `json:"usedTotalBytes"`
}
