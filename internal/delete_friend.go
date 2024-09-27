package internal

import (
	"fmt"
	"strings"
)

func DeleteFriend(surname string) {
	friends := LoadFriends()

	if len(friends) == 0 {
		fmt.Println("No friend to delete. List of friends is empty")
		return
	}

	friendDeleted := false
	var friendDeletedName string

	for index, friend := range friends {

		if strings.EqualFold(friend.Surname, surname) {
			friendDeletedName = friend.Name
			friends = append(friends[:index], friends[index+1:]...)
			friendDeleted = true
			break
		}
	}

	if friendDeleted {
		StoreFriends(friends)
		fmt.Printf("Friend with surname %s and named %s deleted successfully\n", surname, friendDeletedName)
	} else {
		fmt.Printf("No friend called %s to delete", surname)
	}
}
