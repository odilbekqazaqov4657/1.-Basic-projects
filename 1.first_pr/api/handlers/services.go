package handlers

import "fmt"

func (h *Handler) BinToDec() {

	if UserToken != nil {
		if UserToken.Subscribes.Convertor {

			var binNum int

			fmt.Printf("%s", "Enter bin number: ")
			fmt.Scanln(&binNum)

			decNum := h.Service.Convertor.BinToDec(binNum)

			fmt.Println("result: ", decNum)

		} else {
			var cmd int

			fmt.Println("you are not subscribe convertor!")

			fmt.Print(`
				1. Subscribe
				2. Unsubscribe
			`)
			fmt.Scanln(&cmd)
			if cmd == 1 {
				UserToken.Subscribes.Convertor = true
				fmt.Println("you are subscribed convertor!")

			} else if cmd == 2 {
				UserToken.Subscribes.Convertor = false
			}
		}
	} else {
		fmt.Println("you are not registered !")
	}
}
