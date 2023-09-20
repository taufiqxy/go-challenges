package main

import (
    "bufio"
    "fmt"
    "sort"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
 
func sum(arr []int32) int64{
    var sum int64
    for _, data := range arr {
        sum += int64(data)
    }
    return sum
}

func miniMaxSum(arr []int32) {
    // Write your code here
    // sort array
    sort.Slice(arr, func(i, j int) bool {return arr[i] < arr[j]})
    
    // min and max sum
    fmt.Print(sum(arr[:len(arr)-1]), sum(arr[1:]))

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
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
