# Go

---

## Overview

* **Go** is a language developed by Google.
* Go is a statically typed language
* Example
    ```go
    package main
    import "fmt"
    func main() {
        fmt.Println("Hello World!")
    }
    ```

---

## Types
* Basic types - int, float, string, bytes
* @color[red](rune) for unicode characters
* Types are specified after the variable
---

## Variables
```go
//variable definition
var i int

//initialization
var j int = 0
var name string = "Abraham"

//reference variable
var namePtr *string = &name;

//slice 
var nums = []int{}

//multi value assignment
var a int
var b = 2
i, j = a, b
```
* Uninitialized variables are assigned a zero value
    * 0 for int
    * "" for string
* Unused variables cause compilation errors
* @color[red](**_**) placeholder can be used for unused variables

---
## Methods
```go
func add(a int, b int) int {
    return a + b
}
```
* Supports multiple return types
```go
func GetAccount(accountId string) (Account, error) {
    ...
}
```

+++
### Variable Arguments
```go
func addAll(a int, b ...int) int {
    result := a
    for _, i := range b {
        result += i
    }
    return result
}
```

---

## Control Flows

* if... else if... else
* if with assignment
    ```go
    if age := account.age(); age % 2 ==0 {
        fmt.Println("Age is even.")
    }
    ```
* for
    ```go
    for i := 0; i < 10; i++ {
		fmt.Printf("Square of %d is %d.\n", i, i*i)
    }
    ```
* for ... range
    ```go
    arr := []int{1, 3, 6}
	for _, i := range arr {
		fmt.Printf("Square of %d is %d.\n", i, i*i)
	}

    ```
* infinite loop
    ```go
    for {
        ...
        break
        ...
    }
    ```

---
## Collections
```go
//array - fixed length
var nums = [5]int{} // [0,0,0,0,0]

//slice - variable length
var nums = []int{} // []
nums = append(nums, 1) // [1]

//maps
var nameToAgeMap = map[string]int{
    "Abraham": 20,
    "George":  30,
}
for key, value := range nameToAgeMap {
    fmt.Printf("%s is %d years old.\n", key, value)
}
```

---

## Error Handling


---

## Classes and Objects

+++
### Inheritance

+++
### Interfaces
---

## Language Features

