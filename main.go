package main

import (
	"fmt"
	"os"
	"strings"
)

func formatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Erro ao ler diretório: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%-50s %-10s %s\n", "Nome", "Tamanho", "Data Modificação")
	fmt.Println(strings.Repeat("-", 80))

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		size := "<DIR>"
		if !info.IsDir() {
			size = formatSize(info.Size())
		}

		modTime := info.ModTime().Format("02/01/2006 15:04:05")

		fmt.Printf("%-50s %-10s %s\n", info.Name(), size, modTime)
	}
}
