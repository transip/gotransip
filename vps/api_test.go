package vps

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip"
)

func TestGetActiveAddonsForVps(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getactiveaddonsforvps.xml")
	require.NoError(t, err)

	lst, err := GetActiveAddonsForVps(c, "example-vps")
	require.NoError(t, err)
	require.Equal(t, 3, len(lst))
	assert.IsType(t, []Product{}, lst)
	assert.Equal(t, "100 GB extra SSD", lst[0].Description)
	assert.Equal(t, "vpsAddon-100-gb-extra-harddisk", lst[0].Name)
	assert.Equal(t, 2.5, lst[0].Price)
	assert.Equal(t, 7.5, lst[0].RenewalPrice)
	assert.Equal(t, "1 extra CPU core", lst[1].Description)
}

func TestGetAllIps(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getallips.xml")
	assert.NoError(t, err)

	lst, err := GetAllIps(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []net.IP{}, lst)
	assert.Equal(t, net.ParseIP("1.2.3.4"), lst[0])
	assert.Equal(t, net.ParseIP("2a01:7c8::1"), lst[1])
}

func TestGetAllPrivateNetworks(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getallprivatenetworks.xml")
	assert.NoError(t, err)

	lst, err := GetAllPrivateNetworks(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []PrivateNetwork{}, lst)
	assert.Equal(t, "example-privatenetwork", lst[0].Name)
	assert.Equal(t, "example-privatenetwork2", lst[1].Name)
}

func TestSetCustomerLock(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/setcustomerlock.xml")
	require.NoError(t, err)

	err = SetCustomerLock(c, "test-vps", true)
	require.NoError(t, err)
}

func TestGetAvailableAvailabilityZones(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getavailableavailabilityzones.xml")
	require.NoError(t, err)

	lst, err := GetAvailableAvailabilityZones(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []AvailabilityZone{}, lst)
	assert.Equal(t, "ams0", lst[0].Name)
	assert.Equal(t, "nl", lst[0].Country)
	assert.Equal(t, true, lst[0].IsDefault)
	assert.Equal(t, "rtm0", lst[1].Name)
}

func TestGetAvailableAddons(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getavailableaddons.xml")
	require.NoError(t, err)

	lst, err := GetAvailableAddons(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Product{}, lst)
	assert.Equal(t, "vpsAddon-100-gb-extra-harddisk", lst[0].Name)
	assert.Equal(t, "100 GB extra SSD", lst[0].Description)
	assert.Equal(t, 2.5, lst[0].Price)
	assert.Equal(t, 7.5, lst[0].RenewalPrice)
	assert.Equal(t, "vpsAddon-1-extra-ip-address", lst[1].Name)
}

func TestGetAvailableAddonsForVps(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getavailableaddonsforvps.xml")
	require.NoError(t, err)

	lst, err := GetAvailableAddonsForVps(c, "example-vps")
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Product{}, lst)
	assert.Equal(t, "vpsAddon-1-extra-cpu-core", lst[0].Name)
	assert.Equal(t, "1 extra CPU core", lst[0].Description)
	assert.Equal(t, 2.5, lst[0].Price)
	assert.Equal(t, 5.0, lst[0].RenewalPrice)
	assert.Equal(t, "vpsAddon-extra-memory", lst[1].Name)
}

func TestGetAvailableProducts(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getavailableproducts.xml")
	require.NoError(t, err)

	lst, err := GetAvailableProducts(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Product{}, lst)
	assert.Equal(t, "vps-bladevps-x1", lst[0].Name)
	assert.Equal(t, "BladeVPS PureSSD X1", lst[0].Description)
	assert.Equal(t, 5.0, lst[0].Price)
	assert.Equal(t, 10.0, lst[0].RenewalPrice)
	assert.Equal(t, "vps-bladevps-x4", lst[1].Name)
}

func TestGetAvailableUpgrades(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getavailableupgrades.xml")
	require.NoError(t, err)

	lst, err := GetAvailableUpgrades(c, "example-vps")
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Product{}, lst)
	assert.Equal(t, "vps-bladevps-pro-x32", lst[0].Name)
	assert.Equal(t, "BladeVPS PureSSD PRO X32", lst[0].Description)
	assert.Equal(t, 199.0, lst[0].Price)
	assert.Equal(t, 299.0, lst[0].RenewalPrice)
	assert.Equal(t, "vps-bladevps-pro-x24", lst[1].Name)
}

func TestGetCancellableAddonsForVps(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getcancellableaddonsforvps.xml")
	require.NoError(t, err)

	lst, err := GetCancellableAddonsForVps(c, "example-vps")
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Product{}, lst)
	assert.Equal(t, "vpsAddon-100-gb-extra-harddisk", lst[0].Name)
	assert.Equal(t, "100 GB extra SSD", lst[0].Description)
	assert.Equal(t, 2.5, lst[0].Price)
	assert.Equal(t, 7.5, lst[0].RenewalPrice)
	assert.Equal(t, "vpsAddon-1-extra-ip-address", lst[1].Name)
}

func TestOrderVps(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/ordervps.xml")
	require.NoError(t, err)

	err = OrderVps(c, "vps-bladevps-x1", nil, "centos65", "test")
	require.NoError(t, err)
}

func TestOrderVpsInAvailabilityZone(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/ordervpsinavailabilityzone.xml")
	require.NoError(t, err)

	err = OrderVpsInAvailabilityZone(c, "vps-bladevps-x1", nil, "centos65", "test", "ams0")
	require.NoError(t, err)
}

func TestCloneVps(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/clonevps.xml")
	require.NoError(t, err)

	err = CloneVps(c, "test-vps")
	require.NoError(t, err)
}

func TestCloneVpsToAvailabilityZone(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/clonevpstoavailabilityzone.xml")
	require.NoError(t, err)

	err = CloneVpsToAvailabilityZone(c, "test-vps", "ams0")
	require.NoError(t, err)
}

func TestOrderAddon(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/orderaddon.xml")
	require.NoError(t, err)

	err = OrderAddon(c, "test-vps", []string{"vpsAddon-extra-memory"})
	require.NoError(t, err)
}

func TestOrderPrivateNetwork(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/orderprivatenetwork.xml")
	require.NoError(t, err)

	err = OrderPrivateNetwork(c)
	require.NoError(t, err)
}

func TestUpgradeVps(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/upgradevps.xml")
	require.NoError(t, err)

	err = UpgradeVps(c, "test-vps", "vps-bladevps-x4")
	require.NoError(t, err)
}

func TestCancelVps(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/cancelvps.xml")
	require.NoError(t, err)

	err = CancelVps(c, "test-vps", gotransip.CancellationTimeImmediately)
	require.NoError(t, err)
}

func TestCancelAddon(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/canceladdon.xml")
	require.NoError(t, err)

	err = CancelAddon(c, "test-vps", "vpsAddon-extra-memory")
	require.NoError(t, err)
}

func TestCancelPrivateNetwork(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/cancelprivatenetwork.xml")
	require.NoError(t, err)

	err = CancelPrivateNetwork(c, "test-privatenetwork", gotransip.CancellationTimeEnd)
	require.NoError(t, err)
}

func TestGetIpsForVps(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getipsforvps.xml")
	require.NoError(t, err)

	lst, err := GetIpsForVps(c, "example-vps")
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []net.IP{}, lst)
	assert.Equal(t, net.ParseIP("1.2.3.4"), lst[0])
	assert.Equal(t, net.ParseIP("2a01:7c8::1"), lst[1])
}

func TestGetOperatingSystems(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getoperatingsystems.xml")
	require.NoError(t, err)

	lst, err := GetOperatingSystems(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []OperatingSystem{}, lst)
	assert.Equal(t, "gentoo12", lst[0].Name)
	assert.Equal(t, "Gentoo", lst[0].Description)
	assert.Equal(t, true, lst[0].IsPreinstallableImage)
	assert.Equal(t, "centos65", lst[1].Name)
}

func TestAddVpsToPrivateNetwork(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/addvpstoprivatenetwork.xml")
	require.NoError(t, err)

	err = AddVpsToPrivateNetwork(c, "test-vps", "test-privatenetwork")
	require.NoError(t, err)
}

func TestRemoveVpsFromPrivateNetwork(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/removevpsfromprivatenetwork.xml")
	require.NoError(t, err)

	err = RemoveVpsFromPrivateNetwork(c, "test-vps", "test-privatenetwork")
	require.NoError(t, err)
}

func TestGetTrafficInformationForVps(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/gettrafficinformationforvps.xml")
	require.NoError(t, err)

	ti, err := GetTrafficInformationForVps(c, "example-vps")
	require.NoError(t, err)
	assert.IsType(t, TrafficInformation{}, ti)
	x, _ := time.Parse("2006-01-02", "2018-09-01")
	assert.Equal(t, x, ti.From)
	y, _ := time.Parse("2006-01-02", "2018-10-01")
	assert.Equal(t, y, ti.End)
	assert.Equal(t, int64(2483004356), ti.Used)
	assert.Equal(t, int64(4380854877), ti.Total)
	assert.Equal(t, int64(1073741824000), ti.Max)
}

func TestGetPooledTrafficInformation(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getpooledtrafficinformation.xml")
	require.NoError(t, err)

	ti, err := GetPooledTrafficInformation(c)
	require.NoError(t, err)
	assert.IsType(t, TrafficInformation{}, ti)
	x, _ := time.Parse("2006-01-02", "2018-08-01")
	assert.Equal(t, x, ti.From)
	y, _ := time.Parse("2006-01-02", "2018-09-01")
	assert.Equal(t, y, ti.End)
	assert.Equal(t, int64(128356924407), ti.Used)
	assert.Equal(t, int64(200639805776), ti.Total)
	assert.Equal(t, int64(16106127360000), ti.Max)
}

func TestStart(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/start.xml")
	require.NoError(t, err)

	err = Start(c, "test-vps")
	require.NoError(t, err)
}

func TestStop(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/stop.xml")
	require.NoError(t, err)

	err = Stop(c, "test-vps")
	require.NoError(t, err)
}

func TestReset(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/reset.xml")
	require.NoError(t, err)

	err = Reset(c, "test-vps")
	require.NoError(t, err)
}

func TestCreateSnapshot(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/createsnapshot.xml")
	require.NoError(t, err)

	err = CreateSnapshot(c, "test-vps", "test-snapshot")
	require.NoError(t, err)
}

func TestRevertSnapshot(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/revertsnapshot.xml")
	require.NoError(t, err)

	err = RevertSnapshot(c, "test-vps", "test-snapshot")
	require.NoError(t, err)
}

func TestRevertSnapshotToOtherVps(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/revertsnapshottoothervps.xml")
	require.NoError(t, err)

	err = RevertSnapshotToOtherVps(c, "test-vps", "test-snapshot", "test-vps2")
	require.NoError(t, err)
}

func TestRemoveSnapshot(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/removesnapshot.xml")
	require.NoError(t, err)

	err = RemoveSnapshot(c, "test-vps", "test-snapshot")
	require.NoError(t, err)
}

func TestRevertVpsBackup(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/revertvpsbackup.xml")
	require.NoError(t, err)

	err = RevertVpsBackup(c, "test-vps", 1234)
	require.NoError(t, err)
}

func TestGetPrivateNetworksByVps(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getprivatenetworksbyvps.xml")
	require.NoError(t, err)

	lst, err := GetPrivateNetworksByVps(c, "example-vps")
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []PrivateNetwork{}, lst)
	assert.Equal(t, "example-privatenetwork", lst[0].Name)
	assert.Equal(t, "example-privatenetwork2", lst[1].Name)
}

func TestGetVps(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getvps.xml")
	require.NoError(t, err)

	v, err := GetVps(c, "example-vps")
	require.NoError(t, err)
	assert.IsType(t, Vps{}, v)
	assert.Equal(t, "rtm0", v.AvailabilityZone)
	assert.Equal(t, "my vps", v.Description)
	assert.Equal(t, int64(52428800), v.DiskSize)
	assert.Equal(t, net.ParseIP("1.2.3.4"), v.IPv4Address)
	assert.Equal(t, net.ParseIP("2a01:7c8::1"), v.IPv6Address)
	assert.Equal(t, true, v.IsBlocked)
	assert.Equal(t, true, v.IsCustomerLocked)
	assert.Equal(t, "52:54:00:01:02:03", v.MACAddress)
	assert.Equal(t, int64(1048576), v.MemorySize)
	assert.Equal(t, "example-vps", v.Name)
	assert.Equal(t, "ubuntu-18.04", v.OperatingSystem)
	assert.Equal(t, int64(1), v.Processors)
	assert.Equal(t, StatusRunning, v.Status)
}

func TestGetVpsSnapshotsByVps(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getvpssnapshotsbyvps.xml")
	require.NoError(t, err)

	lst, err := GetSnapshotsByVps(c, "example-vps")
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Snapshot{}, lst)
	assert.Equal(t, "1501750169", lst[0].Name)
	x, _ := time.Parse("2006-01-02 15:04:05", "2017-08-03 10:49:29")
	assert.Equal(t, x, lst[0].Created.Time)
	assert.Equal(t, "clean install ubuntu 17.04", lst[0].Description)
	assert.Equal(t, "ams0", lst[0].AvailabilityZone)
	assert.Equal(t, "1501860169", lst[1].Name)
}

func TestGetVpses(t *testing.T) {
	var err error

	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getvpses.xml")
	require.NoError(t, err)

	lst, err := GetVpses(c)
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Vps{}, lst)
	assert.Equal(t, "rtm0", lst[0].AvailabilityZone)
	assert.Equal(t, "my vps", lst[0].Description)
	assert.Equal(t, int64(157286400), lst[0].DiskSize)
	assert.Equal(t, net.ParseIP("1.2.3.4"), lst[0].IPv4Address)
	assert.Equal(t, net.ParseIP("2a01:7c8::1"), lst[0].IPv6Address)
	assert.Equal(t, true, lst[0].IsBlocked)
	assert.Equal(t, true, lst[0].IsCustomerLocked)
	assert.Equal(t, "52:54:00:01:02:03", lst[0].MACAddress)
	assert.Equal(t, int64(1048576), lst[0].MemorySize)
	assert.Equal(t, "example-vps", lst[0].Name)
	assert.Equal(t, "ubuntu-18.04", lst[0].OperatingSystem)
	assert.Equal(t, int64(2), lst[0].Processors)
	assert.Equal(t, StatusRunning, lst[0].Status)
	assert.Equal(t, "example-vps2", lst[1].Name)
}

func TestGetVpsBackupsByVps(t *testing.T) {
	var err error
	c := gotransip.FakeSOAPClient{}
	err = c.FixtureFromFile("testdata/getvpsbackupsbyvps.xml")
	require.NoError(t, err)

	lst, err := GetVpsBackupsByVps(c, "example-vps")
	require.NoError(t, err)
	require.Equal(t, 2, len(lst))
	assert.IsType(t, []Backup{}, lst)
	assert.Equal(t, int64(6996570), lst[0].ID)
	x, _ := time.Parse("2006-01-02 15:04:05", "2018-09-11 01:32:51")
	assert.Equal(t, x, lst[0].Created.Time)
	assert.Equal(t, int64(157286400), lst[0].DiskSize)
	assert.Equal(t, "Ubuntu 18.04 LTS", lst[0].OperatingSystem)
	assert.Equal(t, "ams0", lst[0].AvailabilityZone)
	assert.Equal(t, int64(7039319), lst[1].ID)
}

func TestInstallOperatingSystem(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/installoperatingsystem.xml")
	require.NoError(t, err)

	err = InstallOperatingSystem(c, "test-vps", "centos65", "test")
	require.NoError(t, err)
}

func TestInstallOperatingSystemUnattended(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/installoperatingsystemunattended.xml")
	require.NoError(t, err)

	err = InstallOperatingSystemUnattended(c, "test-vps", "centos65", "cHJlc2VlZAo=")
	require.NoError(t, err)
}

func TestAddIpv6ToVps(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/addipv6tovps.xml")
	require.NoError(t, err)

	err = AddIpv6ToVps(c, "test-vps", net.ParseIP("fe80::f00/64"))
	require.NoError(t, err)
}

func TestUpdatePtrRecord(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/updateptrrecord.xml")
	require.NoError(t, err)

	err = UpdatePtrRecord(c, net.IP{127, 0, 0, 1}, "ptr.not.set")
	require.NoError(t, err)
}

func TestHandoverVps(t *testing.T) {
	c := gotransip.FakeSOAPClient{}
	err := c.FixtureFromFile("testdata/handovervps.xml")
	require.NoError(t, err)

	err = HandoverVps(c, "test-vps", "customer")
	require.NoError(t, err)
}
