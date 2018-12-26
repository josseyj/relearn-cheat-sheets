# Scala

---

## Overview

* Scala is a JVM language
* Too many language features
* Big learning curve
* Supports scripting 
    * scala `filename.sc` - run the instructions in file
* Object Oriented
* Statically Typed
* Enables DSLs

---

## Variable Definitions

* ```var name : String = "Abraham"```
* ```val age : Int = 30``` - Integer constant
* Supports type inference

## Method Deinitions

* ```def add(a : Int, b : Int) : Int = a + b```
* `Unit` is the return type assumed is `=` is not specified. (There is no void keyword)
    * ```scala
        def displayMessage() {
            println("Hello World!")
        }
    ```
---

## Control Flows

* `if` - an expression
    * `val result = if (args.length > 0) args(0) else "default"`

* `match`
    * just like switch in java
    * ```scala
        val result = args(0) match {
            case "A" => 1
            case "B" => 2
            case _   => -1 // catch-all
        }
      ```
* `while` - is a statement with return type `Unit`
    * ```scala
        var i = 0
        while (i < 10) {
            println(s"Square if $i is ${i*i}.")
            i += 1 //i++ won't work - ++ is not deined for Int
        }
      ```
* for - is an expression
    * ```val x = for (i <- 1 to 5) yield 2 * i``` --> `Vector(2, 4, 6, 8, 10)`


---

## Classes and Objects

* ```scala
    class Person(val name: String, val age : Int = 18)

    val p = new Person("George")
  ```
* All fields are public by default (there is no public key word)

+++

### Companion Objects

* Can access the private and protected members of the class

* Same Name, Same Package, Same file

* ```scala
    class Person private (val name: String, val age : Int)

    object Person {

        def apply(name : String, age : Int) : Person  = new Person(name, age)
    
        def apply(name : String) : Person  = new Person(name, 18)

    }

    val p = Person("George")
  ```
