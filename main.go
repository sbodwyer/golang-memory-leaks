package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/myLovelyFunction1", myLovelyHandler1)
    http.HandleFunc("/myLovelyFunction2", myLovelyHandler2)
    http.HandleFunc("/myLovelyFunction3", myLovelyHandler3)

    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", nil)
}

// myLovelyHandler1 handles the /myLovelyFunction1 endpoint
func myLovelyHandler1(w http.ResponseWriter, r *http.Request) {
    myLovelyFunction1()
    fmt.Fprintf(w, "myLovelyFunction1 endpoint called")
}

func myLovelyFunction1() {
    myLovelySlice := make([]int, 0)
    for i := 0; i < 1000000; i++ {
        myLovelySlice = append(myLovelySlice, i)
    }
}

// myLovelyHandler2 handles the /myLovelyFunction2 endpoint
func myLovelyHandler2(w http.ResponseWriter, r *http.Request) {
    myLovelyFunction2()
    fmt.Fprintf(w, "myLovelyFunction2 endpoint called")
}


func myLovelyFunction2() {
    go func() {
        for {
            time.Sleep(1 * time.Second)
        }
    }()
}

// myLovelyHandler3 handles the /myLovelyFunction3 endpoint
func myLovelyHandler3(w http.ResponseWriter, r *http.Request) {
    myLovelyFunction3()
    fmt.Fprintf(w, "myLovelyFunction3 endpoint called")
}

func myLovelyFunction3() {
    myLovelyMap := make(map[int]*int)
    for i := 0; i < 1000000; i++ {
        value := i
        myLovelyMap[i] = &value
    }
}