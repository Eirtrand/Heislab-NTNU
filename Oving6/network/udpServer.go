package network

import (
"fmt"
"net"
"time"
"encoding/json"
)

type StatusMessage struct {
        State   int
        I       int           
}

func ReadUDP(ip string) (elevMsg StatusMessage ) {



        port := ip
        elevMsg.State = 0
        udpAddress, err := net.ResolveUDPAddr("udp",":20012")
        if err != nil {
                fmt.Println("error resolving UDP address on ", port)
                fmt.Println(err)
                elevMsg.State = 0
                return
        }
        conn ,err := net.ListenUDP("udp",udpAddress)

        if err != nil {
                fmt.Println("error listening on UDP port ", port)
                fmt.Println(err)
                elevMsg.State = 0
                return
        }
        defer conn.Close()

        var buf [1024]byte


                
                conn.SetReadDeadline(time.Now().Add(2*time.Second))
                n,address, err := conn.ReadFromUDP(buf[0:])
                if err != nil {
                        fmt.Println("error reading data from connection asd")
                        elevMsg.State = 0
                        return //elevMsg
                }
                if address != nil {

                        fmt.Println("got message from ", address, " with n = ", n)
                        //fmt.Println("Raw message:", buf[0:n])
                        if n> 0 {
                        		err := json.Unmarshal(buf[0:n], &elevMsg)
                        		if err != nil {
                        			fmt.Println("error reading data from connection lol")
                        			fmt.Println(err)
                                                elevMsg.State = 0
                        		return
                                        }
                                elevMsg.State = 1
                                //fmt.Println(elevMsg.I)
                        	//fmt.Printf("%+v",elevMsg)
                        }else{
                                elevMsg.State = 0
                                fmt.Println("Set state to 0")
                        }
                }
                return elevMsg
        

}
