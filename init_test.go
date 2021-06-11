package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func meetingExist(users []string, meetings [][]string) bool {
	for _, meeting := range meetings {
		usersInMeeting := true

		for _, userId := range users {
			if !Contains(meeting, userId) {
				usersInMeeting = false
				break
			}
		}

		if usersInMeeting {
			return true
		}
	}

	return false
}

func TestGetUserFrequencyMeetings(t *testing.T) {
	assert := assert.New(t)
	userId := "user1"
	users := []string{"user1", "user2", "user3", "user4", "user5"}

	previousMeetings := [][]string{
		{"user1", "user2", "user3"},
		{"user5", "user3", "user4"},
		{"user1", "user3", "user4"},
		{"user1", "user3", "user4"},
	}

	result := getUserFrequencyMeetings(
		userId,
		users,
		previousMeetings,
	)

	_, ok := result["user1"]

	assert.Equal(ok, false)
	assert.Equal(result["user2"], 1)
	assert.Equal(result["user3"], 3)
	assert.Equal(result["user4"], 2)
	assert.Equal(result["user5"], 0)
}

func TestGetUserFrequencyMeetingsList(t *testing.T) {
	assert := assert.New(t)
	userId := "user1"
	users := []string{"user1", "user2", "user3", "user4", "user5"}

	previousMeetings := [][]string{
		{"user1", "user2", "user3"},
		{"user5", "user3", "user4"},
		{"user1", "user3", "user4"},
		{"user1", "user3", "user4"},
	}

	result := getUserFrequencyMeetingsList(
		userId,
		users,
		previousMeetings,
	)

	assert.Equal(len(result), 4)
	assert.Equal(result[0], "user5")
	assert.Equal(result[1], "user2")
	assert.Equal(result[2], "user4")
	assert.Equal(result[3], "user3")
}

func TestGetUserListByDate(t *testing.T) {
	assert := assert.New(t)
	users := []string{"user1", "user2", "user3", "user4", "user5", "user6", "user7"}

	previousMeetings := [][]string{
		{"user1", "user2", "user3"},
		{"user6", "user4", "user2"},
		{"user1", "user5", "user3"},
		{"user1", "user2", "user4"},
	}

	meeting := []string{"user1", "user2"}

	result := getMeetingCandidates(
		meeting,
		users,
		previousMeetings,
	)

	// user7 first because neither user1 and user2 meet it
	assert.Equal(result[0], "user7")

	// user5 or user6 can be in [1]
	// user2 didn't meet user5
	if result[1] == "user5" {
		assert.Equal(result[2], "user6")
	} else {
		// user1 didn't meet user6
		assert.Equal(result[1], "user6")
		assert.Equal(result[2], "user5")
	}

	assert.Equal(result[3], "user4")
	assert.Equal(result[4], "user3")
}

func TestOddGetMeetings(t *testing.T) {
	assert := assert.New(t)

	users := []string{"user1", "user2", "user3", "user4", "user5", "user6", "user7"}

	previousMeetings := [][]string{
		{"user1", "user2", "user3"},
		{"user4", "user5", "user6"},
		{"user1", "user5", "user6"},
		{"user2", "user4", "user7"},
	}

	result := getMeetings(users, 3, previousMeetings)

	assert.Equal(2, len(result))
	// todo: random fail
	assert.True(meetingExist([]string{"user1", "user3", "user4", "user7"}, result))
	assert.True(meetingExist([]string{"user2", "user5", "user6"}, result))
}

func TestGetMeetings(t *testing.T) {
	assert := assert.New(t)
	frequency := make(map[string]int)
	users := []string{"user1", "user2", "user3", "user4"}

	previousMeetings := [][]string{}

	numMeetings := 18

	for len(previousMeetings) < numMeetings {
		meetings := getMeetings(users, 2, previousMeetings)

		previousMeetings = append(meetings, previousMeetings...)
	}

	testedUser := "user1"

	for _, meeting := range previousMeetings {
		if Contains(meeting, testedUser) {
			for _, userMeeting := range meeting {
				if userMeeting != testedUser {
					frequency[userMeeting]++
				}
			}
		}
	}

	avgMeetingPerUser := (numMeetings / 2) / (len(users) - 1)
	usersMeetUser1 := GetIntStringKeys(frequency)

	assert.Equal(len(usersMeetUser1), 3)

	for _, userId := range usersMeetUser1 {
		assert.Equal(frequency[userId], avgMeetingPerUser)
	}
}

func TestGetMeetingsTwo(t *testing.T) {
	assert := assert.New(t)
	users := []string{"user1", "user2", "user3", "user4"}

	previousMeetings := [][]string{}

	numMeetings := int(Combinations(len(users), 2))
	combinations := getAllCombinations(users, 2)

	for len(previousMeetings) < numMeetings {
		meetings := getMeetings(users, 2, previousMeetings)

		previousMeetings = append(meetings, previousMeetings...)
	}

	for _, combination := range combinations {
		assert.True(meetingExist(combination, previousMeetings))
	}
}
