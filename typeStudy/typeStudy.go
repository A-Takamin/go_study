package typeStudy

import "fmt"

// ============================ //

type ID int
type Priority int

func ProcessTask(id ID, p Priority) {
	// id, p intでは、引数の順番を指定できないのでバグの恐れ。
	// typeで型を作ってあげると安心
}

func Study1() {
	var id ID = 3
	var priority Priority = 1
	ProcessTask(id, priority)
}

// ============================ //

type Task struct {
	ID     int
	Detail string
	done   bool
}

func Study2() {
	var task Task = Task{
		ID:     1,
		Detail: "buy the milk",
		done:   true,
	}
	fmt.Println(task.ID)
	fmt.Println(task.Detail)
	fmt.Println(task.done)

	task2 := Task{1, "buy the pen", true}
	fmt.Println(task2)
}
