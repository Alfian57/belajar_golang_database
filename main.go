package main

import (
	"belajargolang/entity"
	repository "belajargolang/repository/user"
	"belajargolang/utils"
	"context"
	"fmt"
)

func main() {
	userRepository := repository.NewUserRepository(utils.GetConnection())
	ctx := context.Background()
	var menu string

	for {
		fmt.Println("1. Tampilkan semua users")
		fmt.Println("2. Cari users berdasarkan ID")
		fmt.Println("3. Tambah users")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih Menu [1-4] : ")
		fmt.Scan(&menu)

		if menu == "1" {
			showAllUser(ctx, userRepository)
		} else if menu == "2" {
			findUserById(ctx, userRepository)
		} else if menu == "3" {
			insertUser(ctx, userRepository)
		} else if menu == "4" {
			break
		} else {
			fmt.Println("\n==================")
			fmt.Println("Pilihan Tidak Anda")
			fmt.Print("==================\n\n")
		}
	}

}

func showAllUser(ctx context.Context, userRepository repository.UserRepository) {
	users, err := userRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		printUser(user)
	}
}

func findUserById(ctx context.Context, userRepository repository.UserRepository) {
	var id int32

	fmt.Print("Masukan Id user : ")
	fmt.Scan(&id)

	user, err := userRepository.FindById(ctx, id)
	if err != nil {
		panic(err)
	}

	printUser(user)
}

func insertUser(ctx context.Context, userRepository repository.UserRepository) {
	user := entity.User{}

	fmt.Print("Masukan Username : ")
	fmt.Scan(&user.Username)

	fmt.Print("Masukan Password : ")
	fmt.Scan(&user.Password)

	user, err := userRepository.Insert(ctx, user)
	if err != nil {
		panic(err)
	}

	printUser(user)
}

func printUser(user entity.User) {
	fmt.Println("============================")
	fmt.Println("Username\t :", user.Username)
	fmt.Println("Password\t :", user.Password)
	fmt.Print("============================\n\n")
}
