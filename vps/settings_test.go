package vps

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/transip/gotransip/v6/internal/testutil"
)

func TestSettingsRepository_GetAll(t *testing.T) {
	const apiResponse = `{"settings":[{"name":"blockVpsMailPorts","dataType":"boolean","readOnly":false,"value":{"valueString":"","valueBoolean":true}},{"name":"tcpMonitoringAvailable","dataType":"boolean","readOnly":true,"value":{"valueString":"","valueBoolean":true}}]}`
	server := testutil.MockServer{T: t, ExpectedURL: "/vps/example-vps/settings", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := SettingRepository{Client: *client}

	settings, err := repo.GetAll("example-vps")

	assert.NoError(t, err)
	assert.Len(t, settings, 2)

	assert.ElementsMatch(t, settings, []Setting{
		{
			Name:     "blockVpsMailPorts",
			DataType: SettingDataTypeBoolean,
			ReadOnly: false,
			Value: SettingValue{
				ValueBoolean: true,
				ValueString:  "",
			},
		},
		{
			Name:     "tcpMonitoringAvailable",
			DataType: SettingDataTypeBoolean,
			ReadOnly: true,
			Value: SettingValue{
				ValueBoolean: true,
				ValueString:  "",
			},
		},
	})
}

func TestSettingsRepository_GetByName(t *testing.T) {
	const apiResponse = `{"setting":{"name":"blockVpsMailPorts","dataType":"boolean","readOnly":false,"value":{"valueString":"","valueBoolean":true}}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/vps/example-vps/settings/blockVpsMailPorts", ExpectedMethod: "GET", StatusCode: 200, Response: apiResponse}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := SettingRepository{Client: *client}

	setting, err := repo.GetByName("example-vps", SettingBlockVPSMailPorts)

	assert.NoError(t, err)

	assert.Equal(t, setting, Setting{
		Name:     "blockVpsMailPorts",
		DataType: SettingDataTypeBoolean,
		ReadOnly: false,
		Value: SettingValue{
			ValueBoolean: true,
			ValueString:  "",
		},
	})
}

func TestSettingsRepository_Update(t *testing.T) {
	const apiRequest = `{"setting":{"name":"blockVpsMailPorts","dataType":"boolean","readOnly":false,"value":{"valueBoolean":true,"valueString":""}}}`
	server := testutil.MockServer{T: t, ExpectedURL: "/vps/example-vps/settings/blockVpsMailPorts", ExpectedMethod: "PUT", StatusCode: 204, ExpectedRequest: apiRequest}
	client, tearDown := server.GetClient()
	defer tearDown()
	repo := SettingRepository{Client: *client}

	setting := Setting{
		Name:     "blockVpsMailPorts",
		DataType: SettingDataTypeBoolean,
		ReadOnly: false,
		Value: SettingValue{
			ValueBoolean: true,
		},
	}
	err := repo.Update("example-vps", setting)

	assert.NoError(t, err)
}
