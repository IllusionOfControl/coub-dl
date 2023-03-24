package coub

import (
	"coub-dl/api"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"path/filepath"
)

/*
   if (loop) coub.loop(loop)
   if (crop) coub.crop(crop)
   if (scale) coub.scale(scale)
   if (time) coub.addOption('-t', time)
   if (details) coub.on('info', console.log)
*/

type SaveOptions struct {
	Loop string
}

type Media struct {
	Title    string
	Height   int
	Width    int
	VideoUrl string
	AudioUrl string
}

func NewCoubMedia(metadata map[string]interface{}) *Media {
	title := metadata["title"].(string)
	formats := metadata["file_versions"].(map[string]interface{})
	html5Files := formats["html5"].(map[string]interface{})
	videoQuality := html5Files["video"].(map[string]interface{})
	audioQuality := html5Files["audio"].(map[string]interface{})

	media := Media{
		Title:    title,
		VideoUrl: getUrlFromMetadata(videoQuality),
		AudioUrl: getUrlFromMetadata(audioQuality),
	}
	return &media
}

type Options struct {
	Loop  string
	Crop  string
	Scale string
	Time  string
}

func (m *Media) DownloadFile(url string) *os.File {
	ext := getExtFromUrl(url)
	filenamePattern := fmt.Sprintf("%s.*.%s", m.Title, ext)
	tempFile, err := os.CreateTemp("", filenamePattern)
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Base(tempFile.Name())
	log.Printf("downloading media to %s", filename)

	data, err := api.FetchCoubFile(url)
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := tempFile.Write(data); err != nil {
		log.Fatal(err)
	}
	err = tempFile.Close()
	if err != nil {
		log.Fatalln(err)
	}

	return tempFile
}

func (m *Media) Save(config SaveOptions) {
	videoFile := m.DownloadFile(m.VideoUrl)
	defer os.Remove(videoFile.Name())

	audioFile := m.DownloadFile(m.AudioUrl)
	defer os.Remove(audioFile.Name())

	outputFilename := fmt.Sprintf("./%s.mp4", m.Title)

	videoStream := ffmpeg.Input(videoFile.Name(), ffmpeg.KwArgs{"stream_loop": config.Loop})
	audioStream := ffmpeg.Input(audioFile.Name(), nil)

	err := ffmpeg.Output([]*ffmpeg.Stream{
		videoStream,
		audioStream,
	}, outputFilename, ffmpeg.KwArgs{"c": "copy", "shortest": ""}).
		OverWriteOutput().
		Run()

	if err != nil {
		log.Fatalln(err)
	}
}
