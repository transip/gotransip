package traffic

import (
	"github.com/transip/gotransip/v6/rest"
)

// this struct will be used to unmarshal it only contains a TrafficInformation struct in it
type trafficWrapper struct {
	TrafficInformation TrafficInformation `json:"trafficInformation"`
}

// TrafficInformation struct for TrafficInformation
type TrafficInformation struct {
	// The end date in 'Y-m-d' format
	EndDate rest.Date `json:"endDate"`
	// The maximum amount of bytes that can be used in this period
	MaxInBytes int64 `json:"maxInBytes"`
	// The start date in 'Y-m-d' format
	StartDate rest.Date `json:"startDate"`
	// The usage in bytes for this period
	UsedInBytes int64 `json:"usedInBytes"`
	// The usage in bytes
	UsedTotalBytes int64 `json:"usedTotalBytes"`
}
