package webhosting

import "github.com/transip/gotransip"

// This file holds all WebhostingService methods directly ported from TransIP API

// GetWebhostingDomainNames returns all domain names that have a webhosting
// package attached to them.
func GetWebhostingDomainNames(c gotransip.Client) ([]string, error) {
	var v struct {
		V []string `xml:"item"`
	}
	err := c.Call(gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getWebhostingDomainNames",
	}, &v)

	return v.V, err
}

// GetAvailablePackages returns all available webhosting packages
func GetAvailablePackages(c gotransip.Client) ([]Package, error) {
	var v struct {
		V []Package `xml:"item"`
	}

	err := c.Call(gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getAvailablePackages",
	}, &v)

	return v.V, err
}

// GetInfo retuns information about existing webhosting on a domain.
func GetInfo(c gotransip.Client, domainName string) (Host, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getInfo",
	}
	sr.AddArgument("domainName", domainName)

	var v Host
	err := c.Call(sr, &v)
	return v, err
}

// Order orders a webhosting package for a domain name
func Order(c gotransip.Client, domainName, webhostingPackage string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "order",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("webhostingPackage", webhostingPackage)

	return c.Call(sr, nil)
}

// GetAvailableUpgrades get available upgrades packages for a domain name with
// webhosting. Only those packages will be returned to which the given domain
// name can be upgraded to.
func GetAvailableUpgrades(c gotransip.Client, domainName string) ([]Package, error) {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "getAvailableUpgrades",
	}
	sr.AddArgument("domainName", domainName)

	var v struct {
		V []Package `xml:"item"`
	}

	err := c.Call(sr, &v)
	return v.V, err
}

// Upgrade upgrades the webhosting package of a domain name to a new webhosting
// package
func Upgrade(c gotransip.Client, domainName, newWebhostingPackage string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "upgrade",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("newWebhostingPackage", newWebhostingPackage)

	return c.Call(sr, nil)
}

// Cancel cancels webhosting for a domain
func Cancel(c gotransip.Client, domainName, endTime gotransip.CancellationTime) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "cancel",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("endTime", string(endTime))

	return c.Call(sr, nil)
}

// SetFtpPassword sets a new FTP password for a webhosting package
func SetFtpPassword(c gotransip.Client, domainName, newPassword string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setFtpPassword",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("newPassword", newPassword)

	return c.Call(sr, nil)
}

// CreateCronjob creates a cronjob
func CreateCronjob(c gotransip.Client, domainName string, cronjob CronJob) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "createCronjob",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("cronjob", cronjob)

	return c.Call(sr, nil)
}

// DeleteCronjob deletes a cronjob from a webhosting package.
func DeleteCronjob(c gotransip.Client, domainName, cronjob CronJob) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "deleteCronjob",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("cronjob", cronjob)

	return c.Call(sr, nil)
}

// CreateMailBox creates a MailBox for a webhosting package
func CreateMailBox(c gotransip.Client, domainName string, mailBox MailBox) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "createMailBox",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("mailBox", mailBox)

	return c.Call(sr, nil)
}

// ModifyMailBox updates MailBox settings
func ModifyMailBox(c gotransip.Client, domainName string, mailBox MailBox) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "modifyMailBox",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("mailBox", mailBox)

	return c.Call(sr, nil)
}

// SetMailBoxPassword sets a new password for a MailBox
func SetMailBoxPassword(c gotransip.Client, domainName string, mailBox MailBox, newPassword string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setMailBoxPassword",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("mailBox", mailBox)
	sr.AddArgument("newPassword", newPassword)

	return c.Call(sr, nil)
}

// DeleteMailBox deletes a MailBox from a webhosting package
func DeleteMailBox(c gotransip.Client, domainName string, mailBox MailBox) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "deleteMailBox",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("mailBox", mailBox)

	return c.Call(sr, nil)
}

// CreateMailForward creates a MailForward for a webhosting package
func CreateMailForward(c gotransip.Client, domainName string, mailForward MailForward) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "createMailForward",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("mailForward", mailForward)

	return c.Call(sr, nil)
}

// ModifyMailForward changes an active MailForward object
func ModifyMailForward(c gotransip.Client, domainName string, mailForward MailForward) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "modifyMailForward",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("mailForward", mailForward)

	return c.Call(sr, nil)
}

// DeleteMailForward deletes an active MailForward object
func DeleteMailForward(c gotransip.Client, domainName string, mailForward MailForward) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "deleteMailForward",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("mailForward", mailForward)

	return c.Call(sr, nil)
}

// CreateDatabase creates a new database
func CreateDatabase(c gotransip.Client, domainName string, db Database) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "createDatabase",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("db", db)

	return c.Call(sr, nil)
}

// ModifyDatabase changes a Db object
func ModifyDatabase(c gotransip.Client, domainName string, db Database) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "modifyDatabase",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("db", db)

	return c.Call(sr, nil)
}

// SetDatabasePassword sets a database password for a Db
func SetDatabasePassword(c gotransip.Client, domainName string, db Database, newPassword string) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "setDatabasePassword",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("db", db)
	sr.AddArgument("newPassword", newPassword)

	return c.Call(sr, nil)
}

// DeleteDatabase deletes a Db object
func DeleteDatabase(c gotransip.Client, domainName string, db Database) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "deleteDatabase",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("db", db)

	return c.Call(sr, nil)
}

// CreateSubdomain creates a SubDomain
func CreateSubdomain(c gotransip.Client, domainName string, subDomain SubDomain) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "createSubdomain",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("subDomain", subDomain)

	return c.Call(sr, nil)
}

// DeleteSubdomain deletes a SubDomain
func DeleteSubdomain(c gotransip.Client, domainName string, subDomain SubDomain) error {
	sr := gotransip.SoapRequest{
		Service: serviceName,
		Method:  "deleteSubdomain",
	}
	sr.AddArgument("domainName", domainName)
	sr.AddArgument("subDomain", subDomain)

	return c.Call(sr, nil)
}
