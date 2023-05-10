package main

import (
	"fmt"
    "strings"
)

func main() {

    var n, k, i, len_s int32
    var s string
    flag := false
    var res int32 = -1
    fmt.Scan(&n)
    fmt.Scan(&s)
    len_s = int32(len(s))
    for k = 4; k <= n && !flag; k++ {
        for i = 0; i < len_s - k +1; i++{
//	        fmt.Println(s[i:i+k])
            if strings.Contains(s[i:i+k], "a") && strings.Contains(s[i:i+k],"b") && strings.Contains(s[i:i+k], "c") && strings.Contains(s[i:i+k], "d") {
//                fmt.Println("Yes")
                res = k
                flag = true
                break
            }
        }
    }

	fmt.Println(res)
	
}
