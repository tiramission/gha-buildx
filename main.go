package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("Args: %+v\n", os.Args)
	runt()
}

func runt() {
	// 基本架构信息
	fmt.Printf("操作系统: %s\n", runtime.GOOS)
	fmt.Printf("处理器架构: %s\n", runtime.GOARCH)
	fmt.Printf("编译器版本: %s\n", runtime.Version())
	fmt.Printf("GOROOT: %s\n", runtime.GOROOT())

	// CPU 核心数
	fmt.Printf("逻辑CPU数量: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS设置: %d\n", runtime.GOMAXPROCS(0))

	// Goroutine 信息
	fmt.Printf("当前Goroutine数量: %d\n", runtime.NumGoroutine())
}
