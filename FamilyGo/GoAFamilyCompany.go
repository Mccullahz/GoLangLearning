package main

// Thank you geeks for geeks, I can never remember how to inherit properly off the top of my head

// declaring a struct
type Comic struct {

	// declaring struct variable
	Universe string
}

// function to return the
// universe of the comic
func (comic Comic) ComicUniverse() string {

	// returns comic universe
	return comic.Universe
}

// declaring a struct
type Marvel struct {

	// anonymous field,
	// this is composition where
	// the struct is embedded
	Comic
}

// declaring a struct
type DC struct {

	// anonymous field
	Comic
}

// main function
func main() {

	// creating an instance
	c1 := Marvel{

		// child struct can directly
		// access base struct variables
		Comic{
			Universe: "MCU",
		},
	}
	// child struct can directly
	// access base struct methods

	// printing base method using child
	println("Universe is:", c1.ComicUniverse())

	c2 := DC{
		Comic{
			Universe: "DC",
		},
	}

	// printing base method using child
	println("Universe is:", c2.ComicUniverse())
}
