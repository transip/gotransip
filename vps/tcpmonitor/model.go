package tcpmonitor

import "net"

// TcpMonitor struct for TcpMonitor
type TcpMonitor struct {
	// Allowed time outs (numbers 1-5)
	AllowedTimeouts uint8 `json:"allowedTimeouts"`
	// TcpMonitorContact that will be notified for this monitor
	Contacts []TcpMonitorContact `json:"contacts"`
	// The hours when the TCP monitoring is ignored (no notifications are sent out)
	IgnoreTimes []IgnoreTime `json:"ignoreTimes"`
	// Checking interval in minutes (numbers 1-6)
	Interval int64 `json:"interval"`
	// IP Address that is monitored
	IPAddress net.IP `json:"ipAddress"`
	// Title of the monitor
	Label string `json:"label"`
	// Ports that are monitored
	Ports []uint16 `json:"ports"`
}

// TcpMonitorContact struct for TcpMonitorContact
type TcpMonitorContact struct {
	// Send emails to contact
	EnableEmail bool `json:"enableEmail"`
	// Send SMS text messages to contact
	EnableSMS bool `json:"enableSMS"`
	// Monitoring contact id
	Id uint64 `json:"id"`
}

// IgnoreTime struct for IgnoreTime
type IgnoreTime struct {
	// Start from (24 hour format)
	TimeFrom string `json:"timeFrom"`
	// End at (24 hour format)
	TimeTo string `json:"timeTo"`
}

// MonitoringContact struct for MonitoringContact
type MonitoringContact struct {
	// Id number of the contact
	Id uint64 `json:"id"`
	// Name of the contact
	Name string `json:"name"`
	// Telephone number of the contact
	Telephone string `json:"telephone"`
	// Email address of the contact
	Email string `json:"email"`
}
