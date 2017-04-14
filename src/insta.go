package app

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// getFollowing gets users that we follow.
func (a *App) getFollowing() {
	log.Println("Collecting your 'Following' list")
	resp, err := a.api.SelfTotalUserFollowing()
	if err != nil {
		panic(err)
	}

	for _, user := range resp.Users {
		username := strings.ToLower(user.Username)
		a.following[username] = true
	}
}

// getFollowers gets users that follow us.
func (a *App) getFollowers() {
	log.Println("Collecting your 'Followers' list")
	resp, err := a.api.SelfTotalUserFollowers()
	if err != nil {
		panic(err)
	}

	for _, user := range resp.Users {
		username := strings.ToLower(user.Username)
		a.followers[username] = true
	}
}

// massUnfollow performs the mass unfollow operation.
func (a *App) massUnfollow() {
	fmt.Printf("!!! Beginning the mass unfollow !!!\n\n")

	// Show the progress as we go.
	var (
		reaped    = 1
		remaining = len(a.leeches)
	)

	for _, username := range a.leeches {
		if _, ok := a.following[username]; !ok {
			fmt.Printf("[ERROR] Username %s not found in following map\n", username)
			continue
		}

		// Unfollow.
		userIDStr := a.getUserId(username)
		log.Printf("- [%d of %d] Unfollow: %s\t\t(UID %s)\n", reaped, remaining, username, userIDStr)

		// Convert the user ID from a string to an int.
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			panic(err)
		}

		_, err = a.api.UnFollow(int64(userID))
		if err != nil {
			log.Panicf("Got error when unfollowing %s: %s", username, err)
		}

		reaped++
		time.Sleep(time.Duration(a.Wait) * time.Second)
	}
}

// getUserId gets the Instagram user PK ID (an int64) as a string.
func (a *App) getUserId(username string) string {
	user, err := a.api.GetUserByUsername(username)
	if err != nil {
		log.Panicf("Can't getUserId %s: %s", username, err)
	}

	return strconv.Itoa(int(user.User.ID))
}
