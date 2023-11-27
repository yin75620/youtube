package main

import "testing"

func TestConvert(t *testing.T) {
	inputFile := "../cmd/youtubedr/【アニメ】秘密結社の平日ルーティン✰･.｡ ｡.･ﾟ✽.｡.・ﾟ.mp4"
	outputFile := "output.wav"
	convert(inputFile, outputFile)
}

func TestTextGeter(t *testing.T) {
	url := "https://www.youtube.com/watch?v=2xD8H2h7EIY"
	TextGeter(url, "ja")
}
