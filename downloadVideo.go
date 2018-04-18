package main

import "github.com/knadh/go-get-youtube/youtube"

func downloadVideo(video youtube.Video, id string) error {
	option := &youtube.Option{
		Resume: true,
		Mp3:    true,
	}
	err := video.Download(0, id+".mp4", option)
	if err != nil {
		return err
	}
	return nil
}
