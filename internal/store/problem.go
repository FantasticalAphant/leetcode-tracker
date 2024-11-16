package store

import "time"

// Problem is a struct that holds information about a LeetCode question/problem
// This information includes the completion status and the time it was modified
type Problem struct {
	ID        int       `json:"id"`
	Modified  time.Time `json:"modified"`
	Completed bool      `json:"completed"`
	Notes     string    `json:"notes"`
}
