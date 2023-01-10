package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	var null *computer

	if null == nil {
		fmt.Println("null is nil")
	}

	apple := &computer{brand: "apple"}

	newApple := apple

	if newApple == apple {
		fmt.Println("apples are same!")
		fmt.Printf("Mem address of new => %p\nMem address of old => %p\n", &newApple, &apple)
	}

	sony := &computer{brand: "sony"}
	if sony != apple {
		fmt.Println("sony and apple are NOT same!")
		fmt.Printf("Mem address of sony => %p\nMem address of apple => %p\n", &sony, &apple)
	}

	appleVal := *apple
	fmt.Printf("apple                     : %p\n", &apple)
	fmt.Printf("appleVal                  : %p\n", &appleVal)

	if *apple == appleVal {
		fmt.Println("the value that is pointed by the apple and the new variable are same")

		fmt.Printf("the value that is pointed by the apple: 		%s\n", *apple)
		fmt.Printf("new variable: 						%s\n", appleVal)
	}

	change(sony, "hp")
	fmt.Println(*sony)

	fmt.Printf("appleVal                  : %+v\n", valueOf(sony))

	fmt.Printf("dell's address            : %p\n", newComputer("dell"))
	fmt.Printf("lenovo's address          : %p\n", newComputer("lenovo"))
	fmt.Printf("acer's address            : %p\n", newComputer("acer"))
}

func change(cmp *computer, newVal string) {
	cmp.brand = newVal
}

func valueOf(cmp *computer) computer {
	return *cmp
}

func newComputer(brand string) *computer {
	return &computer{brand: brand}
}
