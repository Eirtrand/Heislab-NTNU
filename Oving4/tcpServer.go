package main

import (
"fmt"
"net"
)


func main() {

        port := "129.241.187.136:33546"


        listener ,err := net.Listen("tcp",":20006")

        if err != nil {
                fmt.Println("error listening on TCP port ", port)
                fmt.Println(err)
                return
        }

        defer listener.Close()

        
        data := make([]byte,1024)

        for {


                //fmt.Println("Waiting for Accept")
                conn, err := listener.Accept()
                _, err = conn.Read(data)

                fmt.Printf("%s\n", data)

                if err != nil {
                        fmt.Println("error reading data from connection")
                        fmt.Println(err)
                        return
                }

                
        }

}