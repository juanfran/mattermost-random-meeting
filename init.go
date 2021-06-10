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

func getUserListByDate(userId string, users []string, previousMeetings [][]string) []string {
	usersList := []string{}

	for _, meeting := range previousMeetings {
		if Contains(meeting, userId) {
			for _, meetingUserId := range meeting {
				if meetingUserId != userId && !Contains(usersList, meetingUserId) && Contains(users, meetingUserId) {
					usersList = Prepend(usersList, meetingUserId)
				}
			}
		}
	}

	// users you have not met with.
	for _, meetingUserId := range users {
		if meetingUserId != userId && !Contains(usersList, meetingUserId) {
			usersList = Prepend(usersList, meetingUserId)
		}
	}

	return usersList
}

// users is the getUserListByDate of the users involve, so the function return how far in time was the last meeting for all array of users
func weighUserPriority(users [][]string) map[string]int {
	priority := make(map[string]int)

	for _, userList := range users {
		for index, userId := range userList {
			priority[userId] = priority[userId] + index
		}
	}

	return priority
}
func weighUserPriorityCandidate(users [][]string) string {
	usersPriorities := weighUserPriority(users)
	usersPrioritiesList := GetIntStringKeys(usersPriorities)

	sort.SliceStable(usersPrioritiesList, func(i, j int) bool {
		return usersPriorities[usersPrioritiesList[i]] < usersPriorities[usersPrioritiesList[j]]
	})

	return usersPrioritiesList[0]
}

func sortMeetingBySize(meetings [][]string) [][]string {
	sort.SliceStable(meetings, func(i, j int) bool {
		return len(meetings[i]) < len(meetings[j])
	})

	return meetings
}

/*
 previousMeetings: from newest to oldest meeting
*/
func getMeetings(users []string, usersPerMeeting int, previousMeetings [][]string) [][]string {
	meetings := [][]string{}

	if len(users) <= 1 {
		return meetings
	}

	// ShuffleArrayStrings(users)

	alreadyInMeeting := []string{}

	for _, userId := range users {
		if Contains(alreadyInMeeting, userId) {
			continue
		}

		meeting := []string{userId}

		for len(meeting) < usersPerMeeting {
			previousUserMeetingsByDate := [][]string{}

			for _, meetingUserId := range meeting {
				userMeetingCandidates := getUserListByDate(meetingUserId, users, previousMeetings)
				userMeetingCandidates = Filter(userMeetingCandidates, append(alreadyInMeeting, meeting[:]...))

				if len(userMeetingCandidates) > 0 {
					previousUserMeetingsByDate = append(previousUserMeetingsByDate, userMeetingCandidates)
				}
			}

			if len(previousUserMeetingsByDate) > 0 {
				candidate := weighUserPriorityCandidate(previousUserMeetingsByDate)
				meeting = append(meeting, candidate)
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

func runMeetings(users []string, usersPerMeeting int, previousMeetings [][]string) [][]string {
	numMeetings := Combinations(len(users), usersPerMeeting)

	for uint64(len(previousMeetings)) < numMeetings {
		meetings := getMeetings(
			users,
			usersPerMeeting,
			previousMeetings,
		)

		for _, meeting := range meetings {
			if uint64(len(previousMeetings)) < numMeetings {
				previousMeetings = append(previousMeetings, meeting)
			}
		}

	}

	return previousMeetings
}
