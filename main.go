package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const confernceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicket {

		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(2)
		sendTickets(userTickets, firstName, lastName, email)

		firstNames := (bookings)
		fmt.Printf("The first names of bookings are: %v\n ", firstNames)

		if remainingTickets == 0 {
			fmt.Print("Our confrerence is booked outt. Come back next year. ")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name  you enterd is too short. ")
		}
		if !isValidEmail {

			fmt.Println("email address you entered doesn't contain @ sign.")
		}
		if !isValidTicket {

			fmt.Println("number of tickets you entered is invalid. ")
		}
		fmt.Println("Your input data is invalid, try again")
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n. ", conferenceName)
	fmt.Printf("We have total of %v tickets and  %v are stil available.\n", confernceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n.", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remianing for %v\n", remainingTickets, conferenceName)

}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v  %v", userTickets, firstName, lastName)
	fmt.Printf("##############\n")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", tickets, email)
	fmt.Printf("##############")
	wg.Done()

}
