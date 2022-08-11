package main

import (
	"encoding/json"
	"fmt"
)



type AppData map[string]interface{}

type ConsumerOptions struct {
	ProducerId      string    `json:"producerId"`
	RTPCapabilities string    `json:"rtpCapabilities"`
	Paused          bool      `json:"paused"`
	PreferredLayers string    `json:"preferredLayers"`
	AppData         AppData   `json:"appData"`
}

func main() {
	opt := ConsumerOptions{
		ProducerId:      "producerId",
		RTPCapabilities: "rtpCapabilities",
		Paused:          false,
		PreferredLayers: "preferredLayers",
		AppData:         make(AppData),
	}

	opt.AppData["streamId"] = 1
	opt.AppData["index"] = 0

	reqData := struct{
		AppData
		Kind string `json:"kind"`
		RTPParameters 	string `json:"rtpParameters"`
		Type string `json:"type"`
		ConsumableRTPEncoding string `json:"consumableRtpEncoding"`
		Paused bool `json:"paused"`
		PreferedLayers string `json:"preferredLayers"`
	} {
		AppData: opt.AppData,
		Kind: "producer",
		RTPParameters: "",
		Type: "",
		ConsumableRTPEncoding: "",
		Paused: false,
		PreferedLayers: "",
	}

	b, err := json.Marshal(reqData)
	if  err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	reqData2 := map[string]interface{}{
		"kind":                   "producer",
		"rtpParameters":          "",
		"type":                   "",
		"consumableRtpEncodings": "",
		"paused":                 false,
		"preferredLayers":        "",
	}

	b2, err :=  json.Marshal(reqData2)
	fmt.Println(string(b2))
}
