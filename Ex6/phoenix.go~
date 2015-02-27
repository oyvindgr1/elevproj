package main

import (
	"net"
	."fmt"
	"os/exec"
	"time"
	"encoding/json"
)

func main() {
	isAlive := true
	counter := 0
	addr, _ := net.ResolveUDPAddr("udp", ":20015")
	conn, _ := net.ListenUDP("udp", addr)

	Println("Backup..")
	for (isAlive) {
		conn.SetReadDeadline(time.Now().Add(time.Second*2))
		someData := make([]byte, 1024)
		n, err := conn.Read(someData[0:])
		err3 := json.Unmarshal(someData[0:n], &counter)		
	
		if err3 != nil {
			Println("Error with Unmarshal")
			Println(err3)
		}

		if err != nil {
			isAlive = false
		} else {
			Println("Counting...", n)
			Println(counter)
		}
	}
	conn.Close()

	Println("Master..")	
	spawnNewBackup()
	UDPaddr, _ := net.ResolveUDPAddr("udp", "localhost:20015")
	UDPconn, _ := net.DialUDP("udp", nil, UDPaddr)
	for {
		
 	   	bytes, err1 := json.Marshal(&counter)
		if err1 != nil {
			Println("Error with JSON", err1)
		}		

		_, err := UDPconn.Write(bytes)
		
		if err != nil {
			Println("Error with UDP.Write: ", err)
		}
		
		if counter == 100 {
			break
		}
		counter++
		Println(counter)
		time.Sleep(time.Second)
	}
	UDPconn.Close()
		
}

func spawnNewBackup() {
	Println("Spawn new Backup")
	cmd := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run phoenix.go")
	err := cmd.Run()

	if err != nil {
		Println("Error with spawning backup: ", err)
		return
	}
	
}


