package main

import "gomall/common/logs"

func main() {
	logs.LogInit("test")

	logs.Error("hello world")
}
