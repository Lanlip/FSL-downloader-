package main

import (
	"fmt"
	"os"
	"io"
	"github.com/sqweek/dialog"
	"time"
)

func main() {
	fmt.Println("discord.gg/standkr Linore Modz\nGrand Theft Auto V 파일을 선택하세요.")
	dirPath, err := dialog.Directory().Title("디렉토리를 선택하세요").Browse()
	if err != nil {
		fmt.Println("디렉토리 선택에 오류가 발생했습니다:", err)
		time.Sleep(3 * time.Second)
		return
	}
	
	fmt.Println("선택된 디렉토리:", dirPath)

	sourceFile := "version.dll" 
	destFile := dirPath + "\\version.dll"

	err = copyFile(sourceFile, destFile)
	if err != nil {
		fmt.Println("파일을 복사하는데 오류가 발생했습니다:", err)
		time.Sleep(3 * time.Second)
	} else {
		fmt.Println("파일을 성공적으로 이동했습니다!")
		time.Sleep(3 * time.Second)
	}
}

func copyFile(src, dst string) error {
	
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()


	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}
