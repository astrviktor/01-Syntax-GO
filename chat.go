package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	s "strings"
)

type client chan<- string

var (
	entering  = make(chan client)
	leaving   = make(chan client)
	messages  = make(chan string)
	nicknames = make(map[string]string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func getNickFromString(str string) (nickname string, ok bool) {
	if s.HasPrefix(str, "name:") {
		nickname = s.Replace(str, "name:", "", 1)
		ok = true
		return nickname, ok
	}

	return "", false
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()

	nicknames[who] = who

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch
	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := input.Text()

		name, ok := getNickFromString(text)

		if ok {
			nicknames[who] = name
			text = "<= change name from " + who
		}

		nickname, _ := nicknames[who]
		messages <- nickname + ": " + text
	}
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
