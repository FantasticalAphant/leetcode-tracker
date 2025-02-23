package leetcode

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getFilePath(path ...string) (string, error) {
	root, err := getProjectRoot()
	if err != nil {
		return "", err
	}

	fullPath := []string{root}
	fullPath = append(fullPath, path...)
	return filepath.Join(fullPath...), nil
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
	dataFile, err := getFilePath("data", "info.json")
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

func GetCodingPatterns() ([]string, error) {
	dataFile, err := getFilePath("data", "patterns.txt")
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	patterns := strings.Split(string(data), "\n")
	var result []string
	for _, pattern := range patterns {
		result = append(result, pattern)
	}

	return result, nil
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

type QuestionInformation struct {
	Name       string
	Difficulty string
}

func GetQuestionRangeInformation(questionIDs ...int) (map[int]QuestionInformation, error) {
	data, err := fetchLeetCodeProblems()

	if err != nil {
		return nil, err
	}

	result := map[int]QuestionInformation{}

	for _, questionID := range questionIDs {

		name, err := GetQuestionNameByID(questionID, data)
		if err != nil {
			return nil, err
		}

		difficulty, err := GetQuestionDifficultyByID(questionID, data)
		if err != nil {
			return nil, err
		}

		result[questionID] = QuestionInformation{
			Name:       name,
			Difficulty: difficulty,
		}
	}

	return result, nil
}

func GetAllQuestionInformation() (map[string]string, error) {
	data, err := fetchLeetCodeProblems()
	info := make(map[string]string)

	if err != nil {
		return nil, err
	}

	for _, pair := range data.StatStatusPairs {
		id := pair.Stat.QuestionID
		name, err := GetQuestionNameByID(id, data)
		if err != nil {
			return nil, err
		}

		difficulty, err := GetQuestionDifficultyByID(id, data)
		if err != nil {
			return nil, err
		}

		// Just prepend the question ID for easy referencing
		info[strconv.Itoa(id)+") "+name] = difficulty
	}

	return info, nil
}
