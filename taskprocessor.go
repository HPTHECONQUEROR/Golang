package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task interface{
	ProcessTask()
}

type BaseTask struct{
	ID int
}

type EmailTask struct{
	BaseTask
	Recipient string
}

func (e EmailTask) ProcessTask(){
	fmt.Printf("Processing Email Task %d for %s\n",e.ID,e.Recipient)
	time.Sleep(time.Second)
}
 

type Worker struct{
	ID int
}

func (w Worker) Start(taskQ <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskQ {
		fmt.Printf("Worker %d started processing task %d\n", w.ID, task.(EmailTask).ID)
		task.ProcessTask()
		fmt.Printf("Worker %d completed task %d\n", w.ID, task.(EmailTask).ID)
	}
}

func main(){
	rand.Seed(time.Now().UnixNano())
	const numWorkers = 3
	const numTasks = 10

	taskQ := make(chan Task, numTasks)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++{
		worker := Worker{ID:1}
		wg.Add(1)
		go worker.Start(taskQ,&wg)
	}



for i := 1; i <= numTasks; i++{
	task := EmailTask{BaseTask: BaseTask{ID:1},Recipient:fmt.Sprintf("mail%d@gmail.com",i)}
	taskQ <- task
}

close(taskQ) 
wg.Wait()        

fmt.Println("All tasks processed!")
}