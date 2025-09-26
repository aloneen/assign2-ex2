package main

import "github.com/aloneen/assign2-ex2/initializers"

func init() {
    initializers.LoadEnvVariable()
    initializers.ConnectToDatabase()
}

func main() {
	
}
