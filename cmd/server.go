package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func startServer(port string) {
	// TCP 소켓 생성 및 지정된 포트에 바인딩
	// net.Listen = socket() + bind() + listen() 의 조합
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("서버 시작 실패: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("서버가 포트 %s에서 대기 중...\n", port)
	fmt.Println("클라이언트 연결을 기다리고 있습니다.")

	for {
		// 클라이언트 연결 요청 대기 (블로킹)
		// Accept() = 시스템 콜 accept()와 동일
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("연결 수락 실패: %v\n", err)
			continue
		}

		fmt.Printf("클라이언트 연결됨: %s\n", conn.RemoteAddr())

		// 각 클라이언트를 별도 goroutine에서 처리
		// 이렇게 하면 여러 클라이언트 동시 처리 가능
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// 클라이언트로부터 메시지 읽기용 스캐너
	scanner := bufio.NewScanner(conn)
	
	// 서버에서 메시지 보내기용 채널과 goroutine
	go handleServerInput(conn)

	// 클라이언트 메시지 수신 루프
	for scanner.Scan() {
		message := scanner.Text()
		if strings.TrimSpace(message) == "quit" {
			fmt.Println("클라이언트가 연결을 종료했습니다.")
			break
		}
		
		fmt.Printf("클라이언트 메시지: %s\n", message)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("메시지 읽기 오류: %v\n", err)
	}
}

func handleServerInput(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("메시지를 입력하세요 (quit으로 종료):")
	
	for scanner.Scan() {
		input := scanner.Text()
		
		if strings.TrimSpace(input) == "quit" {
			conn.Write([]byte("quit\n"))
			break
		}
		
		// 클라이언트에게 메시지 전송
		// Write() = 시스템 콜 send()와 동일
		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Printf("메시지 전송 실패: %v\n", err)
			break
		}
	}
}
