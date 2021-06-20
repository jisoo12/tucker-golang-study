package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 찾은 라인 정보
type LineInfo struct {
	lineNo int
	line string
}

// 파일 내 라인 정보
type FindInfo struct {
	filename string
	lines []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.3 word filepath")
		return
	}

	word := os.Args[1]
	// files := os.Args[2]		// macox, linux 환경에는 안먹힘..
	files := "*.txt"
	
	cnt, ch := FindWordInAllFiles(word, files)
	recvCnt := 0
	for findInfo := range ch {
		fmt.Println(findInfo.filename)
		fmt.Println("------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("------------------------------")
		fmt.Println()
		recvCnt++
		if recvCnt == cnt {
			// all received
			break
		}
	}
}

func GetFileList(pattern string) ([]string, error) {
	// fmt.Println(pattern)
	fileList := []string{}
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// fmt.Println(info.Name())
			matched, _ := filepath.Match(pattern, info.Name())
			if matched {
				fileList = append(fileList, path)
			}
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return fileList, nil
}

func FindWordInAllFiles(word, path string) (int, chan FindInfo) {
	fileList, err := GetFileList(path)
	// fmt.Println(fileList)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err:", err)
		return 0, nil
	}

	ch := make(chan FindInfo)
	cnt := len(fileList)

	for _, filename := range fileList {
		go FindWordInFile(word, filename, ch)
	}

	return cnt, ch
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		ch <- findInfo
		return
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo
}