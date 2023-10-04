# remail
Personal use project that reminds me of my tasks in todoist app

## My initial problem
I use the todoist app for organize my monthly tasks. However, todoist doesn't allow notifications in the free plan and this behavior makes me forget about some tasks.

## Solution
So, I implemented an application to consume the todoist API and send me daily emails to remind me of tasks

## Implementation
I used the standart go library to create a http client and consume the todoist API to get all the data for my tasks. After that, I used the gopkg.in/gomail.v2
package to create an email sending service. So, having this two functionalities working, I used the channel and the tick tool in golang to synchronize
all this services whit my local time and send me recurring emails.
