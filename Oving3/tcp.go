package main

import (
"fmt"
"net"
"os"
"time"
)

func connectToServer() {

        port := "129.241.187.136:33546"

        tcpAddress, err := net.ResolveTCPAddr("ListenTCP",":33546")

        if err != nil {
                fmt.Println("error resolving TCP address on ", port)
                fmt.Println(err)
                return
        }

        listener ,err := net.ListenTCP("tcp",":33546")

        if err != nil {
                fmt.Println("error listening on TCP port ", port)
                fmt.Println(err)
                return
        }

        defer listener.Close()

        var buf [1024]byte
        data := make([]byte,1024)

        for {

                //time.Sleep(1000 * time.Millisecond)
                conn, err := listener.Accept()
                n,address, err := conn.Read(data)

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

func listenToServer(){

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

        for {

                time.Sleep(1000*time.Millisecond)
                n, err := conn.Write([]byte("Connect to: 129.241.187.146:20006\x00"))
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


func main() {
        go connectToServer()
        go listenToServer()


}