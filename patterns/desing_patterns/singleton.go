package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Singleton is a creational design pattern
that lets you ensure that a class has
only one instance, while providing a global access
point to this instance.
*/

type Database struct {
}

func (Database) CreateSingleConnection() {
	fmt.Println("Creating single connection")
	time.Sleep(2 * time.Second)
	fmt.Println("Connection created")
}

var db *Database
var mux sync.Mutex
var once sync.Once

func getDatabaseInstance2() *Database {
	// sync.Once is a struct that has a method Do that takes a function as an argument and executes it only once.
	once.Do(func() {
		fmt.Println("Creating DB connection")
		db = &Database{}
		db.CreateSingleConnection()
	})
	return db
}

func getDatabaseInstance() *Database {
	mux.Lock()
	defer mux.Unlock()
	if db == nil {
		fmt.Println("Creating database instance")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("Database instance already created")
	}
	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getDatabaseInstance2()
			defer wg.Done()
		}()
	}
	wg.Wait()
}
