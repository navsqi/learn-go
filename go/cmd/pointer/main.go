package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// func bukan pointer
// func ChangeCountryToIndonesia(address Address) {
// 	address.Country = "Indonesia"
// }

// func pointer
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

// pointer di method: method harus selalu menggunakan pointer agar data tidak di copy dan makan banyak memori
type Man struct {
	Name string
}

// penggunaan pointer menggunakan kw *
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	// pass by value
	// Address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// Address2 := Address1
	// Address2.City = "Bandung"

	// fmt.Println(Address1) // tidak berubah
	// fmt.Println(Address2)

	// pass by reference
	var Address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var Address2 *Address = &Address1

	Address2.City = "Bandung"
	fmt.Println(Address1) // ikut berubah
	fmt.Println(Address2)
	fmt.Println("===============================================================")

	// pengunaan asterisk (*)
	// Address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(Address1) // tidak ikut berubah
	// fmt.Println(Address2)

	*Address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(Address1) // ikut berubah
	fmt.Println(Address2)
	fmt.Println("===============================================================")

	// penggunaan new
	// var alamat1 *Address = &Address{}
	var alamat1 *Address = new(Address) // sama dengan yg atas, membuat data kosong dengan kw new
	var alamat2 *Address = alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1) // hasil sama dengan alamat2
	fmt.Println(alamat2)
	fmt.Println("===============================================================")

	// pointer di function
	address3 := Address{}
	// ChangeCountryToIndonesia(address3) // func bukan pointer ==> address3.Country akan tetap kosong
	ChangeCountryToIndonesia(&address3) // func pointer ==> address3.Country = Indonesia

	// function
	fmt.Println(address3)
	fmt.Println("===============================================================")

	// pointer di method
	var nauval Man = Man{"Nauval"}
	nauval.Married() // Jika menggunakan pointer = {Mr. Nauval}, jika tidak menggunakan pointer {Nauval}

	fmt.Println(nauval)
	fmt.Println("===============================================================")

}
