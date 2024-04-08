package myffmpeg

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ConvertVideo(inputFile, outputPattern string) (string, error) {

	err := ffmpeg.Input(inputFile).
		Output(outputPattern, ffmpeg.KwArgs{"vf": "fps=1/5,scale=350:-1", "q:v": "1", "c:v": "libwebp", "lossless": "1"}).
		Run()

	if err != nil {
		return "ERROR", err
	}

	return "OK", nil

}
