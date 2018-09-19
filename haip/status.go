package haip

import (
	"encoding/xml"
	"fmt"
	"net"
	"strconv"
	"time"
)

// This part of the HA-IP package is mainly concerned with parsing the HA-IP
// status reports. The structure of the XML the TransIP SOAP API returns is
// relatively complex which the structs represent

type statusXMLOuter struct {
	PortConf []struct { // port configuration
		Port  string `xml:"key"` // port number
		Value struct {
			Item []struct {
				Item []struct {
					Key   string `xml:"key"` // vpsName, ipVersions
					Value struct {
						Cont string `xml:",innerxml"`
					} `xml:"value"`
				} `xml:"item"`
			} `xml:"item"`
		} `xml:"value"`
	} `xml:"item"`
}

type statusXMLInner struct {
	IPVersions []struct {
		Version string `xml:"key"` // ipv4 or ipv6
		Value   struct {
			Item []struct {
				Key   string `xml:"key"` // name of loadbalancer
				Value struct {
					Item []statusKV `xml:"item"` // fields like loadBalancerIp and status
				} `xml:"value"`
			} `xml:"item"`
		} `xml:"value"`
	} `xml:"item"`
}

type statusKV struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

// multi-level structs to represent the status report
type statusReportLb struct {
	Name       string
	IPAddress  net.IP
	State      string
	LastChange time.Time
}

type statusReportIPVersion struct {
	Version      string
	LoadBalancer []statusReportLb
}

type statusReportVPS struct {
	Name      string
	IPVersion []statusReportIPVersion
}

type statusReportPortConfiguration struct {
	Port int
	VPS  []statusReportVPS
}

func parseStatusReportLb(name, ipVersion string, kv []statusKV) (lb statusReportLb, err error) {

	lb.Name = name

	for _, x := range kv {
		switch x.Key {
		case "loadBalancerIp":
			if ipVersion == "ipv4" {
				lb.IPAddress = net.ParseIP(x.Value).To4()
			} else {
				lb.IPAddress = net.ParseIP(x.Value).To16()
			}
		case "state":
			lb.State = x.Value
		case "lastChangeTimestamp":
			var i int64
			if i, err = strconv.ParseInt(x.Value, 10, 64); err == nil {
				t := time.Unix(i, 0)
				lb.LastChange = t
			}
		case "lastChange":
			// ignore lastChange since it is the same as lastChangeTimestamp
			// anyway, but formatted as a string
		default:
			err = fmt.Errorf("unhandled field in parsing ip versions: %s", x.Key)
			return
		}
	}

	return
}

func parseStatusReportBody(data []byte) (StatusReport, error) {
	sr := StatusReport{}

	var v statusXMLOuter
	if err := xml.Unmarshal(data, &v); err != nil {
		return sr, err
	}

	// start going over statusXMLOuter body to parse each port configuration
	// these port configurations each have a key/value pair with the vpsName and
	// then a more complex structure for the statuses per loadbalancer per ipversion
	// these structures are represented by statusXMLInner
	for _, pc := range v.PortConf {
		srpc := statusReportPortConfiguration{}
		if port, err := strconv.ParseInt(pc.Port, 10, 64); err == nil {
			srpc.Port = int(port)
		}
		for _, x := range pc.Value.Item {
			srv := statusReportVPS{}
			for _, xx := range x.Item {
				switch xx.Key {
				case "vpsName":
					srv.Name = xx.Value.Cont
				case "ipVersions":
					ipv := statusReportIPVersion{}
					vv := statusXMLInner{}
					// encapsulate body into some elements for easier parsing
					xml.Unmarshal([]byte("<transip>"+xx.Value.Cont+"</transip>"), &vv)
					for _, xxx := range vv.IPVersions {
						ipv.Version = xxx.Version
						// parse loadbalancers from status report
						for _, xxxx := range xxx.Value.Item {
							lb, err := parseStatusReportLb(xxxx.Key, xxx.Version, xxxx.Value.Item)
							if err != nil {
								return sr, err
							}
							ipv.LoadBalancer = append(ipv.LoadBalancer, lb)
						}
						srv.IPVersion = append(srv.IPVersion, ipv)
					}
				default:
					return sr, fmt.Errorf("unhandled field in parsing key/values: %s", xx.Key)
				}
			}

			srpc.VPS = append(srpc.VPS, srv)
		}

		sr.PortConfiguration = append(sr.PortConfiguration, srpc)
	}

	return sr, nil
}
