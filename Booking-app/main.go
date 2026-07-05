package main

import (
	"fmt"
	"strings"
)

func main() {
	var orgname = "NOTA"
	var confname = "Go Conference"
	const confTickets = 50
	var remainTickets uint = 50

	fmt.Printf(" Welcome to %v\n", orgname)
	greetUsers(confname)
	// fmt.Println(" Welcome to", confname, "booking application")
	fmt.Println(" We have total of", confTickets, "tickets and", remainTickets, "this many are still available")
	fmt.Println(" Book your tickets to attend")
	fmt.Printf(" orgname is %T , remaining tickets is %T \n", orgname, remainTickets)

	// var bookings =[50]string{} //arrays
	bookingslice := []string{} // using slices

	for remainTickets > 0 && len(bookingslice) < confTickets {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for input
		fmt.Println(" Enter your first name : ")
		fmt.Scan(&firstName)

		fmt.Println(" Enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Println(" Enter your email address : ")
		fmt.Scan(&email)

		fmt.Println(" Enter how many tickets are you buying")
		fmt.Scan(&userTickets)

		// check if user input is valid
		isValidName := len(firstName) >= 2 && len(lastName) >= 2

		// EMAIL VERIFICATION
		isValidEmail := strings.Contains(email, "@")

		isValidTicketNum := userTickets > 0 && userTickets <= remainTickets

		//this would check the validity first before we run the booking
		if isValidName && isValidEmail && isValidTicketNum {

			remainTickets = remainTickets - userTickets

			// bookings[0] = firstName + " , " + lastName
			bookingslice = append(bookingslice, firstName+" , "+lastName) // adds inside of it

			//It has the same output method just different input method

			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("Array type: %T\n", bookings)
			// fmt.Printf("The length array: %v\n", len(bookings))

			// fmt.Printf("The whole slice: %v\n", bookingslice)
			// fmt.Printf("Slice type: %T\n", bookingslice)
			// fmt.Printf("The first slice value is %v\n", bookingslice[0])
			// fmt.Printf("The length slice: %v\n", len(bookingslice))

			fmt.Println("Thank you for booking")
			fmt.Printf(" User %v %v bought %v tickets successfully, it has been sent to email address:%v\n", firstName, lastName, userTickets, email)

			//we want to only get first names from the bookingslice
			var firstNames = []string{}
			for _, bookingloop := range bookingslice {
				var names = strings.Split(bookingloop, " , ")
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf(" Hey %v, there are now  %v tickets left for  %v\n", firstNames, remainTickets, confname)

			// fmt.Println(remainTickets)
			// fmt.Println(&remainTickets)

			// a for loop with condition to check if there are no more tickets
			if remainTickets == 0 {
				fmt.Println("Our conference is booked out, come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is short ")
			}
			if !isValidEmail {
				fmt.Println(" email address is invalid ")
			}
			if !isValidTicketNum {
				fmt.Println("number of tickets you entered is invalid ")
			}
			// fmt.Println("Your input data is invalid, try again")
			// fmt.Printf(" We only have %v tickets remaining, so you can't book %v tickets\n", remainTickets, userTickets)
		}
	}

	city := "London"

	// switch city{
	// case "New York" :
	// 	// execute code for booking NY 
	// case "Singapore" :
	// 	//execute code for singapore 
	// case "London":
	// 	// execute code for London
	// case "Mexico City", "Lagos":
	// 	// they both have the same code, so consildate teh code into one 
	// default:
	// 	fmt.Println("No valid city selected")
	// }
}

func greetUsers(confnamefunc string) {
	fmt.Printf(" Welcome to our %v booking application function\n", confnamefunc)
}


