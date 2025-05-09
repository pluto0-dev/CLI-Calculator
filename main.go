package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func calculate(a float64, b float64, op string) float64 {
    switch op {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        if b == 0 {
            fmt.Println("Error: divide by zero")
            return 0
        }
        return a / b
    default:
        fmt.Println("Invalid operator")
        return 0
    }
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter first number: ")
    input1, _ := reader.ReadString('\n')
    num1, _ := strconv.ParseFloat(strings.TrimSpace(input1), 64)

    fmt.Print("Enter operator (+, -, *, /): ")
    operator, _ := reader.ReadString('\n')
    operator = strings.TrimSpace(operator)

    fmt.Print("Enter second number: ")
    input2, _ := reader.ReadString('\n')
    num2, _ := strconv.ParseFloat(strings.TrimSpace(input2), 64)

    // ส่งข้อมูลไปยังฟังก์ชันคำนวณ
    result := calculate(num1, num2, operator)
    fmt.Println("Result:", result)
}
