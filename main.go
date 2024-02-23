package main

import (
	"fmt"
	"os"

	"github.com/papaya147/clippy/store"
	"github.com/papaya147/clippy/util"
)

func main() {
	config := util.DefaultConfig()
	store := store.NewStore(config.StoreFile)

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no command provided")
		return
	}

	if len(args) == 1 && args[0] == "help" {
		// TODO - add help manual
		return
	}

	if len(args) >= 2 {
		switch args[0] {
		case "set":
			if len(args) != 3 {
				fmt.Println("invalid arguments")
				return
			}
			if err := store.Set(args[1], args[2]); err != nil {
				fmt.Println(err)
			}
		case "get":
			if len(args) != 2 {
				fmt.Println("invalid arguments")
				return
			}
			value, err := store.Get(args[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		default:
			fmt.Println("oops, this command is not supported")
		}
	}
}
