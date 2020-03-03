package main

import (
	"fmt"
	"net/http"
	"time"
)

func mirroredQuery() string {
	responses := make(chan string, 3)

	go func() {
		responses <- requestTimer("https://golang.org/")
	}()
	go func() {
		responses <- requestTimer("https://habr.com/ru/")
	}()
	go func() {
		responses <- requestTimer("https://ya.ru/")
	}()

	return <-responses // возврат самого быстрого ответа
}

func requestTimer(hostname string) (res string) {
	t0 := time.Now()

	res = request(hostname)

	t1 := time.Now()
	time := fmt.Sprintf("Затраченное время: %v", t1.Sub(t0))

	return res + time
}

func request(hostname string) (res string) {
	resp, err := http.Get(hostname)

	if err != nil {
		res = fmt.Sprintf("%s : Ошибка соединения. %s\n", hostname, err)
		return res
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		res = fmt.Sprintf("%s : Ошибка. http-статус: %d\n", hostname, resp.StatusCode)
		return res
	}

	res = fmt.Sprintf("%s : Онлайн. http-статус: %d\n", hostname, resp.StatusCode)
	return res
}

func main() {

	fmt.Println("Самый быстрый ответ: \n" + mirroredQuery())

}
