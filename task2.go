package main

import (
	"fmt"

)

func main() {

var n, m, k, res int

fmt.Scan(&n, &m, &k)

if (n*k % m) == 0 {
res = n*k / m
} else {
res = n*k / m +1
}

	fmt.Println(res)
	
}
