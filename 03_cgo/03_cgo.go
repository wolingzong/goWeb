package main

//#include <stdio.h>
import "C"

func main() {
	C.puts(C.CString("你好, GopherChina 2018!\n"))
}
