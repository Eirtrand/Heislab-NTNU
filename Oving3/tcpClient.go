package main

import (
"fmt"
"net"
"os"
"time"
)

func main() {

        if len(os.Args) != 2{
                fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
                os.Exit(1)
        }

        service := os.Args[1]
        
        serverAddress, err := net.ResolveTCPAddr("tcp", service)


        fmt.Println("Connecting to server at ", service)

        conn, err := net.DialTCP("tcp",nil,serverAddress)

        if err != nil {
                fmt.Println("Could not resolve tcp address or connect to it  on " , service)
                fmt.Println(err)
                return
        }

        fmt.Println("Connected to server at ", service)

        //defer conn.Close()

        fmt.Println("About to write to connection")

        

                time.Sleep(1000*time.Millisecond)
                n, err := conn.Write([]byte("Connect to: 129.241.187.146:20006 \x00"))
                if err != nil {
                        fmt.Println("error writing data to server", service)
                        fmt.Println(err)
                        return
                }

                if n > 0 {
                        fmt.Println("Wrote ",n, " bytes to server at ", service)
                }
        

}