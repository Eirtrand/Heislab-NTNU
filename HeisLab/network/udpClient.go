package network

import (
"fmt"
"net"
 //"os"
"time"
"encoding/json"
)


func SendUDP(ip string, elevMsg StatusMessage) {

        type StatusMessage struct {
                State     int
                I    int
                
        }


        b, err := json.Marshal(elevMsg)
        if err != nil {
                fmt.Println("error:", err)
        }

        service := ip
        serverAddress, err := net.ResolveUDPAddr("udp", service)


        //fmt.Println("Connecting to server at ", service)

        conn, err := net.DialUDP("udp",nil,serverAddress)

        if err != nil {
                fmt.Println("Could not resolve udp address or connect to it  on " , service)
                fmt.Println(err)
                return
        }

        fmt.Println("Connected to server at ", service)

        defer conn.Close()

        //fmt.Println("About to write to connection")



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
