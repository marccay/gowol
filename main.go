package main

import (
	"encoding/hex"
	"os"
	"fmt"
	"bufio"
	"strings"
	"net"
)

var port string
var ip string
var macAddr string

func main() {
	args := os.Args[1:]
	for i, x := range args {
		switch x {
			case "-p":
				port = args[i+1]
			case "-ip":
				ip = args[i+1]
			case "-mac":
				macAddr = args[i+1]
			default:
				continue
		}
	}

	if port == "" {
		port = "7"
	}
	if ip == "" || ip == "all"{
		ip = "224.0.0.1"
	}
	if macAddr == "" {
		fmt.Printf("Enter macaddress of machine: \t")
		scan := bufio.NewReader(os.Stdin)
		mac, err := scan.ReadString('\n')
		if err != nil {
			panic(err)
		}
		macAddr = mac[:len(mac)-1]
	}

	wakeLan(macAddr, ip, port)
}


func createMagic(macAddr string) []byte {
	macBytes, err := hex.DecodeString(strings.Join(strings.Split(macAddr, ":"), ""))
	if err != nil {
		panic(err)
	}
	b := []uint8{255,255,255,255,255,255}
	for i := 0; i < 16; i++ {
		b = append(b, macBytes...)
	}
	fmt.Println(b)
	return b
}


func wakeLan(macAddr string, ip string, port string) {
	combined := ip + ":" + port
	conn, err := net.Dial("udp", combined)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	packet := createMagic(macAddr)
	_, err = conn.Write(packet)
	if err != nil {
		panic(err)
	}
}
