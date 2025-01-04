package main

import (


        "fmt"

        "os"

)

func main() {
        file, err := os.Open("my_file.txt")
        if err != nil {
                fmt.Printf("Error opening file: %v\n", err)
                os.Exit(1)
        }
        defer file.Close()

        //Further processing
        //data, err := ioutil.ReadAll(file)
        //if err != nil {
        //   fmt.Printf("Error reading file: %v\n", err)
        //   os.Exit(1)
        //}        
        //fmt.Printf("%s\n", data)
}