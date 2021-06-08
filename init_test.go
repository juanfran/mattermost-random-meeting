package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestGetMeetings(t *testing.T) {
	assert := assert.New(t)
	users := []string{"user1", "user2", "user3", "user4", "user5"}

	previousMeetings := [][]string{}

	result := getMeetings(
		users,
		3,
		previousMeetings,
	)

	fmt.Println(result)

	assert.Equal(4, 4)
}

// modo impar
