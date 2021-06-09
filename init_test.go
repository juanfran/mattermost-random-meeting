package main

import (
	"fmt"
	"testing"

	combinations "github.com/mxschmitt/golang-combinations"
	"github.com/stretchr/testify/assert"
)

func runMeetings(numRuns int, users []string, usersPerMeeting int, previousMeetings [][]string) [][]string {
	for i := 1; i < numRuns; i++ {
		result := getMeetings(
			users,
			usersPerMeeting,
			previousMeetings,
		)

		previousMeetings = append(previousMeetings, result[:]...)
	}

	return previousMeetings
}

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
		if len(combination) == 3 {
			result = append(result, combination)
		}
	}

	return result
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

func TestEvenGetMeetings(t *testing.T) {
	assert := assert.New(t)
	users := []string{"user1", "user2", "user3", "user4", "user5", "user6"}

	previousMeetings := [][]string{}
	combinations := getAllCombinations(users, 3)

	fmt.Println(combinations)

	result := runMeetings(
		len(combinations),
		users,
		3,
		previousMeetings,
	)

	for _, combination := range combinations {
		assert.True(meetingExist(combination, result))
	}
}

// func TestOddGetMeetings(t *testing.T) {
// 	assert := assert.New(t)
// 	users := []string{"user1", "user2", "user3"}

// 	previousMeetings := [][]string{}

// 	result := runMeetings(
// 		len(users)*len(users),
// 		users,
// 		2,
// 		previousMeetings,
// 	)

// 	meeting := []string{"user1", "user2"}
// 	assert.True(meetingExist(meeting, result))

// 	meeting = []string{"user1", "user3"}
// 	assert.True(meetingExist(meeting, result))

// 	meeting = []string{"user2", "user3"}
// 	assert.True(meetingExist(meeting, result))
// }

// modo impar =>
// modo menos gente
// modo añadir uno más de forma justa
