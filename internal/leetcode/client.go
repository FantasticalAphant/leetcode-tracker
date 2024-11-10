package leetcode

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const leetcodeAPI = "https://leetcode.com/api/problems/all/"

type Response struct {
	StatStatusPairs []struct {
		Stat struct {
			QuestionID int    `json:"question_id"`
			Title      string `json:"question__title"`
		} `json:"stat"`
	} `json:"stat_status_pairs"`
}

func fetchLeetCodeProblems() (*Response, error) {
	resp, err := http.Get(leetcodeAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var leetcodeData Response
	if err := json.NewDecoder(resp.Body).Decode(&leetcodeData); err != nil {
		return nil, err
	}
	return &leetcodeData, nil
}

func GetQuestionNameByID(questionID int) (string, error) {
	data, err := fetchLeetCodeProblems()
	if err != nil {
		return "", err
	}

	for _, pair := range data.StatStatusPairs {
		if pair.Stat.QuestionID == questionID {
			return pair.Stat.Title, nil
		}
	}
	return "", fmt.Errorf("question with ID %d not found", questionID)
}
