package main

import (
	"log"
	"net"
	"strings"
)

func main() {
	initRedis()
	defer closeRedis()

	server, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		log.Fatalf("[%v] Fail to resolve UDP address\n", err)
	}

	connection, err := net.ListenUDP("udp", server)
	if err != nil {
		log.Fatalf("[%v] Fail to listen UDP to %s\n", err, server)
	}
	log.Println("Listenning...")
	defer connection.Close()

	echo(connection)
}

func echo(connection *net.UDPConn) {
	buffer := make([]byte, 1024)

	for {
		length, client, err := connection.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("[%v] Fail to read from %s\n", err, client.IP.String())
			continue
		}

		request := strings.TrimSpace(string(buffer[:length]))
		log.Printf("[%v] Recive message : %v\n", client.IP.String(), request)
		loggingRedis(client)

		response := []byte(request)
		sent, err := connection.WriteToUDP(append(response, '\n'), client)
		if err != nil {
			log.Fatalf("[%v] Fail to send %v\n", err, client.IP.String())
		}
		log.Printf("[%v] Send %v byte : %s\n", client.IP.String(), sent-1, response)

		clear(buffer)
	}
}
