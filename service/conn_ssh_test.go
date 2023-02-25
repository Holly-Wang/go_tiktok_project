package service

import (
	"fmt"
	"testing"
)


func TestRun(t *testing.T) {
	username := "root"
	password := "Freedom9"
	ip := "49.232.155.203"
	port := "22"
	client := NewSSHClient(username, password, ip, port)

	//上传
	filename := "..//video_data//111//bear//bear.mp4"
	n, err := client.UploadFile(filename, "/root/plalyy/go/src/go_tiktok_project/video_data/111/bear/bear3.mp4")
	if err != nil {
		fmt.Printf("n: %v\n", n)
		fmt.Printf("upload failed: %v\n", err)
	}

}
