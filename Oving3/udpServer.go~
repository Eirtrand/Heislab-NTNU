package main

import (
"fmt"
"net"

"encoding/json"
)

func main() {

type statusMessage struct {
				State	string
                CurrentFloor	int
                
        }

        port := "129.241.187.136:20006"

        udpAddress, err := net.ResolveUDPAddr("udp",":20006")

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
        var elevMsg statusMessage
        for {

                

                n,address, err := conn.ReadFromUDP(buf[0:])

                if err != nil {
                        fmt.Println("error reading data from connection")
                        fmt.Println(err)
                        return
                }

                if address != nil {

                        fmt.Println("got message from ", address, " with n = ", n)
                        fmt.Println("Raw message:", buf[0:n])
                        if n> 0 {
                        		err := json.Unmarshal(buf[0:n], &elevMsg)
                        		 if err != nil {
                        				fmt.Println("error reading data from connection")
                        				fmt.Println(err)
                        		return
                }
                                fmt.Println(elevMsg.State)
                        		fmt.Printf("%+v",elevMsg)
                        }

                }
        }

}