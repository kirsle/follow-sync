package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// input asks a question to the console.
func input(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(question)

	answer, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(answer)
}

// getAnswer gets a limited set of possible answers.
//
// In case the user hits Return without answering, the *FINAL* answer
// in the `acceptable...` param is the one given. Example prompt `[yN]` where
// `n` should be the safer response.
func getAnswer(question string, acceptable ...string) string {
	for {
		answer := input(question)

		answer = strings.ToLower(answer)
		for _, cmp := range acceptable {
			if answer == strings.ToLower(cmp) {
				return cmp
			}
		}

		fmt.Printf(
			"Unacceptable answer. Choose one of: " +
				strings.Join(acceptable, ", ") + "\n",
		)
	}
}

// getpasswd asks the user for their password (masks the echo).
func getPassword(question string) string {
	fmt.Printf(question)

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(bytePassword))
}
