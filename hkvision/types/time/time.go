package time_streaming

/**
 * @author : Benjamin Lam
 * @created : 9/29/21, Friday
**/
type Time struct {
	TimeMode          string `xml:"timeMode" json:"timeMode"`
	LocalTime         string `xml:"localTime" json:"localTime"`
	TimeZone          string `xml:"timeZone" json:"timeZone"`
	SatelliteInterval int    `xml:"satelliteInterval" json:"satelliteInterval"`
	CarrierInterval   int    `xml:"carrierInterval" json:"carrierInterval"`
}
