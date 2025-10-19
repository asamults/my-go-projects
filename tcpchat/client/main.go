package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Читаем приглашение от сервера
	prompt, _ := reader.ReadString('\n')
	fmt.Print(prompt)

	// Вводим имя пользователя
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Fprintln(conn, scanner.Text())
	}

	// Чтение сообщений с сервера в отдельной горутине
	go func() {
		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("❌ Соединение закрыто")
				os.Exit(0)
			}
			fmt.Print(msg)
		}
	}()

	// Отправка сообщений
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintln(conn, text)
		if text == "/quit" {
			fmt.Println("👋 Вы вышли из чата.")
			return
		}
	}
}
