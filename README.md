# Simple Calculator

A simple calculator made with Go (golang).

- Mathematical operators only have plus and minus.  
- Use **Space** among of values *ex : 1 + 1 + 3 -1* 


## Running

```bash
go run main.go
```

## Usage

```golang
x := "5+2+1"
```
**just change value of x*


## Unit Test
```bash
go test main.go
```

## Function Component

- **Tokenizing** is for split **Space** and remove empty array after split
- **isOperand** is for check value is operand or number
- **value** is for Parsing string to number
- **Evaluate** is for calculate or is main function..