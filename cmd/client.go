package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func startClient(address string) {
	// 서버에 TCP 연결 시도
	// net.Dial = socket() + connect() 의 조합
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("서버 연결 실패: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Printf("서버에 연결되었습니다: %s\n", address)
	fmt.Println("메시지를 입력하세요 (quit으로 종료):")

	// 서버로부터 메시지 수신을 위한 별도 goroutine
	go handleServerMessages(conn)

	// 사용자 입력 처리
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		
		if strings.TrimSpace(input) == "quit" {
			fmt.Println("연결을 종료합니다.")
			break
		}

		// 서버에게 메시지 전송
		// Write() = 시스템 콜 send()와 동일
		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Printf("메시지 전송 실패: %v\n", err)
			break
		}
	}
}

func handleServerMessages(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	
	for scanner.Scan() {
		message := scanner.Text()
		
		if strings.TrimSpace(message) == "quit" {
			fmt.Println("서버가 연결을 종료했습니다.")
			os.Exit(0)
		}
		
		fmt.Printf("서버 메시지: %s\n", message)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("서버 메시지 읽기 오류: %v\n", err)
	}
}
