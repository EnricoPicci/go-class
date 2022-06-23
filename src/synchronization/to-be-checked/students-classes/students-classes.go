package main

import (
	"fmt"
	"sync"
	"time"
)

// https://www.lurklurk.org/effective-rust/deadlock.html

type Student struct {
	Name    string
	Classes map[string]*Class
}

func NewStudent(name string) *Student {
	return &Student{Name: name, Classes: make(map[string]*Class)}
}
func (s *Student) Enroll(c *Class, wg *sync.WaitGroup) {
	s.Classes[c.Name] = c
	c.AddStudent(s)
	wg.Done()
}
func (s *Student) RemoveFromClass(c *Class) {
	delete(s.Classes, c.Name)
}

type Class struct {
	Name      string
	Students  map[string]*Student
	Cancelled bool
}

func NewClass(name string) *Class {
	return &Class{Name: name, Students: make(map[string]*Student)}
}
func (c *Class) Cancel(wg *sync.WaitGroup) {
	c.Cancelled = true
	for _, s := range c.Students {
		s.RemoveFromClass(c)
	}
	wg.Done()
}
func (c *Class) AddStudent(s *Student) {
	c.Students[s.Name] = s
}

func main() {
	s := NewStudent("Andrea")
	c := NewClass("Go")
	counter := 0
	for i := 0; i < 100; i++ {
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
