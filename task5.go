package main

import (
	"fmt"

)

func main() {

    var n int

    fmt.Scan(&n)

    input := make([]int32, n, n)

    var otrezok_raz []int

    for i := 0; i < n; i++ {
        fmt.Scan(&input[i])
    }

    for i := 0; i < n; i++ {
        sum := input[i]
        for j := i+1; j < n; j++ {
            sum = sum + input[j]
            if sum == 0 {
//	            fmt.Println("Yes")
                otrezok_raz = append(otrezok_raz, i, j)                

            }        
        } 
    }

    res := 0

    for i := 1; i < len(otrezok_raz);  {
        res = res + n - otrezok_raz[i] 
        i = i + 2        
    }

	fmt.Println(res)	
}


