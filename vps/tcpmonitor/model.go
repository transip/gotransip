package tcpmonitor

// TcpMonitor struct for TcpMonitor
type TcpMonitor struct {
	// Allowed time outs (numbers 1-5)
	AllowedTimeouts uint8 `json:"allowedTimeouts"`
	// Contact that will be notified for this monitor
	Contacts []Contact `json:"contacts"`
	// The hours when the TCP monitoring is ignored (no notifications are sent out)
	IgnoreTimes []IgnoreTime `json:"ignoreTimes"`
	// Checking interval in minutes (numbers 1-6)
	Interval int64 `json:"interval"`
	// IP Address that is monitored
	IpAddress string `json:"ipAddress"`
	// Title of the monitor
	Label string `json:"label"`
	// Ports that are monitored
	Ports []uint16 `json:"ports"`
}

// Contact struct for Contact
type Contact struct {
	// Send emails to contact
	EnableEmail bool `json:"enableEmail"`
	// Send SMS text messages to contact
	EnableSMS bool `json:"enableSMS"`
	// Monitoring contact id
	Id int64 `json:"id"`
}

// IgnoreTime struct for IgnoreTime
type IgnoreTime struct {
	// Start from (24 hour format)
	TimeFrom string `json:"timeFrom"`
	// End at (24 hour format)
	TimeTo string `json:"timeTo"`
}

// TcpMonitors struct for TcpMonitors
type TcpMonitors struct {
	// TcpMonitor list
	TcpMonitor []TcpMonitor `json:"tcpMonitors"`
}

// MonitoringContacts struct for MonitoringContacts
type MonitoringContacts struct {
	// List of monitoring contacts
	Contacts []Contact `json:"contacts"`
}
