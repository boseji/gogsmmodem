package gogsmmodem

import "time"

type Packet interface{}

// +ZPASR
type ServiceStatus struct {
	Status string
}

// +ZDONR
type NetworkStatus struct {
	Network string
}

// +CFUN
type PhoneFunction struct {
	Function int
}

// +CMTI
type MessageNotification struct {
	Storage string
	Index   int
}

// +CSCA
type SMSCAddress struct {
	Args []interface{}
}

// +CMGR
type Message struct {
	Index     int
	Status    string
	Telephone string
	Timestamp time.Time
	Body      string
	Last      bool
}

// +CPMS=?
type StorageAreas struct {
	Received []string
	Sent     []string
	New      []string
}

// +CPMS=...
type StorageInfo struct {
	UsedSpace1, MaxSpace1, UsedSpace2, MaxSpace2, UsedSpace3, MaxSpace3 int
}

// +CMGL
type MessageList []Message

// +CUSD
type USSDResponse struct {
	Command   string
	Response  string
	Processed []string
}

// Simple OK response
type OK struct{}

// Simple ERROR response
type ERROR struct{}

// Unknown
type UnknownPacket struct {
	Command string
	Args    []interface{}
}
