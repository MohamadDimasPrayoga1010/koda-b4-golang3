package main

import (
	"fmt"

	"main.go/search"
)

func main() {
	users := []string{"Yoga", "Dimas", "Itsna", "Sidik", "Fiki", "Ari", "Anggi", "Federus"}

	var keyword string
	fmt.Println("Masukan Nama yang ingin dicari")
	fmt.Scan(&keyword)

	userName := search.SearchUser(users, keyword)

	fmt.Println("Hasil Pencarian:", userName)

	// if len(userName) > 0{
	// 	fmt.Println("Nama Ditemukan : ", userName)
	// } else{
	// 	fmt.Println("Nama Tidak Ditemukan")
	// }
}
