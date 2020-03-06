package traffic

import (
	"github.com/transip/gotransip/v6/rest/response"
)

// Traffic struct for Traffic
type Traffic struct {
	Traffic TrafficInformation `json:"trafficInformation"`
}

// this struct will be used to unmarshal it only contains a TrafficInformation struct in it
type trafficWrapper struct {
	TrafficInformation TrafficInformation `json:"trafficInformation"`
}

// TrafficInformation struct for TrafficInformation
type TrafficInformation struct {
	// The end date in 'Y-m-d' format
	EndDate response.Date `json:"endDate"`
	// The maximum amount of bytes that can be used in this period
	MaxInBytes uint64 `json:"maxInBytes"`
	// The start date in 'Y-m-d' format
	StartDate response.Date `json:"startDate"`
	// The usage in bytes for this period
	UsedInBytes uint64 `json:"usedInBytes"`
	// The usage in bytes
	UsedTotalBytes uint64 `json:"usedTotalBytes"`
}
