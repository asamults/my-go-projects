package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
    "time"
)

var clients = make(map[net.Conn]string)

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer ln.Close()

    fmt.Println("💬 TCP-чат сервер запущен на порту 8080...")

    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)

    conn.Write([]byte("Введите ваше имя:\n"))
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    clients[conn] = name
    fmt.Printf("✅ %s подключился\n", name)
    broadcast(fmt.Sprintf("📢 %s присоединился к чату\n", name), conn)

    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            disconnect(conn)
            return
        }

        msg = strings.TrimSpace(msg)
        if msg == "/quit" {
            conn.Write([]byte("👋 До встречи!\n"))
            disconnect(conn)
            return
        }

        timestamp := time.Now().Format("15:04:05")
        broadcast(fmt.Sprintf("[%s] %s: %s\n", timestamp, name, msg), conn)
    }
}

func broadcast(message string, sender net.Conn) {
    for conn := range clients {
        if conn != sender {
            conn.Write([]byte(message))
        }
    }
}

func disconnect(conn net.Conn) {
    name := clients[conn]
    fmt.Printf("❌ %s отключился\n", name)
    delete(clients, conn)
    broadcast(fmt.Sprintf("🚪 %s покинул чат\n", name), conn)
}
