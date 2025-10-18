package main
/* На вход подается целое число, сумма денег, которые у вас есть. Ваша задача - вывести марку телефона, которую вы
можете себе позволить купить.
Если сумма больше 1000 - вывести Apple
Если сумма от 500 до 1000 (включительно) - вывести Samsung
Если сумма меньше 500 - вывести Nokia с фонариком
*/

import "fmt"

func main() {
     var money int
    fmt.Scan(&money)

    if money > 1000 {
        fmt.Println("Apple")
    } else if money >= 500 {
        fmt.Println("Samsung")
    } else {
        fmt.Println("Nokia с фонариком")
    }
}