package main

import (
"fmt"
"net"
 "os"
"time"
"encoding/json"
)


func main() {

        type statusMessage struct {
                State     string
                CurrentFloor    int
                
        }
        elevator1 := statusMessage{
                State:     "IDLE",
                CurrentFloor:   4,        
        }

        b, err := json.Marshal(elevator1)
        if err != nil {
                fmt.Println("error:", err)
        }
        os.Stdout.Write(b)
        if len(os.Args) != 2{
                fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
                os.Exit(1)
        }

        service := os.Args[1]
        serverAddress, err := net.ResolveUDPAddr("udp", service)


        fmt.Println("Connecting to server at ", service)

        conn, err := net.DialUDP("udp",nil,serverAddress)

        if err != nil {
                fmt.Println("Could not resolve udp address or connect to it  on " , service)
                fmt.Println(err)
                return
        }

        fmt.Println("Connected to server at ", service)

        defer conn.Close()

        fmt.Println("About to write to connection")

        for {

                time.Sleep(1000*time.Millisecond)
                n, err := conn.Write([]byte(b))
                if err != nil {
                        fmt.Println("error writing data to server", service)
                        fmt.Println(err)
                        return
                }

                if n > 0 {
                        fmt.Println("Wrote ",n, " bytes to server at ", service)
                }
        }

}
