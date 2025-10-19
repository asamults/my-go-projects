package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    "strings"
    "sync"
    "time"
)

type FileStats struct {
    Lines int
    Words int
    Chars int
}

func countFileStats(filename string) (FileStats, error) {
    file, err := os.Open(filename)
    if err != nil {
        return FileStats{}, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    stats := FileStats{}

    for scanner.Scan() {
        stats.Lines++
        line := scanner.Text()
        stats.Words += len(strings.Fields(line))
        stats.Chars += len(line) + 1 // +1 для символа перевода строки
    }

    return stats, scanner.Err()
}

func main() {
    start := time.Now()

    jsonOutput := flag.Bool("json", false, "вывод в JSON-формате")
    flag.Parse()
    args := flag.Args()

    if len(args) == 0 {
        fmt.Println("Использование: linecounter [опции] <файл1> <файл2> ...")
        flag.PrintDefaults()
        return
    }

    results := make(map[string]FileStats)
    var mu sync.Mutex
    var wg sync.WaitGroup

    for _, filename := range args {
        wg.Add(1)
        go func(f string) {
            defer wg.Done()
            stats, err := countFileStats(f)
            if err != nil {
                log.Printf("Ошибка при чтении %s: %v", f, err)
                return
            }
            mu.Lock()
            results[f] = stats
            mu.Unlock()
        }(filename)
    }

    wg.Wait() // Ждём, пока все горутины завершатся
    elapsed := time.Since(start)

    if *jsonOutput {
        fmt.Print("{\n")
        for name, stats := range results {
            fmt.Printf("  \"%s\": {\"lines\": %d, \"words\": %d, \"chars\": %d},\n", name, stats.Lines, stats.Words, stats.Chars)
        }
        fmt.Printf("  \"elapsed_ms\": %d\n}\n", elapsed.Milliseconds())
    } else {
        for name, stats := range results {
            fmt.Printf("%s: %d строк(и), %d слов(а), %d символ(ов)\n", name, stats.Lines, stats.Words, stats.Chars)
        }
        fmt.Printf("Время выполнения: %d мс\n", elapsed.Milliseconds())
    }
}
