package techpalace
import ("strings" 
        "fmt"
       )

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    stars := strings.Repeat("*", numStarsPerLine)
	result := fmt.Sprintf("%s\n%s\n%s", stars, welcomeMsg, stars)
    return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    noStars := strings.ReplaceAll(oldMsg, "*", "")
	return  strings.TrimSpace(noStars)
}
