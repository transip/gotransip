package vps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSettingsRepository_GetAll(t *testing.T) {
	const apiResponse = `{"settings":[{"name":"blockVpsMailPorts","dataType":"boolean","readOnly":false,"value":{"valueString":"","valueBoolean":true}},{"name":"tcpMonitoringAvailable","dataType":"boolean","readOnly":true,"value":{"valueString":"","valueBoolean":true}}]}` //nolint
	server := mockServer{t: t, expectedURL: "/vps/example-vps/settings", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := SettingsRepository{Client: *client}

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
	const apiResponse = `{"setting":{"name":"blockVpsMailPorts","dataType":"boolean","readOnly":false,"value":{"valueString":"","valueBoolean":true}}}` //nolint
	server := mockServer{t: t, expectedURL: "/vps/example-vps/settings/blockVpsMailPorts", expectedMethod: "GET", statusCode: 200, response: apiResponse}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := SettingsRepository{Client: *client}

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
	const apiRequest = `{"setting":{"name":"blockVpsMailPorts","dataType":"boolean","readOnly":false,"value":{"valueBoolean":true,"valueString":""}}}` //nolint
	server := mockServer{t: t, expectedURL: "/vps/example-vps/settings/blockVpsMailPorts", expectedMethod: "PUT", statusCode: 204, expectedRequest: apiRequest}
	client, tearDown := server.getClient()
	defer tearDown()
	repo := SettingsRepository{Client: *client}

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
