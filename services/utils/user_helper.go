package utils

import (
	"fmt"
	"goexample/models"
	"sort"
)

//Write a utility function for a User object. You will be implementing
//three different functions,
//each with the following function signature: func ([]User) bool.

//AtLeastTwice is a function that returns true if there is a person
//who is at least twice as old as any other person in the list,
//otherwise, the function returns false.
func AtLeastTwice(users []models.UserWithAge) bool {
	fmt.Println("------------------------------")
	fmt.Printf("Inside AtLEastTwice with users: %v\n", users)

	for i := 0; i < len(users); i++ {
		// fmt.Printf("checking user I: %v\n", users[i])
		for j := i + 1; j < len(users); j++ {
			// fmt.Printf("checking user J: %v\n", users[j])
			if users[i].Age*2 <= users[j].Age {
				fmt.Printf("ZOOMER: %v\n", users[i])
				fmt.Printf("BOOMER: %v\n", users[j])
				return true
			}
		}
	}

	return false
}

//TODO: check if w/ range is faster?
func AtLEastTwiceAlt(users []models.UserWithAge) bool {
	fmt.Println("------------------------------")
	fmt.Printf("Inside AtLEastTwiceAlt with users: %v\n", users)

	for _, s := range users {
		// fmt.Printf("checking user: %v\n", s)
		for _, t := range users {
			// fmt.Printf("checking user: %v\n", t)
			if s.Age*2 <= t.Age {
				fmt.Printf("ZOOMER: %v\n", s)
				fmt.Printf("BOOMER: %v\n", t)
				return true
			}
		}
	}
	return false
}

//ExactlyTwice is a function that returns true if there is a person
//who is exactly twice as old as any other person in the list,
//otherwise the function returns false.
func ExactlyTwice(users []models.UserWithAge) bool {
	fmt.Println("------------------------------")
	fmt.Printf("Inside ExactlyTwice with users: %v\n", users)

	for i := 0; i < len(users); i++ {
		// fmt.Printf("checking user I: %v\n", users[i])
		for j := 0; j < len(users); j++ {
			// fmt.Printf("checking user J: %v\n", users[j])
			if users[i].Age == users[j].Age*2 {
				fmt.Printf("user who is exactly twice younger is: %v\n", users[j])
				fmt.Printf("user who is exactly double age is: %v\n", users[i])
				return true
			}
		}
	}

	return false
}

//ConstrainedExactlyTwice is a function that behaves like ExactlyTwice,
//but input age values are guaranteed to always be within the range 18 to 80,
//and this function must perform extremely well
//(can be over-optimized for performance,
//considering time and space complexity).
//TL;DR implement binary search
//the input data has to be ordered by age
func ConstrainedExactlyTwice(users []models.UserWithAge) bool {
	fmt.Println("------------------------------")
	fmt.Printf("Inside ConstrainedExactlyTwice with users: %v\n", users)

	for i, s := range users {
		fmt.Printf("checking user I: %v\n", s)
		//since the binary search checks by age
		//specify the target age as the second parameter
		return BinarySearchAge(users[i+1:], s.Age*2)
	}

	return false
}

//binary search helper
func BinarySearchAge(users []models.UserWithAge, age int) bool {
	fmt.Println("------------------------------")
	fmt.Printf("Inside BinarySearch with users: %v\n", users)

	fmt.Printf("age is %v\n", age)

	var left int = 0
	var right int = len(users) - 1
	var mid int = 0

	for left <= right {
		mid = (left + right) / 2
		if users[mid].Age == age {
			fmt.Printf("user who is exactly twice older is: %v\n", users[mid])
			return true
		} else if users[mid].Age < age {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

//return slice within limits 18 to 80
func GetUsersWithinLimits(users []models.UserWithAge, limit1 int, limit2 int) []models.UserWithAge {
	fmt.Println("------------------------------")
	fmt.Printf("Inside getUsersWithinLimits with users: %v\n", users)

	var usersWithinLimits []models.UserWithAge
	for _, s := range users {
		// fmt.Printf("checking user: %v\n", s)
		//append to slice if age is within limits
		if s.Age >= limit1 && s.Age <= limit2 {
			usersWithinLimits = append(usersWithinLimits, s)
		}
	}
	return usersWithinLimits
}

//sort unordered slice by age
//implemented in services/users@L49
func SortUsersByAge(users []models.User) []models.User {
	fmt.Println("------------------------------")
	fmt.Printf("Inside sortUsersByAge with users: %v\n", users)

	sort.Slice(users, func(i, j int) bool {
		return users[i].Dob > users[j].Dob
	})

	// for _, s := range users {
	// 	fmt.Printf("sorted user: %v\n", s)
	// }
	return users

}
