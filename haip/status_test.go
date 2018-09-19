package haip

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseStatusReportBody(t *testing.T) {

	// example output, but shortened for brevity
	data := []byte(`<transip>
	<item>
		<key xsi:type="xsd:int">80</key>
		<value SOAP-ENC:arrayType="ns2:Map[4]" xsi:type="SOAP-ENC:Array">
			<item xsi:type="ns2:Map">
				<item>
					<key xsi:type="xsd:string">vpsName</key>
					<value xsi:type="xsd:string">example-vps</value>
				</item>
				<item>
					<key xsi:type="xsd:string">ipVersions</key>
					<value xsi:type="ns2:Map">
						<item>
							<key xsi:type="xsd:string">ipv4</key>
							<value xsi:type="ns2:Map">
								<item>
									<key xsi:type="xsd:string">lb0.ams0</key>
									<value xsi:type="ns2:Map">
										<item>
											<key xsi:type="xsd:string">loadBalancerIp</key>
											<value xsi:type="xsd:string">1.2.3.4</value>
										</item>
										<item>
											<key xsi:type="xsd:string">state</key>
											<value xsi:type="xsd:string">up</value>
										</item>
										<item>
											<key xsi:type="xsd:string">lastChange</key>
											<value xsi:type="xsd:string">2018-08-30 20:32:33</value>
										</item>
										<item>
											<key xsi:type="xsd:string">lastChangeTimestamp</key>
											<value xsi:type="xsd:string">1535653953</value>
										</item>
									</value>
								</item>
								<item>
									<key xsi:type="xsd:string">lb0.rtm0</key>
									<value xsi:type="ns2:Map">
									  innerXML is removed for brevity
									</value>
								</item>
							</value>
						</item>
						<item>
							<key xsi:type="xsd:string">ipv6</key>
							<value xsi:type="ns2:Map">
								innerXML is removed for brevity
							</value>
						</item>
					</value>
				</item>
			</item>
			<item xsi:type="ns2:Map">
				<item>
					<key xsi:type="xsd:string">vpsName</key>
					<value xsi:type="xsd:string">example-vps2</value>
				</item>
				<item>
					<key xsi:type="xsd:string">ipVersions</key>
					<value xsi:type="ns2:Map">
						innerXML is removed for brevity
					</value>
				</item>
			</item>
		</value>
	</item>
	<item>
		<key xsi:type="xsd:int">443</key>
		<value SOAP-ENC:arrayType="ns2:Map[4]" xsi:type="SOAP-ENC:Array">
			innerXML removed for brevity
		</value>
	</item>
</transip>`)

	sr, err := parseStatusReportBody(data)
	if err != nil {
		t.Fatal(err.Error())
	}

	// go over returned body and see if all essential structs are there
	assert.Equal(t, 2, len(sr.PortConfiguration))
	assert.Equal(t, 80, sr.PortConfiguration[0].Port)
	assert.Equal(t, 443, sr.PortConfiguration[1].Port)
	assert.Equal(t, 2, len(sr.PortConfiguration[0].VPS))
	assert.Equal(t, "example-vps", sr.PortConfiguration[0].VPS[0].Name)
	assert.Equal(t, "example-vps2", sr.PortConfiguration[0].VPS[1].Name)
	assert.Equal(t, 2, len(sr.PortConfiguration[0].VPS[0].IPVersion))
	assert.Equal(t, "ipv4", sr.PortConfiguration[0].VPS[0].IPVersion[0].Version)
	assert.Equal(t, "ipv6", sr.PortConfiguration[0].VPS[0].IPVersion[1].Version)
	assert.Equal(t, 2, len(sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer))
	assert.Equal(t, "lb0.ams0", sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[0].Name)
	assert.Equal(t, "lb0.rtm0", sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[1].Name)
	assert.Equal(t, net.IP{1, 2, 3, 4}, sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[0].IPAddress)
	assert.Equal(t, "up", sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[0].State)
	assert.Equal(t, time.Unix(1535653953, 0), sr.PortConfiguration[0].VPS[0].IPVersion[0].LoadBalancer[0].LastChange)
}
