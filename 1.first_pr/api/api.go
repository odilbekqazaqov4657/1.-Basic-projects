package api

import (
	"app/api/handlers"
	"app/services"
	"app/storage/db"
	"fmt"
)

func Api(userRepo db.UserRepo, service services.Service) {
	var cmd int
	h := handlers.NewHandler(userRepo, service)

	for cmd != -1 {

		fmt.Println(`
	1. Registration
	2. Log-in
	3. Log-out
	4. Get_users
	5. Services/Convertor/BinToDec
	6. Service/Todo/Create
	7. Service/Todo/GetAll
	-1. Exit
	`)
		fmt.Println("choose a command: ")
		fmt.Scanln(&cmd)

		switch cmd {
		case 1:
			h.Registration()
		case 2:
			h.Login()
		case 3:
			h.LogOut()
		case 4:
			h.GetUsers()
		case 5:
			h.BinToDec()
		case 6:
			h.CreateTodo()
		case 7:
			h.GetTodos()
		}
	}

	fmt.Println("Xizmatdan foydalanganingiz uchun raxmat !")
}
