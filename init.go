package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, World!")
}

func hola() string {
	return "Hello, world!"
}

/*
past meetings:

user1, user3, user2
user1, user4, user5

number of users: 3

available users: ["user1", "user2", "user3", "user4", "user5", "user6"]
*/

/*
{
  user3: 2,
  user1: 0,
}
*/

func getUserFrequencyMeetings(userId string, users []string, previousMeetings [][]string) map[string]int {
	frequncy := make(map[string]int)

	for _, pairUserId := range users {
		if pairUserId != userId {
			frequncy[pairUserId] = 0
		}
	}

	for _, meeting := range previousMeetings {
		if Contains(meeting, userId) {
			for _, pairUserId := range meeting {
				if pairUserId != userId {
					frequncy[pairUserId]++
				}
			}
		}
	}

	return frequncy
}

func getUserFrequencyMeetingsList(userId string, users []string, previousMeetings [][]string) []string {
	ShuffleArrayStrings(users)

	userFrequencyMeetings := getUserFrequencyMeetings(userId, users, previousMeetings)
	userFrequencyMeetingsList := GetIntStringKeys(userFrequencyMeetings)

	sort.SliceStable(userFrequencyMeetingsList, func(i, j int) bool {
		return userFrequencyMeetings[userFrequencyMeetingsList[i]] < userFrequencyMeetings[userFrequencyMeetingsList[j]]
	})

	return userFrequencyMeetingsList
}

func getMeetings(users []string, usersPerMeeting int, previousMeetings [][]string) [][]string {
	ShuffleArrayStrings(users)

	meetings := [][]string{}
	alreadyInMeeting := []string{}

	for _, userId := range users {
		if Contains(alreadyInMeeting, userId) {
			continue
		}

		meeting := []string{userId}
		frequencyMeetingList := getUserFrequencyMeetingsList(userId, users, previousMeetings)

		frequencyMeetingList = Filter(frequencyMeetingList, alreadyInMeeting)

		for len(meeting) < usersPerMeeting && len(frequencyMeetingList) > 0 {
			newUserId := ""
			newUserId, frequencyMeetingList = Shift(frequencyMeetingList)

			meeting = append(meeting, newUserId)
		}

		meetings = append(meetings, meeting)
		alreadyInMeeting = append(alreadyInMeeting, meeting[:]...)
	}

	return meetings
}

// "user1", "user2", "user3", "user4", "user5"

// ["user1", "user2", "user3"]
// ["user4", "user5", "random"]
