package main

import "fmt"

// QueueElem - тип элемента очереди
type QueueElem string

// Queue - структура очереди
type Queue struct {
	name   string
	head   int
	tail   int
	length int
	body   map[int]QueueElem
}

// Init - Инициализация очереди
func (q *Queue) Init(name string) {
	q.name = name
	q.head = 0
	q.tail = 0
	q.length = 0
	q.body = make(map[int]QueueElem)
}

// Push - Добавить элемент в очередь
func (q *Queue) Push(elem QueueElem) {
	q.body[q.tail] = elem
	q.tail++
	q.length++
}

// Pop - Забрать элемент из очереди
func (q *Queue) Pop() (elem QueueElem) {
	if q.length == 0 {
		fmt.Print("Ошибка: очередь пуста, элементов нет")
		return
	}

	elem = q.body[q.head]
	delete(q.body, q.head)
	q.head++
	q.length--
	return elem
}

// Info - Информация про очередь
func (q *Queue) Info() {
	fmt.Println("Queue name:", q.name)
	fmt.Println("  head:", q.head)
	fmt.Println("  tail:", q.tail)
	fmt.Println("  length:", q.length)
	fmt.Println("  elems:", q.body)
}

// main - точка входа
func main() {

	var queue Queue

	queue.Init("main queue")
	queue.Info()

	queue.Push("11111")
	queue.Push("22222")
	queue.Push("33333")
	queue.Info()

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

	queue.Info()
}
