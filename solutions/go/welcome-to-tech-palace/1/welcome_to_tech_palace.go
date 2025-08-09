package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    cnt := 0
    f := ""
	for cnt < numStarsPerLine{
        f += "*"
        cnt++
    }
    return f + "\n" + welcomeMsg + "\n" + f
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	ans := strings.Split(oldMsg, "\n")
    return strings.Trim(ans[1], "* ")
}
