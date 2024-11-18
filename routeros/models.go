package routeros

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type Interfaces struct {
	Wireless []WirelessInterface `json:"wireless"`
	Ethernet []EthernetInterface `json:"ethernet"`
}

type Ip struct {
	Arp []ARPItem `json:"arp"`
}

type EthernetInterface struct {
	Name            string `json:"name"`
	Speed           string `json:"speed"`
	FullDuplex      string `json:"full_duplex"`
	TxBitsBerSecond string `json:"tx_bits_ber_second"`
	RxBitsBerSecond string `json:"rx_bits_ber_second"`
}

type WirelessInterface struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Ssid              string `json:"ssid"`
	Channel           string `json:"channel"`
	TxCCQ             string `json:"tx_ccq"`
	RxCCQ             string `json:"rx_ccq"`
	MacAddress        string `json:"mac_address"`
	TxSignalStrength  string `json:"tx_signal_strength"`
	TxSignalStrength0 string `json:"tx_signal_strength0"`
	TxSignalStrength1 string `json:"tx_signal_strength1"`
	RxSignalStrength  string `json:"rx_signal_strength"`
	RxSignalStrength0 string `json:"rx_signal_strength0"`
	RxSignalStrength1 string `json:"rx_signal_strength1"`
	TxRate            string `json:"tx_rate"`
	RxRate            string `json:"rx_rate"`
	RadioName         string `json:"radio_name"`
	TxBitsBerSecond   string `json:"tx_bits_ber_second"`
	RxBitsBerSecond   string `json:"rx_bits_ber_second"`
}

type ARPItem struct {
	IPAddress   string `json:"ip_address"`
	MacAddress  string `json:"mac_address"`
	Interface   string `json:"interface"`
	HostComment string `json:"host_comment,omitempty"`
}

type RouterData struct {
	Interfaces Interfaces `json:"interfaces"`
	Ip         Ip         `json:"ip"`
}
