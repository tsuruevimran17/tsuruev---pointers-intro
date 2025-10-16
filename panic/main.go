package main

import "fmt"

func main(){
	 var p *int
    fmt.Println(p)
}
//panic: runtime error: invalid memory address or nil pointer dereference
//[signal 0xc0000005 code=0x0 addr=0x0 pc=0x7ff65711ed56]

//goroutine 1 [running]:
//main.main()
//	c:/Users/GameStore/Desktop/intucode/go- pointer-intro/2/main.go:7 +0x16