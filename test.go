package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/BenjaminLam1202/test-go-config-cameras/hkvision/client"
	"github.com/BenjaminLam1202/test-go-config-cameras/hkvision/types"
	"github.com/golang-module/carbon"
)

/**
 * @author : Donald Trieu
 * @created : 9/25/21, Saturday
**/
// const layout = "2021-10-29T15:19:15+07:00"

// const layout = "Jan 2, 2006 at 3:04pm (MST)"
func GetInformationFromURI(schannel string, susername string, spass string, s string) (string, string, string, string) {
	var username, pass, ip, channel string
	reUserPass := regexp.MustCompile(`(?:([^\s@\/]+?)[@])`)
	userPassMatch := reUserPass.FindStringSubmatch(s)
	reIp := regexp.MustCompile(`[@]([^\s\/:]+)`)
	ipMatch := reIp.FindStringSubmatch(s)
	reChannel := regexp.MustCompile(`[/]([^\s\/:]+?$)`)
	channelMatch := reChannel.FindStringSubmatch(s)
	if userPassMatch != nil {
		subUserPass := userPassMatch[1]
		username = strings.Split(subUserPass, ":")[0]
		pass = strings.Split(subUserPass, ":")[1]
		ip = ipMatch[1]
		channel = channelMatch[1]
	}
	if !strings.Contains(s, "Channels") {
		if schannel != "" {
			channel = schannel
		}
		channel = "101"
	}
	if !strings.Contains(s, "@") {
		if susername != "" || spass != "" {
			reaIp := regexp.MustCompile(`(?:[/]([^\s\/]+?)[:])`)
			ipMatch := reaIp.FindStringSubmatch(s)
			username = susername
			pass = spass
			ip = ipMatch[1]
		}
		errors.New("username or password is missing")
	}
	return channel, username, pass, ip
}

func GetDateFromString(syear int, smonth int, sday int, shour int, smin int, sns int, s string) string {
	var year, month, day, hour, min, ns int

	if s == "" {
		if syear == 0 {
			syear := carbon.Now().Year()
			year = syear
			fmt.Println(year)
		} else {
			year = syear
		}
		if smonth == 0 {
			smonth := carbon.Now().Month()
			month = smonth
			fmt.Println(month)
		} else {
			month = smonth
		}
		if sday == 0 {
			sday := carbon.Now().Day()
			day = sday
			fmt.Println(day)
		} else {
			day = sday
		}
		if shour == 0 {
			shour := carbon.Now().Hour()
			hour = shour
			fmt.Println(hour)
		} else {
			hour = shour
		}
		if smin == 0 {
			smin := carbon.Now().Minute()
			min = smin
			fmt.Println(min)
		} else {
			min = smin
		}
		if sns == 0 {
			sns := carbon.Now().Second()
			ns = sns
			fmt.Println(ns)
		} else {
			ns = sns
		}
		sstring := carbon.CreateFromDateTime(year, month, day, hour, min, ns).ToDateTimeString()
		stringdate := carbon.Parse(sstring).ToRfc3339String()
		return stringdate
	}
	stringdate := carbon.Parse(s).ToRfc3339String()
	return stringdate
}
func main() {
	newClient, _ := client.NewClient(types.ConstructClient{
		Username: "admin",
		Password: "AI_team123",
		Host:     "192.168.1.150",
		Proto:    "http",
	})
	// channels, err := newClient.GetStreamChannels()
	// if err != nil {
	// 	log.Print(err)
	// }
	// bytesa, _ := json.Marshal(channels)
	// log.Printf("Result : %s\n", string(bytesa))
	channel, err := newClient.GetIntegrate()
	if err != nil {
		log.Print(err)
	}
	bytes, _ := json.Marshal(channel)
	log.Printf("Result : %s\n", string(bytes))

	// *channel.Video.FixedQuality = 40
	// channel.LocalTime = "2021-10-29T15:10:33+07:00"
	// channel.CarrierInterval = 0
	// channel.SatelliteInterval = 0

	// s1 := GetDateFromString(0, 0, 0, 0, 0, 0, "")
	// channel.LocalTime = "2004-09-01T13:56:37+07:00"

	// channel.LocalTime = "2021-08-29T22:55:20+07:00"
	// RFC3339local := carbon.CreateFromDateTime(2020, 0, 0, 0, 0, 15).ToDateTimeString()
	// RFC3339local1 := carbon.Parse(RFC3339local).ToRfc3339String()
	// // RFC3339local1 := carbon.Parse("2021-11-30T00:00:00+07:00").ToRfc3339String()
	// fmt.Println(RFC3339local1)

	// year := carbon.Now().Month()
	// fmt.Println(year)
	// channel.LocalTime = RFC3339local1
	//2020-12-31T00:00:00+07:00
	//2021-10-29T15:35:27+07:00
	*channel.ONVIF.Enable = true
	// s1, _ := time.Parse(layout, "Feb 4, 2014 at 6:05pm (PST)")
	// fmt.Println(s1)
	// channel.Video.SVC.Enabled = false
	// channel.Video.SVC.Enabled = false
	// channel.Video.SVC.Enabled = true
	// channel.Video.SVC.SVCMode = ""
	// prechannel, err := newClient.GetSingleStreamChannel("2")
	// if err != nil {
	// 	log.Print(err)
	// }
	// fmt.Println("---------pre-test--------------")
	// fmt.Println(channel.Video.FixedQuality)
	// fmt.Println(*channel.Video.FixedQuality)
	// fmt.Println(&channel.Video.FixedQuality)
	// fmt.Println("---------pre test--------------")

	// fmt.Println(*channel.Video.FixedQuality)

	// 	channel.Video.ConstantBitRate = 67
	// 	if channel.Video.ConstantBitRate <= 0 {
	// 		channel.Video.ConstantBitRate = prechannel.Video.ConstantBitRate
	// 	}
	// 	channel.Video.SVC.Enabled = true
	// 	channel.Video.SVC.SVCMode = "manual"
	// 	if channel.Video.SVC.SVCMode == "manual" {
	// 		channel.Video.SVC.SVCMode = prechannel.Video.SVC.SVCMode
	// 	}

	// log.Printf("maxFrameRate: %v\n", channel.Video.ConstantBitRate)
	err = newClient.PutIntegrate(channel)
	if err != nil {
		log.Print(err)
	}
}
