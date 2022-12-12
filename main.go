package main

import "os"

func main() {
	res := os.Getenv("MY_NAME")
	print("Key : ", res   )
}
