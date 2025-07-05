package main

import "fmt"


func calculateSum(num1 float32, num2 float32) (float32, float32) {
  return num1 + num2, num1 * num2;
}

func calculateTax(price float32) (stateTax float32, federalTax float32) {
  stateTax = price * 0.05
  federalTax = price * 0.10
  return
}

func main() {
  fmt.Println("Hello, World!")

  sum, _ := calculateSum(3.1, 4.3)
  fmt.Println("Sum of 1.5 and 2.5 is:", sum )

  sTax, fTax := calculateTax(100.00)
  fmt.Println("Calculating taxes for $100.00: ", sTax, fTax)
}
