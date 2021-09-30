package onvif

/**
 * @author : Benjamin Lam
 * @created : 9/29/21, Friday
**/
type Integrate struct {
	ONVIF *ONVIF `xml:"ONVIF" json:"ONVIF"`
}
type ONVIF struct {
	Enable *bool `xml:"enable" json:"enable"`
}
