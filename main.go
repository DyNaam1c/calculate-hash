package main

import (
    "bufio"
    "time"
    "crypto/sha256"
    "fmt"
    "io"
    "os"
    "path/filepath"

    "github.com/fatih/color"
)

func main() {
    for {
        fmt.Print(color.MagentaString("[*] Enter file path: "))
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        path := scanner.Text()

        file, err := os.Open(path)
        if err != nil {
            fmt.Printf(color.RedString("[*] Error opening file: %v\n", err))
            continue
        }
        defer file.Close()

        hash := sha256.New()
        if _, err := io.Copy(hash, file); err != nil {
            fmt.Printf(color.RedString("[*] Error grabbing hash: %v\n", err))
            continue
        }

        absPath, err := filepath.Abs(path)
        if err != nil {
            fmt.Printf(color.RedString("[*] Error getting path: %v\n", err))
            continue
        }

        fileInfo, err := os.Stat(path)
        if err != nil {
            fmt.Printf(color.RedString("[*] Error getting file info: %v\n", err))
            continue
        }

        fileSize := fileInfo.Size()

        fileSizeStr := ""
        if fileSize < 1024 {
            fileSizeStr = fmt.Sprintf("%d bytes", fileSize)
        } else if fileSize < 1024*1024 {
            fileSizeStr = fmt.Sprintf("%.2f KB", float64(fileSize)/1024)
        } else if fileSize < 1024*1024*1024 {
            fileSizeStr = fmt.Sprintf("%.2f MB", float64(fileSize)/(1024*1024))
        } else {
            fileSizeStr = fmt.Sprintf("%.2f GB", float64(fileSize)/(1024*1024*1024))
        }

        fmt.Print(color.YellowString("\n[*] Starting..."))
        time.Sleep(1 * time.Second)

        fmt.Print(color.YellowString("\n[*] Getting Hash..."))
        time.Sleep(1 * time.Second)

        fmt.Print(color.YellowString("\n[*] Done!"))
        time.Sleep(1 * time.Second)

        fmt.Printf("\n%s %s\n", color.WhiteString("Algorithm:      "), color.WhiteString("SHA256"))
        fmt.Printf("%s %s\n", color.WhiteString("Hash:           "), color.GreenString("%x", hash.Sum(nil)))
        fmt.Printf("%s %s\n", color.WhiteString("Path:           "), color.WhiteString(absPath))
        fmt.Printf("%s %s\n\n", color.WhiteString("File size:      "), color.WhiteString(fileSizeStr))

        for {
            fmt.Print(color.YellowString("[*] Close program? (y/n): "))
            scanner := bufio.NewScanner(os.Stdin)
            scanner.Scan()
            answer := scanner.Text()

            if answer == "y" {
                return
            } else if answer == "n" {
                break
            }
        }
    }
}
