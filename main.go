package main

import (
	// "crypto/rand"
	"fmt"
	"math/rand/v2"
	"strconv"
)

func main() {
	var number string
    var count int = 0
	n := rand.IntN(100)
	for {
		fmt.Print("Введите число: ")
		fmt.Scan(&number)

        number, err := strconv.Atoi(number)

        if err != nil{
            fmt.Printf("Вы ввели не число.\n")
            
        } else if number >= n {
			fmt.Printf("Число меньше чем %d!\n", number)
            count++
		
        } else if number <= n {
			fmt.Printf("Число больше чем %d!\n", number)
            count++

        } else if number == n {
            count++
			fmt.Printf("Вы нашли число! оно было %d\n", n)
            fmt.Printf("Количество попыток: %d\n", count)
			break
		
        }
	}
}
