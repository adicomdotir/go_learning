package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	fmt.Printf("sum is %d\n", sum([]int{1, 2, 3, 4, 5}))
	person := getPersons()[0]
	person.changeName("Jane", 25)
	person.sayHello()

	selectSample()
}

func sum(a []int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

// create struct for Person
type Person struct {
	name string
	age  int
}

// add method for person
func (p Person) sayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

// method for change person values
func (p *Person) changeName(name string, age int) {
	p.name = name
	p.age = age
}

// api for get persons
// convert below code to api
func getPersons() []Person {
	return []Person{
		{"John", 30},
		{"Jane", 25},
	}
}

// waitgroup sample
func waitGroupSample() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine")
	}()
	wg.Wait()
}

// channel sample
func channelSample() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		c1 <- "hello"
	}()
	go func() {
		c2 <- "world"
	}()
}

// select sample
func selectSample() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		c1 <- "hello"
	}()
	go func() {
		c2 <- "world"
	}()
	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case msg2 := <-c2:
		fmt.Println(msg2)
	}
}

// goroutine sample
func goroutineSample() {
	go func() {
		fmt.Println("goroutine")
	}()
}
