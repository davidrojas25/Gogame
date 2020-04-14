package main

import "fmt"

func getNames() (iFirstName string, iLastName string) {

	fmt.Println("Enter your first name:")
	fmt.Scanf("%s", &iFirstName)
	fmt.Println("Enter your last name:")
	fmt.Scanf("%s", &iLastName)

	return iFirstName, iLastName
}

func selectWeapon(coins int) (iweapon string, icoins int) {

	fmt.Println("Welcome to the Weapons Depot")
	fmt.Println("Select from the following list:")
	fmt.Println("Hammer - $200")
	fmt.Println("Sword - $400")
	fmt.Println("Gun - $750")
	fmt.Println("Rifle - $1000")
	fmt.Println("Enter Choice of Weapon(type name only):")
	fmt.Scanf("%s", &iweapon)
	var checkout int
	switch iweapon {
	case "Hammer":
		checkout = 200
	case "Sword":
		checkout = 400
	case "Gun":
		checkout = 750
	case "Rifle":
		checkout = 1000
	}
	if coins >= checkout {
		fmt.Println("You have sucessfully purchased a", iweapon)
		icoins = coins - checkout

	}
	return iweapon, icoins
}

func main() {

	var firstName, lastName = getNames()
	fmt.Println(lastName)
	fmt.Println("Welcome ", firstName, "", lastName)
	coins := 500
	weapon := ""
	weapon, coins = selectWeapon(coins)
	fmt.Println("Ready to go to battle with your ", weapon)
	fmt.Println("You now have only", coins, "left")
}
