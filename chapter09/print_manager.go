package chapter09

import "fmt"

type PrintManager[T any] struct {
	queue *Queue[T]
}

func (p *PrintManager[T]) QueuePrintJob(job T) {
	p.queue.Enqueue(job)
}

func (p *PrintManager[T]) Run() {
	for !p.queue.Empty() {
		v, _ := p.queue.Dequeue()
		fmt.Println("now printing: ", v)
	}
}
