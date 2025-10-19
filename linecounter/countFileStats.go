package main

import (
    "bufio"
    "os"
    "strings"
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
        stats.Chars += len(line) + 1
    }

    return stats, scanner.Err()
}
