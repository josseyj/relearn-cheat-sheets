# Scala

---

## Overview

* Scala is a JVM language
* Too many language features
* Supports scripting 
    ```bash
    # run the instructions in file
    scala `filename.sc`
    ```
* Object Oriented
* Statically Typed
* Enables DSLs
* @color[red](**_**) is a keyword and has many meanings

---
## Types
    * Scala is a staticly typed language
    * Supports type inference
    * The type is specified after the variable separated by a colon.
        ```scala
        var a : Int
        var names : List[String]
        def square(a : Int) : Int = a * a
        ```
    * @color[red](**Any**) is a catch-all type
    * @color[red](**Unit**) is the type of a statement

---
## Variables

* @color[red](**var**) and @color[red](**val**) are ways to define variables and constants
    ```scala
        var name : String = "Abraham"
        name = "Lincoln" // value can be changed

        val age : Int = 30 // Integer constant
    ```
* Supports type inference
    ```scala
        var name = "Abraham"
        val age = 30 // Integer constant
    ```
---
## Methods

* Method signature and implemention separated by @color[red](**=**)
    ```scala
    def add(a : Int, b : Int) : Int = a + b
    ```
* @color[red](**Unit**) is the return type assumed is @color[green](**=**) is not specified. (There is no void keyword)
    ```scala
        def displayMessage() {
            println("Hello World!")
        }
    ```
+++
### Variable Arguments
* Scala supports variable arguments
    ```scala
    def addAll(first : Int, rest : Int*) : Int = 
        first + rest.reduce(_ + _)
    
    addAll(2, 4, 5, 7) // Result: 18
    ```
* The var-arg argument can be treated as an array
* Passing an array to var-arg
    ```scala
    val nums = Array(1, 2, 3)
    addAll(1, nums:_*) //result 7
    //addAll(1, nums) - won't compile
    ```

+++
### Named Arguments
* Scala supports named arguements
    ```scala
    def divide(divident : Int, divisor : Int) : Int =
        divident / divisor

    divide(6, 2) // Result: 3
    divide(divisor = 2, divident = 6) // Result: 3
    ```

+++ 
### Method Overloading
* Scala supports Method Overloading
### Operator Overloading
* Scala does not support operator overloading
* But scala supports symbolic method names
    ```scala
    def +=(a : Int): Unit = {
        this.age += a
    }
    ```

---
## Control Flows

* @color[red](**if**) - an expression
    ```scala
    val result = if (args.length > 0) args(0) else "default"
    ```

* @color[red](**match**)
    * just like switch in java
    * ```scala
        val result = args(0) match {
            case "A" => 1
            case "B" => 2
            case _   => -1 // catch-all
        }
      ```
+++

* @color[red](**while**) - is a statement with return type `Unit`
    * ```scala
        var i = 0
        while (i < 10) {
            println(s"Square if $i is ${i*i}.")
            i += 1 
            //i++ won't work - ++ is not deined for Int
        }
      ```

+++
* @color[red](**for**)  - is an expression
    * @color[red](**for**) with @color[red](**yield**)
      ```scala 
      val x = for (i <- 1 to 5) yield 2 * i
      //Output: Vector(2, 4, 6, 8, 10)
      ``` 
    * @color[red](**for**) can have multiple @color[aqua](*generators, filters, assignments*)
        ```scala
        for {
            i <- 1 to 5     // generator
            j <- 1 to i     // generator
            k = i * j       // assignment
            if k % 2 == 0   // filter
         } yield k
        //Output: Vector(2, 4, 6, 4, 8, 12, 16, 10, 20)
        ```

---
## Collections
* @color[red](**Array**)
    * values a mutable
    * length is fixed
    * Uses the java array implementation
      ```scala
        val nums = Array(1, 2, 3)
        nums(1) = 4 // Array(1, 4, 3)

        //Insert at tail
        nums :+ 5 // Array(1, 4, 3, 5)

        //Insert at head
        5 +: nums // Array(5, 1, 4, 3)

        //concatenation
        nums ++ nums // Array(1, 4, 3, 1, 4, 3)
      ```

+++
* @color[red](**List**)
    * Immutable
      ```scala
        val nums = List(1, 2, 3) 
        //nums(0) = 4  wont compile beause immutable       

        //concatenation
        nums ::: nums // List(1, 4, 3, 1, 4, 3)
      ```

* @color[red](**Vector**)
    * Scala implementation of Array
      ```scala
        val nums = Vector(1, 2, 3)
      ```
    * There are immutable and mutable versions

+++
* Sequences (@color[red](**Seq**))
    * Ordered collection includes Array, List and Vector
    * ```scala
        Seq(1, 2, 3) // in turn creates a List
      ```
      ```scala
        def addAll(nums : Seq[Int]) : Int = 
            nums.reduce((a: Int, b : Int)  => a +b)

        //all of the following would compile
        addAll(Array(1, 2, 3))
        addAll(List(1, 2, 3))
        addAll(Vector(1, 2, 3))
        addAll(Seq(1, 2, 3))
      ```
+++      
* @color[red](**Range**)
    * Is also a sequence
    * By deault `Range` is inclusive
    * ```scala
        val nums : Seq[Int] = 1 to 5
        //num: Seq[Int] = Range 1 to 5

        //loop and double
        for (i <- 1 to 3) yield 2 * i
        //Output: Vector(2, 4, 6)
      ```

+++
* @color[red](**Set**)
    * unordered collection of unique items
      ```scala
        val nums = Set(1, 2, 3)
        nums ++ Set(2, 3, 4)
        //Output: Set(1, 2, 3, 4)
      ```
    * There are immutable and mutable versions
        * scala.collection.immutable.Set
        * scala.collection.mutable.Set

---
## Error Handling

* @color[red](**try**) ... @color[red](**catch**) ... @color[red](**finally**)
* @color[red](**try**) ... @color[red](**catch**) is an expression, has return type
* @color[red](**finally**) does not affect return type
    ```scala
            def divide(a : Int, b : Int) : Int = {
                try {
                    a / b
                } catch {
                    case e : ArithmeticException => -1
                    case _ : Exception => -2
                } finally {
                    println("Division Complete")
                }
            }

            divide(4, 2) // 2
            divide(2, 0) // -1
    ```

---
## Classes and Objects

* ```scala
    class Person(val name: String, val age : Int = 18)

    val p = new Person("George")
  ```
* All fields are public by default (there is no public key word)

+++
### Inheritance

* TODO

### Interfaces

* TODO

+++

### Companion Objects

* Can access the private and protected members of the class
* Same Name, Same Package, Same file
  ```scala
    class Person private (val name: String, val age : Int)

    object Person {

        def apply(name : String, age : Int) : Person  = new Person(name, age)
    
        def apply(name : String) : Person  = new Person(name, 18)

    }

    val p = Person("George")
  ```
---
## Functional Programming
* Functions are objects
* Supports higher order functions
```scala
def addMethod(a : Int, b : Int) : Int = a + b
val addFunction1 : (Int, Int) => Int = (a, b) => a + b

//placeholder syntax
val addFunction2 : (Int, Int) => Int =  = _ + _
```
+++
### Partially Applied Functions
* Placeholder can be used to create partial functions from other functions
    ```scala
    val multiply : (Int, Int) => Int = _ * _
    val twice = multiply(2, _)
    val thrice = multiply(_, 3)
    ```
+++
### Partial Functions
* Different from Partially Applied Functions.
* @color[red](**PartialFunction[T, R]**) extends from Function[T, R]
* They might not return a result always
* They start with a @color[red](**case**) statement after the @color[green](**{**).
    ```scala
    val squareRoot : (Int) => Double = {
        case x : Int if x > 0 => scala.math.sqrt(x)
    }

    squareRoot(4) //Result: 2
    squareRoot(-4) //Result: MatchError
    ```
* @color[red](**match**) and @color[red](**catch**) uses @color[red](**PartialFunction**)


---
## Unique Language Features
### String Interpolation
```scala
val x = 4

//Resolve variables
println(s"Square of $x is ${ x * x }.")

//Formatting
println(s"Square of ${x}%08.3f is ${ x * x }.")

```
### Code Re-writing

* If the scala code does not compile as it is, the compiler rewrites the code. 
* This enables good support for DSL

+++
* @color[red](**apply**) method
    * If the instance itself is used as a method, the apply method of the instance is used
        ```scala
        Array(1, 2, 3) // Array.apply(1, 2, 3)

        val n = Array(1, 2, 3)
        n(1)    // n.apply(1)
        ```
+++
* In fix method
    * If an instance method has only one argument, the `.` and paranthesis can be skiped.
        ```scala
        "ABCD" charAt 1 // "ABCD".charAt(1) ----> B
        a + b // a.+(b)
        ```

* Right associative method
    * colon (':') ending methods are invoked on the instance on the right.
        ```scala
        Nil :+ 1 :+ 2 // Inserts 1 and then 2
        1 +: 2 +: Nil // Inserts 2 and then 1
        1 :: 2 :: Nil // Inserts 2 and then 1
        ```
