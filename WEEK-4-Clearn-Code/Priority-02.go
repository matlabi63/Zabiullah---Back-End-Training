package main

import "fmt"


type Set struct {
    elements map[int]struct{}
}


func NewSet() *Set {
    return &Set{elements: make(map[int]struct{})}
}


func (s *Set) Add(value int) {
    s.elements[value] = struct{}{}
}


func (s *Set) Get() []int {
    var result []int
    for key := range s.elements {
        result = append(result, key)
    }
    return result
}


func (s *Set) Remove(value int) {
    delete(s.elements, value)
}

func main() {
    s := NewSet()

    s.Add(1)
    s.Add(2)
    s.Add(1) 
    s.Add(3)
    s.Add(4)
    s.Add(5)

    s.Remove(100)
    fmt.Println(s.Get())
}
