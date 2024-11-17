package handlers

import (
	"app/models"
	"app/pkg/utils"
	"app/storage/cache"
	"fmt"

	"github.com/google/uuid"
)

func (h *Handler) Registration() {
	var newUser models.User

	uid := uuid.New()
	newUser.UserId = uid.String()

	fmt.Println("Enter firstname: ")
	fmt.Scanln(&newUser.Firstname)

	fmt.Println("Enter lastname: ")
	fmt.Scanln(&newUser.Lastname)

	fmt.Println("Enter age: ")
	fmt.Scanln(&newUser.Age)

	fmt.Println("Enter phone number: ")
	fmt.Scanln(&newUser.PhoneNumber)

	fmt.Println("Enter password: ")
	fmt.Scanln(&newUser.Password)

	otp := utils.GenerateOTP()

	r := cache.NewCache()
	err := r.Set(otp, newUser)
	if err != nil {
		fmt.Println("error on set to cache: ", err)
		return
	}

	fmt.Printf("\n This is your otp code: %d\n\n", otp)

	var createdOTP int
	var phoneNumber string

	fmt.Printf("%s", `		----
	We have sent an otp code create in here
	code: `)

	fmt.Scanln(&createdOTP)
	fmt.Println("Enter your phone number again !")
	fmt.Scanln(&phoneNumber)

	user, err := cache.NewCache().Get(createdOTP, phoneNumber)
	if err != nil {
		fmt.Println("error on getting data from cache", err)
		return
	}

	h.userRepo.CreateUser(*user)

	fmt.Println("Congratulations, you are registered ! ")
}

func (h *Handler) Login() {

	var password string
	var phoneNumber string

	fmt.Println("Enter phone number: ")
	fmt.Scanln(&phoneNumber)

	fmt.Println("Enter password: ")
	fmt.Scanln(&password)

	user, err := h.userRepo.Login(password, phoneNumber)
	if err != nil {
		fmt.Println("error on login")
		return
	}
	UserToken = &Token{}
	UserToken.UserId = user.UserId
	UserToken.Username = user.Firstname
	UserToken.Subscribes = user.Subscribes
}

func (h *Handler) LogOut() {
	UserToken = nil
}

func (h *Handler) GetUsers() {

	users := h.userRepo.GetUsers()

	for _, user := range users {
		fmt.Println("user_id:", user.UserId)
		fmt.Println("firstname:", user.Firstname)
		fmt.Println("last_name:", user.Lastname)
		fmt.Println("age", user.Age)
		fmt.Println("created_at:", user.CreatedAt)
		fmt.Println()

	}

}
