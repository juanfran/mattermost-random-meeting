package main

import (
	"sort"

	combinations "github.com/mxschmitt/golang-combinations"
)

// ?
func getUserFrequencyMeetings(userId string, users []string, previousMeetings [][]string) map[string]int {
	frequency := make(map[string]int)

	for _, pairUserId := range users {
		if pairUserId != userId {
			frequency[pairUserId] = 0
		}
	}

	for _, meeting := range previousMeetings {
		if Contains(meeting, userId) {
			for _, pairUserId := range meeting {
				if pairUserId != userId {
					frequency[pairUserId]++
				}
			}
		}
	}

	return frequency
}

// ?
func getUserFrequencyMeetingsList(userId string, users []string, previousMeetings [][]string) []string {
	ShuffleArrayStrings(users)

	userFrequencyMeetings := getUserFrequencyMeetings(userId, users, previousMeetings)
	userFrequencyMeetingsList := GetIntStringKeys(userFrequencyMeetings)

	sort.SliceStable(userFrequencyMeetingsList, func(i, j int) bool {
		return userFrequencyMeetings[userFrequencyMeetingsList[i]] < userFrequencyMeetings[userFrequencyMeetingsList[j]]
	})

	return userFrequencyMeetingsList
}

// ?
func getAllCombinations(users []string, size int) [][]string {
	result := [][]string{}
	combinations := combinations.All(users)

	for _, combination := range combinations {
		if len(combination) == size {
			result = append(result, combination)
		}
	}

	return result
}

func getCountUserMeet(userId string, user2Id string, previousMeetings [][]string) int {
	count := 0

	for _, meeting := range previousMeetings {
		if Contains(meeting, userId) && Contains(meeting, user2Id) {
			count++
		}
	}

	return count
}

func getMeetingCandidates(userInMeeting []string, users []string, previousMeetings [][]string) []string {
	priority := make(map[string]int)

	for _, userMeeting := range userInMeeting {
		usersList := []string{}

		for _, meeting := range previousMeetings {
			if Contains(meeting, userMeeting) {
				for _, meetingUserId := range meeting {
					if !Contains(userInMeeting, meetingUserId) && !Contains(usersList, meetingUserId) && Contains(users, meetingUserId) {
						usersList = Prepend(usersList, meetingUserId)
					}
				}
			}
		}

		for _, meetingUserId := range users {
			if !Contains(usersList, meetingUserId) && !Contains(userInMeeting, meetingUserId) {
				priority[meetingUserId] = priority[meetingUserId] - len(users)
			}
		}

		for index, userId := range usersList {
			priority[userId] = priority[userId] + (index * index) + getCountUserMeet(userMeeting, userId, previousMeetings)
		}
	}

	usersPrioritiesList := GetIntStringKeys(priority)

	sort.SliceStable(usersPrioritiesList, func(i, j int) bool {
		return priority[usersPrioritiesList[i]] < priority[usersPrioritiesList[j]]
	})

	return usersPrioritiesList
}

func sortMeetingBySize(meetings [][]string) [][]string {
	sort.SliceStable(meetings, func(i, j int) bool {
		return len(meetings[i]) < len(meetings[j])
	})

	return meetings
}

func getMeetingsShuffleUsers(users []string, usersPerMeeting int, previousMeetings [][]string) [][]string {
	ShuffleArrayStrings(users)

	return getMeetings(users, usersPerMeeting, previousMeetings)
}

/*
 previousMeetings: from newest to oldest meeting
*/
func getMeetings(users []string, usersPerMeeting int, previousMeetings [][]string) [][]string {
	meetings := [][]string{}

	if len(users) <= 1 {
		return meetings
	}

	alreadyInMeeting := []string{}

	for _, userId := range users {
		if Contains(alreadyInMeeting, userId) {
			continue
		}

		meeting := []string{userId}

		for len(meeting) < usersPerMeeting {
			userMeetingCandidates := getMeetingCandidates(meeting, users, previousMeetings)
			userMeetingCandidates = Filter(userMeetingCandidates, append(alreadyInMeeting, meeting[:]...))

			if len(userMeetingCandidates) > 0 {
				meeting = append(meeting, userMeetingCandidates[0])
			} else {
				break
			}
		}

		if len(meeting) < usersPerMeeting {
			for _, userIdFromIncrompleMeeting := range meeting {
				meetings = sortMeetingBySize(meetings)
				meetings[0] = append(meetings[0], userIdFromIncrompleMeeting)
			}
		} else {
			meetings = append(meetings, meeting)
		}

		alreadyInMeeting = append(alreadyInMeeting, meeting[:]...)
	}

	return meetings
}
