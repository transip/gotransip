package webhosting

import (
	"fmt"

	"github.com/transip/gotransip/v5"
)

const (
	serviceName string = "WebhostingService"
)

// Package represents a Transip_WebhostingPackage as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_WebhostingPackage.html
type Package struct {
	Name         string  `xml:"name"`
	Description  string  `xml:"description"`
	Price        float64 `xml:"price"`
	RenewalPrice float64 `xml:"renewalPrice"`
}

// SpamCheckStrength defines the possible strengths for the spam checker
type SpamCheckStrength string

var (
	// SpamCheckStrengthAverage defines average spam checking
	SpamCheckStrengthAverage SpamCheckStrength = "AVERAGE"
	// SpamCheckStrengthOff disables spam checking
	SpamCheckStrengthOff SpamCheckStrength = "OFF"
	// SpamCheckStrengthLow defines low spam checking
	SpamCheckStrengthLow SpamCheckStrength = "LOW"
	// SpamCheckStrengthHigh defines high spam checking
	SpamCheckStrengthHigh SpamCheckStrength = "HIGH"
)

// MailBox represents a Transip_Mailbox object as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_Mailbox.html
type MailBox struct {
	Address              string            `xml:"address"`
	SpamCheckerStrength  SpamCheckStrength `xml:"spamCheckerStrength"`
	MaxDiskUsage         int               `xml:"maxDiskUsage"`
	HasVacationReply     bool              `xml:"hasVacationReply"`
	VacationReplySubject string            `xml:"vacationReplySubject"`
	VacationReplyMessage string            `xml:"vacationReplyMessage"`
}

// EncodeParams returns MailBox parameters ready to be used for constructing a signature
// the order of parameters added here has to match the order in the WSDL
// as described at http://api.transip.nl/wsdl/?service=WebhostingService
func (m MailBox) EncodeParams(prm gotransip.ParamsContainer, prefix string) {
	if len(prefix) == 0 {
		prefix = fmt.Sprintf("%d", prm.Len())
	}

	prm.Add(fmt.Sprintf("%s[address]", prefix), m.Address)
	prm.Add(fmt.Sprintf("%s[spamCheckerStrength]", prefix), string(m.SpamCheckerStrength))
	prm.Add(fmt.Sprintf("%s[maxDiskUsage]", prefix), fmt.Sprintf("%d", m.MaxDiskUsage))
	prm.Add(fmt.Sprintf("%s[hasVacationReply]", prefix), m.HasVacationReply)
	prm.Add(fmt.Sprintf("%s[vacationReplySubject]", prefix), m.VacationReplySubject)
	prm.Add(fmt.Sprintf("%s[vacationReplyMessage]", prefix), m.VacationReplyMessage)
}

// EncodeArgs returns MailBox XML body ready to be passed in the SOAP call
func (m MailBox) EncodeArgs(key string) string {
	return fmt.Sprintf(`<%s xsi:type="ns1:MailBox">
	<address xsi:type="xsd:string">%s</address>
	<spamCheckerStrength xsi:type="xsd:string">%s</spamCheckerStrength>
	<maxDiskUsage xsi:type="xsd:int">%d</maxDiskUsage>
	<hasVacationReply xsi:type="xsd:boolean">%t</hasVacationReply>
	<vacationReplySubject xsi:type="xsd:string">%s</vacationReplySubject>
	<vacationReplyMessage xsi:type="xsd:string">%s</vacationReplyMessage>
</%s>`, key, m.Address, m.SpamCheckerStrength, m.MaxDiskUsage, m.HasVacationReply,
		m.VacationReplySubject, m.VacationReplyMessage, key)
}

// Host represents a Transip_WebHost as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_WebHost.html
type Host struct {
	DomainName  string        `xml:"domainName"`
	CronJobs    []CronJob     `xml:"cronjobs>item"`
	MailBoxes   []MailBox     `xml:"mailBoxes>item"`
	Database    []Database    `xml:"dbs>item"`
	MailForward []MailForward `xml:"mailForwards>item"`
	SubDomains  []SubDomain   `xml:"subDomains>item"`
}

// SubDomain represents a Transip_SubDomain object
// as described at https://api.transip.nl/docs/transip.nl/class-Transip_SubDomain.html
type SubDomain struct {
	Name string `xml:"name"`
}

// EncodeParams returns SubDomain parameters ready to be used for constructing a signature
// the order of parameters added here has to match the order in the WSDL
// as described at http://api.transip.nl/wsdl/?service=WebhostingService
func (s SubDomain) EncodeParams(prm gotransip.ParamsContainer, prefix string) {
	if len(prefix) == 0 {
		prefix = fmt.Sprintf("%d", prm.Len())
	}

	prm.Add(fmt.Sprintf("%s[name]", prefix), s.Name)
}

// EncodeArgs returns SubDomain XML body ready to be passed in the SOAP call
func (s SubDomain) EncodeArgs(key string) string {
	return fmt.Sprintf(`<%s xsi:type="ns1:SubDomain">
	<name xsi:type="xsd:string">%s</name>
</%s>`, key, s.Name, key)
}

// Forward represents a Transip_Forward object as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_Forward.html
type Forward struct {
	DomainName        string `xml:"domainName"`
	ForwardTo         string `xml:"forwardTo"`
	ForwardMethod     string `xml:"forwardMethod"`
	FrameTitle        string `xml:"frameTitle"`
	FrameIcon         string `xml:"frameIcon"`
	ForwardEverything bool   `xml:"forwardEverything"`
	ForwardSubdomains bool   `xml:"forwardSubdomains"`
	ForwardEmailTo    string `xml:"forwardEmailTo"`
}

// MailForward represents a Transip_MailForward object as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_MailForward.html
type MailForward struct {
	Name          string `xml:"name"`
	TargetAddress string `xml:"targetAddress"`
}

// EncodeParams returns MailForward parameters ready to be used for constructing a signature
// the order of parameters added here has to match the order in the WSDL
// as described at http://api.transip.nl/wsdl/?service=WebhostingService
func (m MailForward) EncodeParams(prm gotransip.ParamsContainer, prefix string) {
	if len(prefix) == 0 {
		prefix = fmt.Sprintf("%d", prm.Len())
	}

	prm.Add(fmt.Sprintf("%s[name]", prefix), m.Name)
	prm.Add(fmt.Sprintf("%s[targetAddress]", prefix), m.TargetAddress)
}

// EncodeArgs returns MailForward XML body ready to be passed in the SOAP call
func (m MailForward) EncodeArgs(key string) string {
	return fmt.Sprintf(`<%s xsi:type="ns1:MailForward">
	<name xsi:type="xsd:string">%s</name>
	<targetAddress xsi:type="xsd:string">%s</targetAddress>
</%s>`, key, m.Name, m.TargetAddress, key)
}

// Database represents a Transip_Db object as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_Db.html
type Database struct {
	Name         string `xml:"name"`
	Username     string `xml:"username"`
	MaxDiskUsage int    `xml:"maxDiskUsage"`
}

// EncodeParams returns Database parameters ready to be used for constructing a signature
// the order of parameters added here has to match the order in the WSDL
// as described at http://api.transip.nl/wsdl/?service=WebhostingService
func (db Database) EncodeParams(prm gotransip.ParamsContainer, prefix string) {
	if len(prefix) == 0 {
		prefix = fmt.Sprintf("%d", prm.Len())
	}

	prm.Add(fmt.Sprintf("%s[name]", prefix), db.Name)
	prm.Add(fmt.Sprintf("%s[username]", prefix), db.Username)
	prm.Add(fmt.Sprintf("%s[maxDiskUsage]", prefix), fmt.Sprintf("%d", db.MaxDiskUsage))
}

// EncodeArgs returns Database XML body ready to be passed in the SOAP call
func (db Database) EncodeArgs(key string) string {
	return fmt.Sprintf(`<%s xsi:type="ns1:Db">
	<name xsi:type="xsd:string">%s</name>
	<username xsi:type="xsd:string">%s</username>
	<maxDiskUsage xsi:type="xsd:int">%d</maxDiskUsage>
</%s>`, key, db.Name, db.Username, db.MaxDiskUsage, key)
}

// CronJob represents a Transip_Cronjob object as described at
// https://api.transip.nl/docs/transip.nl/class-Transip_Cronjob.html
type CronJob struct {
	Name           string `xml:"name"`
	URL            string `xml:"url"`
	Email          string `xml:"email"`
	MinuteTrigger  string `xml:"minuteTrigger"`
	HourTrigger    string `xml:"hourTrigger"`
	DayTrigger     string `xml:"dayTrigger"`
	MonthTrigger   string `xml:"monthTrigger"`
	WeekdayTrigger string `xml:"weekdayTrigger"`
}

// EncodeParams returns CronJob parameters ready to be used for constructing a signature
// the order of parameters added here has to match the order in the WSDL
// as described at http://api.transip.nl/wsdl/?service=WebhostingService
func (c CronJob) EncodeParams(prm gotransip.ParamsContainer, prefix string) {
	if len(prefix) == 0 {
		prefix = fmt.Sprintf("%d", prm.Len())
	}

	prm.Add(fmt.Sprintf("%s[name]", prefix), c.Name)
	prm.Add(fmt.Sprintf("%s[url]", prefix), c.URL)
	prm.Add(fmt.Sprintf("%s[email]", prefix), c.Email)
	prm.Add(fmt.Sprintf("%s[minuteTrigger]", prefix), c.MinuteTrigger)
	prm.Add(fmt.Sprintf("%s[hourTrigger]", prefix), c.HourTrigger)
	prm.Add(fmt.Sprintf("%s[dayTrigger]", prefix), c.DayTrigger)
	prm.Add(fmt.Sprintf("%s[monthTrigger]", prefix), c.MonthTrigger)
	prm.Add(fmt.Sprintf("%s[weekdayTrigger]", prefix), c.WeekdayTrigger)
}

// EncodeArgs returns CronJob XML body ready to be passed in the SOAP call
func (c CronJob) EncodeArgs(key string) string {
	return fmt.Sprintf(`<%s xsi:type="ns1:Cronjob">
	<name xsi:type="xsd:string">%s</name>
	<url xsi:type="xsd:string">%s</url>
	<email xsi:type="xsd:string">%s</email>
	<minuteTrigger xsi:type="xsd:string">%s</minuteTrigger>
	<hourTrigger xsi:type="xsd:string">%s</hourTrigger>
	<dayTrigger xsi:type="xsd:string">%s</dayTrigger>
	<monthTrigger xsi:type="xsd:string">%s</monthTrigger>
	<weekdayTrigger xsi:type="xsd:string">%s</weekdayTrigger>
</%s>`, key, c.Name, c.URL, c.Email, c.MinuteTrigger, c.HourTrigger,
		c.DayTrigger, c.MonthTrigger, c.WeekdayTrigger, key)
}
