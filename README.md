# go-fizzbuzz
:rabbit: Learn to write testable fizzbuzz program in Golang

## The program
**FizzBuzz** is a toy program that has very simple rules: You give **FizzBuzz** a number and it will:

- Return **Fizz** if that number is a multiples of **3**
- Return **Buzz** if that number is a multiples of **5**
- Return **FizzBuzz** if that number is both multiples of **3** and **5**
- Otherwise, return **the number itself** 

We will implement it as `FizzBuzz()` function in `main.go` and test it in `main_test.go()`

## Unit testing - Meet the tools

Go has a bunch of awesome built-in features, one of them is `go test`.

### Write the tests
To write test for your code:

- Create a new `*_test.go` file, where `*` is the source file you want to test. In this example, it will be `main_test.go`.
- Write the test function with the name of `Test*(t *testing.T) { }`, where `*` is whatever meaningful describe you want.

### Run the tests
To test them, open Terminal and run:

```bash
$ go test
```

If you are using editors like `vim`, install `vim-go` plugin, then type: `:GoTest`. For `Sublime`, install the [official build system integration](https://github.com/golang/sublime-build), then run `Test` command in the `Command Pallete`.

### Code Coverage
It's not a must, but it's recommeneded to run coverage to see how many percent of your code are not being tested yet.

To run coverage, open Terminal and run:

```bash
$ go test -cover
```

The output will tell you how much coverage your code is. For example:

```
PASS
coverage: 81.8% of statements
ok      github.com/huytd/go-fizzbuzz    0.421s
```

This output mean: the project is 81.8% coverage.

If this is not enough, Go also provide you a deeper look into your code, so you can see what lines of code are not tested. To do it, you need to tell Go to generate some output when it run coverage:

```bash
$ go test -coverprofile=cover.out
```

Then you can use the `cover` tool to read this output in HTML format:

```bash
$ go tool cover -html=cover.out
```

If you are new to unit test, reading the coverage report also let you know what need to test in your code.

## TDD approach - Write test first, the code can wait

## Write code first, test it later

## Changes are OK, and you can handle it

## Table Driven Test

## How much coverage is enough?
