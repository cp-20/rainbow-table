package main

const mode = "createTable"

// const mode = "reverse"

func main() {
	switch mode {
	case "createTable":
		createTable()
	case "reverse":
		reverse()
	}
}
