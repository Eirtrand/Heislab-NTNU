package main

import (
"fmt"
"net"
///"time"
)

func main() {

        port := "129.241.187.136:30000"

        udpAddress, err := net.ResolveUDPAddr("udp4",":30000")

        if err != nil {
                fmt.Println("error resolving UDP address on ", port)
                fmt.Println(err)
                return
        }

        conn ,err := net.ListenUDP("udp",udpAddress)

        if err != nil {
                fmt.Println("error listening on UDP port ", port)
                fmt.Println(err)
                return
        }

        defer conn.Close()

        var buf [1024]byte

        for {

                //time.Sleep(1000 * time.Millisecond)

                n,address, err := conn.ReadFromUDP(buf[0:])

                if err != nil {
                        fmt.Println("error reading data from connection")
                        fmt.Println(err)
                        return
                }

                if address != nil {

                        fmt.Println("got message from ", address, " with n = ", n)

                        if n> 0 {
                                fmt.Println("from address", address, "got message:", string(buf[0:n]), n)
                        }
                }
        }

}