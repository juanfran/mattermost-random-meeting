package main

import (
	"fmt"
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

// func getAllCombinations(users []string, size int) [][]string {
// 	result := [][]string{}
// 	combinations := combinations.All(users)

// 	for _, combination := range combinations {
// 		if len(combination) == size {
// 			result = append(result, combination)
// 		}
// 	}

// 	return result
// }

// func TestGetUserWithExtraMeeting(t *testing.T) {
// 	assert := assert.New(t)

// 	previousMeetings := [][]string{
// 		{"user1", "user2"},
// 		{"user3", "user4"},
// 		{"user5", "user6", "user7"},
// 	}

// 	result := getUserWithExtraMeeting(previousMeetings)

// 	assert.Equal(result, "user7")
// }

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
	userId := "user1"
	users := []string{"user1", "user2", "user3", "user4", "user5", "user6", "user7"}

	previousMeetings := [][]string{
		{"user1", "user2", "user3"},
		{"user6", "user7", "user4"},
		{"user1", "user5", "user3"},
		{"user1", "user2", "user4"},
	}

	result := getUserListByDate(
		userId,
		users,
		previousMeetings,
	)

	assert.Equal(result[0], "user7")
	assert.Equal(result[1], "user6")
	assert.Equal(result[2], "user4")
	assert.Equal(result[3], "user5")
	assert.Equal(result[4], "user3")
	assert.Equal(result[5], "user2")
}

func TestWeighUserPriority(t *testing.T) {
	assert := assert.New(t)

	users := [][]string{
		{"user1", "user2", "user3"},
		{"user3", "user1", "user2"},
	}

	result := weighUserPriority(users)

	assert.Equal(1, result["user1"])
	assert.Equal(3, result["user2"])
	assert.Equal(2, result["user3"])
}

func TestEvenGetMeetings(t *testing.T) {
	assert := assert.New(t)

	users := []string{"user1", "user2", "user3", "user4", "user5", "user6"}

	previousMeetings := [][]string{
		{"user1", "user2", "user4"},
		{"user1", "user5", "user6"},
		{"user3", "user1", "user4"},
		{"user3", "user2", "user5"},
		{"user3", "user6", "user5"},
	}

	result := getMeetings(users, 3, previousMeetings)

	assert.Equal(2, len(result))
	assert.Equal("user1", result[0][0])
	assert.Equal("user3", result[0][1])
	assert.Equal("user6", result[0][2])
}

func TestOddGetMeetings(t *testing.T) {
	assert := assert.New(t)

	users := []string{"user1", "user2", "user3", "user4", "user5", "user6", "user7"}

	previousMeetings := [][]string{
		{"user1", "user2", "user3"},
		{"user4", "user5", "user6"},
	}

	result := getMeetings(users, 3, previousMeetings)

	assert.Equal(2, len(result))
	assert.Equal("user1", result[0][0])
	assert.Equal("user7", result[0][1])
	assert.Equal("user6", result[0][2])
	assert.Equal("user3", result[0][3])
}

func TestGetMeetings(t *testing.T) {
	// assert := assert.New(t)

	users := []string{"user1", "user2", "user3", "user4", "user5", "user6", "user7"}

	previousMeetings := [][]string{}

	numMeetings := 20

	for len(previousMeetings) < numMeetings {
		meetings := getMeetings(users, 3, previousMeetings)

		for _, meeting := range meetings {

			if len(previousMeetings) < numMeetings {
				previousMeetings = append(previousMeetings, meeting)
			}
		}
	}

	fmt.Println(previousMeetings)
}
