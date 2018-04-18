package main

import "github.com/knadh/go-get-youtube/youtube"

func getVideoStream(id string) (youtube.Video, error) {
	video, err := youtube.Get(id)
	if err != nil {
		return video, err
	}
	return video, nil
}
