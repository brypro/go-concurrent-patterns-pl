package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func main() {
	var wg sync.WaitGroup
	var mux sync.RWMutex
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &mux)
	}
	wg.Wait()
	fmt.Println(Balance(&mux))
}

// Deposit es una funcion que deposita dinero
func Deposit(amount int, wg *sync.WaitGroup, mux *sync.RWMutex) {
	mux.Lock() //bloquea 1 a la vez
	b := balance
	balance = b + amount
	mux.Unlock() //desbloquea
	defer wg.Done()
}

// Balance es una funcion que devuelve el balance
func Balance(mux *sync.RWMutex) int {
	mux.RLock()
	b := balance
	mux.RUnlock()
	return b
}

/*
Lock bloquea lecturas (con RLock) y escrituras (con Lock) de otras goroutines
Unlock permite nuevas lecturas (con Rlock) y/o otra escritura (con Lock)
RLock bloquea escrituras (Lock) pero no bloquea lecturas (RLock)
RUnlock permite nuevas escrituras (y también lecturas, pero por la naturaleza de RLock, estas no se vieron bloqueadas nunca)
En esencia, RLock de RWLock garantiza una secuencia de lecturas en donde el valor que lees no se verá alterado por nuevos escritores, a diferencia de no usar nada.
https://stackoverflow.com/questions/19148809/how-to-use-rwmutex
*/
