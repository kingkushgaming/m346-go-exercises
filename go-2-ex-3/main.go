package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Modul 104",
		117: "Modul 117",
		346: "Modul 346",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117)
	modules[104] = "New Modul 104"
	modules[200] = "New Modul 200"
	fmt.Println(modules)
}
