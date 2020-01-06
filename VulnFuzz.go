package main

import (
        "bufio"
        "fmt"
        "log"
        "net"
)

//Script for fuzzing the Vulnserver application with go.

func fuzzclient(target, command string) {
        for i := 1; i < 7500; i++ {
                conn, err := net.Dial("tcp", target)
                if err != nil {
                        log.Fatalf("[!] Error at offset %d: %s\n\n", i, err)
                }
                bufio.NewReader(conn).ReadString('\n')
                var buff string
                buff = ""
                for n := 0; n <= i; n++ {
                        buff += "A"
                }
                raw := string(command)
                raw += " %s\n"
                fmt.Fprintf(conn, raw,buff)
                bufio.NewReader(conn).ReadString('\n')

                if err := conn.Close(); err != nil {
                        log.Println("[!] Unable to close connection to %s. Is service alive?\n\n",)
                }
        }
}

func main() {
        commands := []string{"STATS", "RTIME", "LTIME", "SRUN", "TRUN /.:/", "GMON", "GDOG", "KSTET", "GTER", "HTER", "LTER", "KSTAN"}
        target := "127.0.0.1:9999"
        cmdcnt := len(commands)
        for x := 0; x < cmdcnt; x++ {
                cmd := string(commands[x])
                fmt.Printf("[-] Fuzzing %s with command %s\n", target, cmd)
                fuzzclient(target, cmd)
        }
}
