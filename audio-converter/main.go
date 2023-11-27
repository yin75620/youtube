package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 输入的mp4文件名和输出的wav文件名
	inputFile := "input.mp4"
	outputFile := "output.wav"
	convert(inputFile, outputFile)
}

func convert(inputFile, outputFile string) {

	// 使用FFmpeg执行转换操作
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-y", "-ar", "16000", outputFile)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("转换失败：%v\n", err)
		return
	}

	fmt.Printf("转换成功，输出文件：%s\n", outputFile)
}
