package kubernetes

import (
	"github.com/transip/gotransip/v6/rest"
)

// Release is a Kubernetes release version
type Release struct {
	Version             string    `json:"version"`
	ReleaseDate         rest.Date `json:"releaseDate"`
	MaintenanceModeDate rest.Date `json:"maintenanceModeDate"`
	EndOfLifeDate       rest.Date `json:"endOfLifeDate"`
}

type releaseWrapper struct {
	Release Release `json:"release"`
}

type releasesWrapper struct {
	Releases []Release `json:"releases"`
}
