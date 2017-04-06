package main

import (
	"log"

	"github.com/bemobi/gorm"
	_ "github.com/lib/pq"
)

type Task struct {
	ID   int64
	Name string
}

func (Task) TableName() string {
	return "potato"
}

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db = db.Debug()

	task := &Task{
		Name: "Babao",
	}
	if err := db.Create(task).Error; err != nil {
		log.Printf("unable to create task: %s", err)
		return
	}
	log.Printf("Created: %+v ", *task)

	tasks := make([]Task, 0)
	if err := db.Find(&tasks, "id = 1").Error; err != nil {
		log.Printf("unable to create task: %s", err)
		return
	}
	for i := range tasks {
		log.Printf("Task: %+v", tasks[i])
	}
}
