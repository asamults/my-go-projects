package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    "time"
)

func countLines(filename string) (int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    count := 0
    for scanner.Scan() {
        count++
    }

    return count, scanner.Err()
}

func main() {
    start := time.Now()

    // Используем стандартный пакет flag для CLI-аргументов
    jsonOutput := flag.Bool("json", false, "вывод в JSON-формате")
    flag.Parse()
    args := flag.Args()

    if len(args) == 0 {
        fmt.Println("Использование: linecounter [опции] <файл1> <файл2> ...")
        flag.PrintDefaults()
        return
    }

    results := make(map[string]int)

    for _, filename := range args {
        count, err := countLines(filename)
        if err != nil {
            log.Printf("Ошибка при чтении %s: %v", filename, err)
            continue
        }
        results[filename] = count
    }

    elapsed := time.Since(start)

    if *jsonOutput {
        fmt.Print("{\n")
        for name, count := range results {
            fmt.Printf("  \"%s\": %d,\n", name, count)
        }
        fmt.Printf("  \"elapsed_ms\": %d\n}\n", elapsed.Milliseconds())
    } else {
        for name, count := range results {
            fmt.Printf("%s: %d строк(и)\n", name, count)
        }
        fmt.Printf("Время выполнения: %d мс\n", elapsed.Milliseconds())
    }
}
