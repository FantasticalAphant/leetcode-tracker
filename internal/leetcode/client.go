package leetcode

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// TODO: maybe pass the path directory through arguments?
func getFilePath() (string, error) {
	root, err := getProjectRoot()
	if err != nil {
		return "", err
	}

	return filepath.Join(root, "data", "info.json"), nil
}

func getProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found")
		}
		dir = parent
	}
}

// TODO: also get the difficulty (tags if possible as well)
type Response struct {
	StatStatusPairs []struct {
		Stat struct {
			QuestionID int    `json:"question_id"`
			Title      string `json:"question__title"`
		} `json:"stat"`
		Difficulty struct {
			Level int `json:"level"`
		}
	} `json:"stat_status_pairs"`
}

func fetchLeetCodeProblems() (*Response, error) {
	dataFile, err := getFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	var leetcodeData Response
	if err := json.Unmarshal(data, &leetcodeData); err != nil {
		return nil, err
	}
	return &leetcodeData, nil
}

func GetQuestionNameByID(questionID int, data *Response) (string, error) {
	for _, pair := range data.StatStatusPairs {
		if pair.Stat.QuestionID == questionID {
			return pair.Stat.Title, nil
		}
	}
	return "", fmt.Errorf("question with ID %d not found", questionID)
}

func GetQuestionDifficultyByID(questionID int, data *Response) (string, error) {
	for _, pair := range data.StatStatusPairs {
		if pair.Stat.QuestionID == questionID {
			switch pair.Difficulty.Level {
			case 1:
				return "Easy", nil
			case 2:
				return "Medium", nil
			case 3:
				return "Hard", nil
			default:
				return "", fmt.Errorf("difficulty level not supported: %d", pair.Difficulty.Level)
			}
		}
	}
	return "", fmt.Errorf("question with ID %d not found", questionID)
}

func GetQuestionInformation(questionID int) (map[string]string, error) {
	data, err := fetchLeetCodeProblems()

	if err != nil {
		return nil, err
	}

	name, err := GetQuestionNameByID(questionID, data)
	if err != nil {
		return nil, err
	}

	difficulty, err := GetQuestionDifficultyByID(questionID, data)
	if err != nil {
		return nil, err
	}

	return map[string]string{"name": name, "difficulty": difficulty}, nil
}
