package main

import (
	"fmt"
	"os"
	"main.go/search"
)

func main() {
	// defer fmt.Println("Program Telah Di Jalankan")
	defer func(){
		if r := recover(); r != nil{
			fmt.Println("Terjadi Kesalahan", r)
			os.Exit(1)
		} else{
			fmt.Println("Program Berjalan")
		}
	}()
	users := []string{"Yoga", "Dimas", "Itsna", "Sidik", "Fiki", "Ari", "Anggi", "Federus"}

	var keyword string
	fmt.Println("Masukan Nama yang ingin dicari")
	fmt.Scan(&keyword)
	
	
	userName := search.SearchUser(users, &keyword)

	if len(userName) == 0 {
		fmt.Println("Hasil Pencarian:", userName)
		panic("Nama Tidak Ditemukan")
	}
	fmt.Println("Hasil Pencarian:", userName)


	// if len(userName) > 0{
	// 	fmt.Println("Nama Ditemukan : ", userName)
	// } else{
	// 	fmt.Println("Nama Tidak Ditemukan")
	// }
}
