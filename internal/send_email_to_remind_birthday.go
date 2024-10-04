package internal

import (
	"fmt"

	"github.com/marou9916/birthdayReminder.git/models"
	"gopkg.in/gomail.v2"
)

// SendBirthdayReminder sends a birthday reminder email to yourself
func SendBirthdayReminderEmail(friend models.Friend) {
	// Configure the email
	m := gomail.NewMessage()
	m.SetHeader("From", "youraddress@gmail.com") // Replace with your email address
	m.SetHeader("To", "youraddress@gmail.com")   // Send the email to yourself
	m.SetHeader("Subject", "Birthday Reminder")
	m.SetBody("text/plain", fmt.Sprintf("Don't forget %s %s's birthday on %s!", friend.Name, friend.Surname, friend.Birthday.Format("2006-01-02")))

	// Send the email
	d := gomail.NewDialer("smtp.gmail.com", 587, "youraddress@gmail.com", "your_application_password") // Replace with your app password

	// Try to send the email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Printf("Birthday reminder sent for %s %s\n", friend.Surname, friend.Name)
	}
}
