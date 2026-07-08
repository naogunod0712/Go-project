package helper

import "strings"

func GetFirstname(bookingslice []string) []string {
	var firstNames = []string{}
	for _, bookingloop := range bookingslice {
		var names = strings.Split(bookingloop, " , ")
		firstNames = append(firstNames, names[0])
	}
	return firstNames
	// fmt.Printf(" Hey %v, there are now  %v tickets left for  %v\n", firstNames, remainTickets, confname)

}

