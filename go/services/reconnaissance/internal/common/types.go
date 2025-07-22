package common

import "time"

type Payload struct {
	Value   string `json:"Value"`
	Context string `json:"Context"`
}

type Vulnerability struct {
	Type        string
	Severity    string
	URL         string
	Parameter   string
	Payload     string
	Evidence    string
	Confidence  string
	CWE         string
	CVSS        float64
	Timestamp   time.Time
	Request     string
	Response    string
}

type Target struct {
	URL string
}
