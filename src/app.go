package app

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ahmdrz/goinsta"
)

const Version = "0.1.0"

// Type App is the Follow-Sync Application.
type App struct {
	// User configurable.
	Wait      int // How long to wait between unfollows.
	api       *goinsta.Instagram
	username  string
	password  string
	authed    bool            // are we logged in successfully?
	following map[string]bool // who are we following?
	followers map[string]bool // who follows us?
	leeches   []string        // users we follow who don't follow us back
}

// New creates a new app.
func New() *App {
	return &App{
		Wait:      60,
		following: map[string]bool{},
		followers: map[string]bool{},
		leeches:   []string{},
	}
}

// Run is the entry point to the program.
func (a *App) Run() {
	// Ask for login.
	a.login()
	defer a.logout()

	// Collect data.
	log.Println("Beginning the data collection process...")
	a.getFollowers()
	a.getFollowing()
	a.writeCSV()

	// Compare data.
	log.Println("Comparing the lists to each other...")
	a.compareLists()

	// Lay it all out there. If the user doesn't consent to the mass
	// unfollow that will come, the program exits here.
	log.Println("Telling it how it is...")
	a.tellItHowItIs()

	// Reap the leeches.
	a.massUnfollow()
}

// login asks the user for their credentials and logs them in via Instagram's
// unofficial API.
func (a *App) login() {
	// Prompt and try logging in loop.
	for !a.authed {
		// Ask the user for their credentials.
		for a.username == "" {
			a.username = input("Instagram Username: ")
		}
		for a.password == "" {
			a.password = getPassword("Password (no echo): ")
		}

		// Try the login.
		a.api = goinsta.New(a.username, a.password)

		if err := a.api.Login(); err != nil {
			fmt.Printf("Login error: %s\n", err)
			a.username = ""
			a.password = ""

			// Ask to retry?
			if getAnswer("Retry? [yn] ", "y", "n") == "n" {
				os.Exit(1)
			}

			continue
		}

		a.authed = true
	}
}

// compareLists compares the following to the followers.
func (a *App) compareLists() {
	// See who we're following that doesn't love us back.
	for username, _ := range a.following {
		if _, ok := a.followers[username]; !ok {
			a.leeches = append(a.leeches, username)
		}
	}
}

// tellItHowItIs does as it says.
func (a *App) tellItHowItIs() {
	if len(a.leeches) == 0 {
		fmt.Println("Congrats! Everybody you follow also follows you back!")
		os.Exit(0)
	}

	// Sum up the numbers.
	var (
		numFollowing = len(a.following)
		numFollowers = len(a.followers)
		numLeeches   = len(a.leeches)
		numLoyal     = len(a.following) - len(a.leeches) // The ones who mutually follow us!
	)

	// Show them the list.
	fmt.Printf("You are following these Instagram users who don't follow you back:\n\n"+
		strings.Join(a.leeches, ", ")+
		"\n\n"+
		"If you'd like to compare details, open `follower-lists-%s.csv`.\n\n",
		a.username,
	)

	// Sum up the numbers.
	fmt.Printf("Of the %d users you follow, %d follow you back and %d do not.\n",
		numFollowing,
		numLoyal,
		numLeeches,
	)
	fmt.Println("If you unfollow these leeches, your new Followers/Following ratio will be:")
	fmt.Printf("Before:  %d followers / %d following\n", numFollowers, numFollowing)
	fmt.Printf("After:   %d followers / %d following\n\n", numFollowers, numLoyal)

	answer := getAnswer(
		fmt.Sprintf("Is it OKAY to unfollow these %d users? [yN] ", len(a.leeches)),
		"y", "n",
	)
	if answer != "y" {
		fmt.Println("Bailing!")
		os.Exit(0)
	}
}

// logout logs out of Instagram.
func (a *App) logout() {
	a.api.Logout()
}
