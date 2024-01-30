package main

func main() {
	sever := NewAPIServer(":3000")
	sever.Run()
}
