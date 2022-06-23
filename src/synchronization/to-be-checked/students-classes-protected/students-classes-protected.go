package main

import (
	"fmt"
	"sync"
	"time"
)

// https://www.lurklurk.org/effective-rust/deadlock.html

type Student struct {
	mu      sync.Mutex
	Name    string
	Classes map[string]*Class
}

func NewStudent(name string) *Student {
	return &Student{Name: name, Classes: make(map[string]*Class)}
}
func (s *Student) Enroll(c *Class, wg *sync.WaitGroup) {
	fmt.Printf("1 Enroll Lock student %p\n", s)
	s.mu.Lock()
	fmt.Printf("2 Enroll Locked student %p\n", s)
	s.Classes[c.Name] = c
	fmt.Printf("3 Enroll Unlock student %p\n", s)
	s.mu.Unlock()
	fmt.Printf("4 Enroll Unlocked student %p\n", s)
	c.AddStudent(s)
	wg.Done()
}
func (s *Student) RemoveFromClass(c *Class) {
	fmt.Printf("1 RemoveFromClass Lock student %p\n", s)
	s.mu.Lock()
	fmt.Printf("2 RemoveFromClass Locked student %p\n", s)
	delete(s.Classes, c.Name)
	fmt.Printf("3 RemoveFromClass Unlock student %p\n", s)
	s.mu.Unlock()
	fmt.Printf("4 RemoveFromClass Unlocked student %p\n", s)
}

type Class struct {
	mu        sync.Mutex
	Name      string
	Students  map[string]*Student
	Cancelled bool
}

func NewClass(name string) *Class {
	return &Class{Name: name, Students: make(map[string]*Student)}
}
func (c *Class) Cancel(wg *sync.WaitGroup) {
	fmt.Printf("1 Cancel Lock class %p\n", c)
	c.mu.Lock()
	fmt.Printf("2 Cancel Locked class %p\n", c)
	c.Cancelled = true
	fmt.Printf("3 Cancel Unock class %p\n", c)
	c.mu.Unlock()
	fmt.Printf("4 Cancel Unlocked class %p\n", c)
	for _, s := range c.Students {
		s.RemoveFromClass(c)
	}
	wg.Done()
}
func (c *Class) AddStudent(s *Student) {
	fmt.Printf("1 AddStudent Lock class %p\n", c)
	c.mu.Lock()
	fmt.Printf("2 AddStudent Locked class %p\n", c)
	c.Students[s.Name] = s
	fmt.Printf("3 AddStudent Unock class %p\n", c)
	c.mu.Unlock()
	fmt.Printf("4 AddStudent Unlocked class %p\n", c)
}

func main() {
	s := NewStudent("Andrea")
	c := NewClass("Go")
	counter := 0
	for i := 0; i < 10; i++ {
		counter++
		fmt.Println(">>>>>>>>>> Iteration", counter)
		time.Sleep(1 * time.Millisecond)
		launchEnrollAndCancel(s, c)
	}
}

func launchEnrollAndCancel(s *Student, c *Class) {
	var wg sync.WaitGroup
	wg.Add(2)
	go s.Enroll(c, &wg)
	go c.Cancel(&wg)
	wg.Wait()
}
