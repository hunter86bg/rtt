package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

var URL = os.Getenv("URL")
// URL := "https://liveengage.liveperson.net"

func getRTT() float64 {
    start := time.Now()
    resp, err := http.Get(URL)
    if err != nil {
        fmt.Println("Error:", err)
        return -1
    }
    defer resp.Body.Close()

    return time.Since(start).Seconds()
}

func average(lst []float64) float64 {
    sum := 0.0
    for _, val := range lst {
        sum += val
    }
    return sum / float64(len(lst))
}

func main() {
    var values []float64

    for i := 0; i < 10; i++ {
        values = append(values, getRTT())
        time.Sleep(1 * time.Second)
    }

    avg := average(values)
    fmt.Printf("%.3f\n", avg)
}
