package structs

type RequestPayload struct {
	ServiceType          string `json:"serviceType"`
	OriginTon            uint8  `json:"origin.ton,omitempty"`
	OriginNpi            uint8  `json:"origin.npi,omitempty"`
	OriginNumber         string `json:"origin.number"`
	DestinationTon       uint8  `json:"destination.ton,omitempty"`
	DestinationNpi       uint8  `json:"destination.npi,omitempty"`
	DestinationNumber    string `json:"destination.number"`
	ValidityPeriod       string `json:"validity_period"`
	ScheduleDeliveryTime string `json:"schedule_delivery_time"`
	ProtocolID           uint8  `json:"protocol_id"`
	EsmeClass            uint8  `json:"esmeClass"`
	PriorityFlag         uint8  `json:"priority_flag"`
	RegisteredDelivery   uint8  `json:"registered_delivery"`
	ReplaceIfPresentFlag uint8  `json:"replace_if_present_flag"`
	Data                 string `json:"data"`
	DataHeaderIndicator  uint8  `json:"data_header_indicator"`
	DataCodingScheme     uint8  `json:"data_coding_scheme"`
	DataLength           uint16 `json:"data_length"`
	MessageType          uint8  `json:"messagetype"`
	TLVTag               int    `json:"TLV_TAG"`
	TLVLength            int    `json:"TLV_LENGTH"`
	TLVValue             string `json:"TLV_VALUE"`
}
