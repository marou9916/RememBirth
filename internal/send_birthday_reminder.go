package internal

import (
	"fmt"

	"github.com/marou9916/birthdayReminder.git/models"
	"gopkg.in/gomail.v2"
)

// SendBirthdayReminder sends a birthday reminder email to the friend
func SendBirthdayReminder(friend models.Friend) {
	// Configure the email
	m := gomail.NewMessage()
	m.SetHeader("From", "youraddress@gmail.com") // Replace with your email address
	m.SetHeader("To", friend.Email)              // Send to the friend's email address
	m.SetHeader("Subject", "Birthday Reminder")
	m.SetBody("text/plain", fmt.Sprintf("Hello dear %s,\n\nHope you are doing great. Just to whish you Happy Birthday!\n\nBest regards,\nYour birthday reminder", friend.Surname))

	// Send the email
	d := gomail.NewDialer("smtp.gmail.com", 587, "youraddress@gmail.com", "your_application_password") // Replace with your app password

	// Try to send the email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Printf("Birthday reminder email sent to %s %s\n", friend.Surname, friend.Name)
	}
}
