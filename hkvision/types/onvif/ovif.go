package onvif

/**
 * @author : Benjamin Lam
 * @created : 9/29/21, Friday
**/
type Integrate struct {
	ONVIF *ONVIF `xml:"ONVIF" json:"ONVIF"`
	ISAPI *ISAPI `xml:"ISAPI" json:"ISAPI"`
}
type ONVIF struct {
	Enable *bool `xml:"enable" json:"enable"`
}
type ISAPI struct {
	Enable *bool `xml:"enable" json:"enable"`
}
