package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
    timeFormat := s[len(s)-2:]
    hour, err := strconv.Atoi(s[:2])
    if err != nil {
        return "Error While Convertion String To Int"
    }
    if timeFormat == "AM" {
        hour = hour % 12
        return fmt.Sprintf("%02d", hour) + s[2:len(s)-2]
    } else if timeFormat == "PM" {
        if hour < 12 {
            hour += 12
        }
        return fmt.Sprintf("%02d", hour) + s[2:len(s)-2]
    } else {
        return "Format '" + timeFormat + "' not recognized!"
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
