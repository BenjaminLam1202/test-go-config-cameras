package streaming

/**
 * @author : Donald Trieu
 * @created : 9/24/21, Friday
**/
type StreamStatus struct {
	TotalStreamingSessions     int `xml:"totalStreamingSessions" json:"totalStreamingSessions"`
	StreamingSessionStatusList []struct {
		StreamingSessionStatus struct {
			ClientAddress struct {
				IpAddress string `xml:"ipAddress" json:"ipAddress"`
			} `xml:"clientAddress" json:"clientAddress"`
		} `xml:"StreamingSessionStatus" json:"streamingSessionStatus"`
	} `xml:"StreamingSessionStatusList" json:"streamingSessionStatusList"`
}
