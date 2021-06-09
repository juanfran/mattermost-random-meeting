package main

import (
	"fmt"
	"testing"

	combinations "github.com/mxschmitt/golang-combinations"
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

func TestEvenGetMeetings(t *testing.T) {
	assert := assert.New(t)
	users := []string{"user1", "user2", "user3", "user4", "user5", "user6"}

	previousMeetings := [][]string{}
	combinations := getAllCombinations(users, 3)

	result := runMeetings(
		users,
		3,
		previousMeetings,
	)

	fmt.Println(len(result))

	assert.Equal(len(combinations), Combinations(len(users), 3))
	assert.Equal(len(combinations), len(result))

	for _, combination := range combinations {
		assert.True(meetingExist(combination, result))
	}
}

//cuando termina puede haber personas sin asignar

// func TestOddGetMeetings(t *testing.T) {
// 	assert := assert.New(t)
// 	users := []string{"user1", "user2", "user3"}

// 	previousMeetings := [][]string{}

// 	result := runMeetings(
// 		users,
// 		2,
// 		previousMeetings,
// 	)

// 	combinations := getAllCombinations(users, 2)
// 	fmt.Println(combinations)

// 	fmt.Println(result)

// 	assert.True(true)

// 	// meeting := []string{"user1", "user2"}
// 	// assert.True(meetingExist(meeting, result))

// 	// meeting = []string{"user1", "user3"}
// 	// assert.True(meetingExist(meeting, result))

// 	// meeting = []string{"user2", "user3"}
// 	// assert.True(meetingExist(meeting, result))
// }

// modo impar =>
// modo menos gente
// modo añadir uno más de forma justa
