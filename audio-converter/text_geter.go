package main

import (
	"fmt"
	"os/exec"
)

func TextGeter(url string, lang string) {

	//cmd := exec.Command("youtube-dl", "--skip-download", "--write-sub", "--sub-lang=en", "--convert-subs=srt", "--verbose", url)
	//cmd := exec.Command("yt-dlp", "--skip-download", "--write-sub", "--sub-lang", "en", "--convert-subs", "srt", "--verbose", url)
	cmd := exec.Command("yt-dlp", "--write-subs", "--sub-langs", lang, url)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Output:\n%s\n", output)
}
