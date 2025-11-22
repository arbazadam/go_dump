package main

import "fmt"

/*

Alright — here’s a clean, tough, real-world flavored problem that forces you to think about state, caching, and correctness. No snippets, no scaffolding.

⸻

:jigsaw: Problem: Event Stream Sessionizer

You are given a stream of user activity events.
Each event contains:
	•	a userID (string)
	•	a timestamp in milliseconds (int, strictly increasing per user but NOT globally)
	•	an eventType (string)

Your job is to compute user sessions.

Session rules:

A session for a user continues as long as each consecutive event from that same user happens within 30 minutes (1800000 ms) of the previous event.

If more than 30 minutes pass since the user’s last event, you must start a new session.

You must compute for every user:
	•	How many total sessions they had.
	•	The average length of their sessions, where length =
(lastEventTimestamp - firstEventTimestamp) of that session.

Input characteristics:
	•	You may process up to 50 million events.
	•	Events come in random order across users:
	•	Example: user A @ t=1000, user B @ t=50, user A @ t=2000, etc.
	•	Events for the same user are always in increasing timestamp order.
	•	You cannot sort the entire event list (too large).
	•	Memory must stay efficient — think O(k) where k = number of active users.

Output:

For each user, compute:
totalSessions
averageSessionLength (float64, two-decimal precision is fine)
Order of users does not matter.

Example scenario (conceptual):

User A events:
	•	10:00
	•	10:10
	•	10:20
	•	11:00 → new session
	•	11:05

This forms:
	•	Session 1: 10:00 → 10:20
	•	Session 2: 11:00 → 11:05

totalSessions = 2
averageSessionLength = ((20 min) + (5 min)) / 2 = 12.5 minutes

Your task:

Write the code that processes the stream and computes the result.
*/

func main() {
	userEvents := []UserActivity{
		{UserId: 1, Timestamp: 1000, EventType: "login"},
		{UserId: 2, Timestamp: 2000, EventType: "view"},
		{UserId: 1, Timestamp: 1000000, EventType: "click"},
		{UserId: 1, Timestamp: 3000000, EventType: "logout"},
		{UserId: 2, Timestamp: 1500000, EventType: "click"},
		{UserId: 3, Timestamp: 500, EventType: "login"},
		{UserId: 3, Timestamp: 2000000, EventType: "play"},
	}

	//A session for a user continues as long as each consecutive event from that same user happens within 30 minutes (1800000 ms) of the previous event.
	userStats := map[int]UserStats{}
	for _, userEvent := range userEvents {
		val, ok := userStats[userEvent.UserId]
		if !ok {
			userStats[userEvent.UserId] = UserStats{TotalSessions: 1, CumulativeSessionTime: 0, LastEventTime: userEvent.Timestamp, CurrentSessionStart: userEvent.Timestamp}
			continue
		}
		if isNewSession := userEvent.Timestamp-val.LastEventTime <= 1800000; !isNewSession {
			val.CumulativeSessionTime += val.LastEventTime - val.CurrentSessionStart
			val.TotalSessions += 1
			val.CurrentSessionStart = userEvent.Timestamp
		}
		val.LastEventTime = userEvent.Timestamp
		userStats[userEvent.UserId] = val
	}
	fmt.Println(userStats)
}

type UserActivity struct {
	UserId    int
	Timestamp int
	EventType string
}

type UserStats struct {
	TotalSessions         int
	CumulativeSessionTime int
	LastEventTime         int
	CurrentSessionStart   int
}

//pick one from the loop process it and add it to the map.
