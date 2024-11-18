package routeros

import "github.com/go-routeros/routeros"

func FetchRouterData(client *routeros.Client) (*RouterData, error) {
	data := &RouterData{}

	interfaceWirelessPrintReply, err := client.RunArgs([]string{
		"/interface/wireless/print",
		"=.proplist=.id,name,mac-address",
	})
	if err != nil {
		return nil, err
	}
	for _, re1 := range interfaceWirelessPrintReply.Re {
		id := re1.Map[".id"]
		name := re1.Map["name"]
		wirelessInterface := WirelessInterface{
			Id:         id,
			Name:       name,
			MacAddress: re1.Map["mac-address"],
		}

		cmd2 := []string{
			"/interface/wireless/monitor",
			"=.id=" + id,
			"=.proplist=ssid,channel,tx-ccq,rx-ccq,tx-signal-strength,tx-signal-strength-ch0,tx-signal-strength-ch1," +
				"signal-strength,signal-strength-ch0,signal-strength-ch1,tx-rate,rx-rate,radio-name",
			"=once=",
		}
		interfaceMonitorReply, err2 := client.RunArgs(cmd2)
		if err2 != nil {
			return nil, err2
		}
		for _, re := range interfaceMonitorReply.Re {
			wirelessInterface.Ssid = re.Map["ssid"]
			wirelessInterface.Channel = re.Map["channel"]
			wirelessInterface.TxCCQ = re.Map["tx-ccq"]
			wirelessInterface.RxCCQ = re.Map["rx-ccq"]
			wirelessInterface.TxSignalStrength = re.Map["tx-signal-strength"]
			wirelessInterface.TxSignalStrength0 = re.Map["tx-signal-strength-ch0"]
			wirelessInterface.TxSignalStrength1 = re.Map["tx-signal-strength-ch1"]
			wirelessInterface.RxSignalStrength = re.Map["signal-strength"]
			wirelessInterface.RxSignalStrength0 = re.Map["signal-strength-ch0"]
			wirelessInterface.RxSignalStrength1 = re.Map["signal-strength-ch1"]
			wirelessInterface.TxRate = re.Map["tx-rate"]
			wirelessInterface.RxRate = re.Map["rx-rate"]
			wirelessInterface.RadioName = re.Map["radio-name"]
		}

		cmd3 := []string{
			"/interface/monitor-traffic",
			"=interface=" + name,
			"=.proplist=tx-bits-per-second,rx-bits-per-second",
			"=once=",
		}
		monitorTrafficReply, err3 := client.RunArgs(cmd3)
		if err3 != nil {
			return nil, err3
		}
		for _, re3 := range monitorTrafficReply.Re {
			wirelessInterface.TxBitsBerSecond = re3.Map["tx-bits-per-second"]
			wirelessInterface.RxBitsBerSecond = re3.Map["rx-bits-per-second"]
		}

		data.Interfaces.Wireless = append(data.Interfaces.Wireless, wirelessInterface)
	}

	ethernetReply, err := client.RunArgs([]string{
		"/interface/ethernet/print",
		"=.proplist=name,speed,full-duplex",
	})
	if err != nil {
		return nil, err
	}
	for _, re := range ethernetReply.Re {
		name := re.Map["name"]

		ethernetInterface := EthernetInterface{
			Name:       name,
			Speed:      re.Map["speed"],
			FullDuplex: re.Map["full-duplex"],
		}

		cmd2 := []string{
			"/interface/monitor-traffic",
			"=interface=" + name,
			"=.proplist=tx-bits-per-second,rx-bits-per-second",
			"=once=",
		}
		monitorTrafficReply, err2 := client.RunArgs(cmd2)
		if err2 != nil {
			return nil, err2
		}
		for _, re2 := range monitorTrafficReply.Re {
			ethernetInterface.TxBitsBerSecond = re2.Map["tx-bits-per-second"]
			ethernetInterface.RxBitsBerSecond = re2.Map["rx-bits-per-second"]
		}

		data.Interfaces.Ethernet = append(data.Interfaces.Ethernet, ethernetInterface)
	}

	arpReply, err := client.RunArgs([]string{
		"/ip/arp/print",
		"=.proplist=address,mac-address,interface,comment",
	})
	if err != nil {
		return nil, err
	}
	for _, re := range arpReply.Re {
		data.Ip.Arp = append(data.Ip.Arp, ARPItem{
			IPAddress:   re.Map["address"],
			MacAddress:  re.Map["mac-address"],
			Interface:   re.Map["interface"],
			HostComment: re.Map["comment"],
		})
	}

	return data, nil
}
