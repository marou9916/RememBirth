# RememBirth

## Storytelling

The idea for RememBirth came from a personal struggle: I often found it challenging to remember my friends' birthdays. I thought, "What if I had a tool that would help me remember these special dates?" This inspired me to create RememBirth, a solution that centralizes all my friends' birthdays in one place, allowing me to easily track and receive reminders about their special days.

## RememBirth

RememBirth is a command-line tool designed to help you remember and receive birthday reminders for your friends. This project allows you to add friends, check for upcoming birthdays, and automate reminders to yourself via email, ensuring you never miss an important date again. 

Currently, RememBirth sends email reminders to yourself, allowing you to remember your friends' birthdays. However, I'm considering future enhancements, such as the possibility of allowing RememBirth to send birthday wishes directly to your friends! 🎉✨

## Why a Command-Line Tool?

As a developer, I spend a significant amount of time working on my PC, and I wanted to create an efficient tool that fits seamlessly into my workflow. Command-line tools offer speed, flexibility, and minimal distraction, allowing me to quickly add friends, check birthdays, and automate reminders without navigating through complex interfaces. Additionally, building RememBirth as a command-line application was a fun challenge that allowed me to explore new programming concepts.

## Table of Contents
- [Getting Started](#getting-started)
- [How It Works](#how-it-works)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Cron Job Automation](#cron-job-automation)
- [License](#license)

## Getting Started

To get started with RememBirth, you'll need to have Go installed on your machine. You can download Go from the official [Go website](https://golang.org/dl/).

### Prerequisites

- Go version 1.15 or higher
- Access to an SMTP server (e.g., Gmail) for sending emails
- A valid `friends.json` file or you can create one as per the structure mentioned below

### Installation

    1. Clone the repository to your local machine:
    ```bash
    git clone https://github.com/yourusername/birthdayReminder.git
    cd birthdayReminder

    2. Install dependencies:
    ```bash
    go mod tidy

## How It Works

The RememBirth tool works by managing a list of friends stored in a JSON file (friends.json). You can perform various operations such as adding friends, deleting friends, listing friends, checking for upcoming birthdays, and sending email reminders.

    Adding Friends: You can add a friend's name, surname, and birthday to the friends.json file.
    Listing Friends: List all friends stored in your friends.json.
    Checking Upcoming Birthdays: Check for friends whose birthdays are coming up in the next day.
    Sending Emails: Automated emails will be sent to remind you of upcoming birthdays.

## Project Structure

The project is organized as follows:

├── cmd
│   └── main.go                   # The main entry point for the application
├── friends.json                  # JSON file containing friend data
├── go.mod                        # Go module dependencies
├── go.sum                        # Go module dependency checksum
├── internal                      # Internal package containing core functionalities
│   ├── add_friend.go            # Logic for adding a friend
│   ├── birthday_upcoming_check.go # Logic to check upcoming birthdays
│   ├── delete_friend.go         # Logic for deleting a friend
│   ├── list_friends.go          # Logic for listing friends
│   ├── load_friends.go          # Logic for loading friends from JSON
│   ├── send_email_to_remind_birthday.go # Logic to send email reminders
│   └── store_friends.go         # Logic to store friends to JSON
├── models                        # Data models
│   └── friend.go                # Definition of the Friend struct
├── README.md                     # Project documentation
└── remembirth_cron.sh           # Shell script for running the reminder job

## Usage

    1. Running the Application:

        To run the application, use the following command:
            ```bash
            go run ./cmd/main.go

        This will present you with a command-line interface to interact with the RememBirth tool.

    2. Adding friends:

        You can add friends by following the prompts in the application. Ensure that the friend's birthday is entered in the correct format (e.g., YYYY-MM-DD).

    3. Sending Email Reminders:

        Set up your SMTP configuration in the code to send email reminders. Make sure you have the necessary permissions for sending emails.

    4. Checking Upcoming Birthdays:

        The application will check for upcoming birthdays and notify you via email if you have set up the reminder feature.

## Cron Job Automation

To automate the reminder emails, you can set up a cron job. Here's how to do it:

    1. Create a shell script:

        Edit the remembirth_cron.sh file to include the following:

        #!/usr/bin/bash
        cd /path/to/your/birthdayReminder
        go run ./cmd remind

    2. Set Up Cron Job:

        Open your crontab file:
        ```bash
    crontab -e

        Add the following line to schedule the script (this example runs the script daily at 8 AM):
        ```bash
        0 8 * * * /path/to/your/birthdayReminder/remembirth_cron.sh

## License

This project is licensed under the MIT License. See the LICENSE file for details.

Feel free to adjust any sections, such as the GitHub repository link or the file paths, to better suit your project's specifics!
