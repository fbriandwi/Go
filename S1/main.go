package main
import "fmt"
func main() {
	// var name string = "Febrian"
	// var age int = 23

	// fmt.Println("namanya adalah", name)
	// fmt.Println("usianya adalah", age)

/* sama aja */

	// var name string
	// var age int = 17

	// name = "febrian"
	// age = 23

	// fmt.Println("namanya adalah", name)
	// fmt.Println("umur adalah", age)

	// var name = "pep"
	// var age = 23
	// var address = "larantuka"

	// fmt.Printf("halo namaku %s, umurku %d, aku tinggal di %s", name , age, address)

	/* float */
	// var decimalNumber float32 = 3.63
	// fmt.Printf("number decimal %f\n", decimalNumber)
	// fmt.Printf("number decimal %.3f\n", decimalNumber)

	/* boolean */
	// var condition bool = true
	// fmt.Printf("is it pemitted? %t\n", condition)

	// var message = `selamat datang di hacktiv8 "peserta".`
	// fmt.Println(message)
	
	// const full_name string = "pep"
	// fmt.Printf("nama saya %s", full_name)

	/* Iota */
	const (
		c1 = iota 
		c2 = iota
		c3 = iota
	) 
	fmt.Println(c1, c2 ,c3)
}
