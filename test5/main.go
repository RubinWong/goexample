package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	MediaTypeAudio = iota
	MediaTypeVideo
	MediaTypeUnknown
)

type publishedStream struct {
	Peerid string `json:"peerId"`
	Id     uint64 `json:"streamId"`
	Kind   string `json:"mediaKind"`
	Mute   bool   `json:"muted"`
}

type publishedStreams struct {
	Streams []publishedStream `json:"publishedStreams"`
}

type NewPublish struct {
	Notify bool             `json:"notification"`
	Method string           `json:"method"`
	Data   publishedStreams `json:"data"`
}

func main() {
	notiMsg := NewPublish{
		Notify: true,
		Method: "newPublish",
	}
	data := publishedStream{
		Peerid: "123",
		Id:     456,
		Mute:   false,
	}

	data.Kind = "video"
	notiMsg.Data.Streams = append(notiMsg.Data.Streams, data)

	fmt.Println("len of streams: ", len(notiMsg.Data.Streams))
	x, _ := json.Marshal(notiMsg)
	fmt.Println(string(x))

	// json.Marshal()
	// json.Unmarshal()
	{
		fmt.Println(time.Now().UnixNano())
		fmt.Println(time.Now().Nanosecond())
		fmt.Println(time.Now().Unix())
	}

	// {
	// 	a := 1
	// 	b := 2
	// 	defer calc("1", a, calc("10", a, b))
	// 	a = 0
	// 	defer calc("2", a, calc("20", a, b))
	// 	b = 1
	// }

	{
		s := make([]int, 5)
		s = append(s, 1, 2, 3)
		fmt.Println("slice: ", s)
	}
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
