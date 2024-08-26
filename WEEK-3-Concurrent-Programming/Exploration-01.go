package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

func fetchUserData(url string, ch chan<- []User, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		ch <- nil
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Received non-200 response code")
		ch <- nil
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		ch <- nil
		return
	}

	var users []User
	if err := json.Unmarshal(body, &users); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		ch <- nil
		return
	}

	ch <- users
}

func main() {
	url := "https://fakestoreapi.com/users"
	ch := make(chan []User)
	var wg sync.WaitGroup

	wg.Add(1)
	go fetchUserData(url, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	users := <-ch
	if users == nil {
		fmt.Println("Failed to fetch data")
		return
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("ID %d: Username: %s. Email: %s. First name: %s. Last name: %s.\n",
			user.ID, user.Username, user.Email, user.FirstName, user.LastName)
	}
}
