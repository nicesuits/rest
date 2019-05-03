package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("leader"),
		os.Getenv(""),
		os.Getenv("issues"),
	)
	a.Run(":9001")
}
