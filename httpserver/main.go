package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		fmt.Println("Done!", ctx.Err())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout!")
		fmt.Fprintf(w, "Hello!\n")
	}
}

func Video(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Video!\n")
	video, err := os.Open("video.mp4")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer video.Close()
	http.ServeContent(w, r, "video.mp4", time.Now(), video)
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/video", Video)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
}
