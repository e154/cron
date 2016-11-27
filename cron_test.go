package cron

import (
	"log"
	"time"
	"testing"
)

func TestCore(t *testing.T) {
	ptr := NewCron()

	ptr.NewTask("* * * * * *", func(){log.Println("task1")})
	ptr.NewTask("10,44 * * * * *", func(){log.Println("task2")})
	ptr.NewTask("5 * * * * *", func(){log.Println("task3")})
	//ptr.NewTask("*/5 * * * * *", func(){log.Println("task4")})	// not supported
	ptr.NewTask("1-45,47,50 * * * * *", func(){log.Println("task5")})
	ptr.NewTask("45-59,20 * * * * *", func(){log.Println("task6")})
	ptr.NewTask("0 23 * * * *", func(){log.Println("task7")})
	ptr.NewTask("* 24 16 * * *", func(){log.Println("task8")})
	ptr.NewTask("0 24 16 * * *", func(){log.Println("task9")})
	ptr.NewTask("0 0 21 * * *", func(){log.Println("task10")})

	ptr.Run()

	for {
		time.Sleep(time.Second)
	}
}