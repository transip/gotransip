package webhosting

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/transip/gotransip/v5"
)

func TestMailBoxEncoding(t *testing.T) {
	mailbox := MailBox{
		Address:              "info@example.org",
		SpamCheckerStrength:  SpamCheckStrengthAverage,
		MaxDiskUsage:         100,
		HasVacationReply:     true,
		VacationReplySubject: "out of office",
		VacationReplyMessage: "I'm out of office, ktnxbye",
	}

	fixtArgs := `<mailbox xsi:type="ns1:MailBox">
	<address xsi:type="xsd:string">info@example.org</address>
	<spamCheckerStrength xsi:type="xsd:string">AVERAGE</spamCheckerStrength>
	<maxDiskUsage xsi:type="xsd:int">100</maxDiskUsage>
	<hasVacationReply xsi:type="xsd:boolean">true</hasVacationReply>
	<vacationReplySubject xsi:type="xsd:string">out of office</vacationReplySubject>
	<vacationReplyMessage xsi:type="xsd:string">I'm out of office, ktnxbye</vacationReplyMessage>
</mailbox>`
	assert.Equal(t, fixtArgs, mailbox.EncodeArgs("mailbox"))

	prm := gotransip.TestParamsContainer{}
	mailbox.EncodeParams(&prm, "")
	assert.Equal(t, "00[address]=info@example.org&280[spamCheckerStrength]=AVERAGE&610[maxDiskUsage]=100&830[hasVacationReply]=1&1070[vacationReplySubject]=out of office&1480[vacationReplyMessage]=I'm out of office, ktnxbye", prm.Prm)
}

func TestMailForwardEncoding(t *testing.T) {
	fwd := MailForward{
		Name:          "foobar@example.org",
		TargetAddress: "info@example.org",
	}

	fixtArgs := `<mailForward xsi:type="ns1:MailForward">
	<name xsi:type="xsd:string">foobar@example.org</name>
	<targetAddress xsi:type="xsd:string">info@example.org</targetAddress>
</mailForward>`
	assert.Equal(t, fixtArgs, fwd.EncodeArgs("mailForward"))

	prm := gotransip.TestParamsContainer{}
	fwd.EncodeParams(&prm, "")
	assert.Equal(t, "00[name]=foobar@example.org&270[targetAddress]=info@example.org", prm.Prm)
}

func TestDatabaseEncoding(t *testing.T) {
	db := Database{
		Name:         "test",
		Username:     "foobar",
		MaxDiskUsage: 100,
	}

	fixtArgs := `<database xsi:type="ns1:Db">
	<name xsi:type="xsd:string">test</name>
	<username xsi:type="xsd:string">foobar</username>
	<maxDiskUsage xsi:type="xsd:int">100</maxDiskUsage>
</database>`
	assert.Equal(t, fixtArgs, db.EncodeArgs("database"))

	prm := gotransip.TestParamsContainer{}
	db.EncodeParams(&prm, "")
	assert.Equal(t, "00[name]=test&130[username]=foobar&340[maxDiskUsage]=100", prm.Prm)
}

func TestSubDoimainEncoding(t *testing.T) {
	sd := SubDomain{
		Name: "demo.example.org",
	}

	fixtArgs := `<subdomain xsi:type="ns1:SubDomain">
	<name xsi:type="xsd:string">demo.example.org</name>
</subdomain>`
	assert.Equal(t, fixtArgs, sd.EncodeArgs("subdomain"))

	prm := gotransip.TestParamsContainer{}
	sd.EncodeParams(&prm, "")
	assert.Equal(t, "00[name]=demo.example.org", prm.Prm)
}

func TestCronjobEncoding(t *testing.T) {
	cron := CronJob{
		Name:           "test",
		URL:            "http://example.org/?foobar",
		Email:          "info@example.org",
		MinuteTrigger:  "1",
		HourTrigger:    "2",
		DayTrigger:     "3",
		MonthTrigger:   "4",
		WeekdayTrigger: "5",
	}

	fixtArgs := `<cronjob xsi:type="ns1:Cronjob">
	<name xsi:type="xsd:string">test</name>
	<url xsi:type="xsd:string">http://example.org/?foobar</url>
	<email xsi:type="xsd:string">info@example.org</email>
	<minuteTrigger xsi:type="xsd:string">1</minuteTrigger>
	<hourTrigger xsi:type="xsd:string">2</hourTrigger>
	<dayTrigger xsi:type="xsd:string">3</dayTrigger>
	<monthTrigger xsi:type="xsd:string">4</monthTrigger>
	<weekdayTrigger xsi:type="xsd:string">5</weekdayTrigger>
</cronjob>`
	assert.Equal(t, fixtArgs, cron.EncodeArgs("cronjob"))

	prm := gotransip.TestParamsContainer{}
	cron.EncodeParams(&prm, "")
	assert.Equal(t, "00[name]=test&130[url]=http://example.org/?foobar&490[email]=info@example.org&770[minuteTrigger]=1&980[hourTrigger]=2&1170[dayTrigger]=3&1360[monthTrigger]=4&1570[weekdayTrigger]=5", prm.Prm)
}
