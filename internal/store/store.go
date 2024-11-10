package store

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

const appName = "heatcold"

// Get directory to store the data based on the user's OS
func getDataDir() (string, error) {
	// First check XDG_DATA_HOME environment variable (Unix/Linux)
	if xdgDataHome := os.Getenv("XDG_DATA_HOME"); xdgDataHome != "" {
		return filepath.Join(xdgDataHome, appName), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	switch runtime.GOOS {
	case "darwin":
		// macOS: use ~/Library/Application Support
		return filepath.Join(homeDir, "Library", "Application Support", appName), nil

	case "windows":
		// Windows: use %LocalAppData%
		if localAppData := os.Getenv("LocalAppData"); localAppData != "" {
			return filepath.Join(localAppData, appName), nil
		}
		return filepath.Join(homeDir, "AppData", "Local", appName), nil

	default:
		// Unix/Linux: use ~/.local/share
		return filepath.Join(homeDir, ".local", "share", appName), nil
	}
}

// ProblemStore holds the list of problems and the path to store the data
type ProblemStore struct {
	Problems map[int]*Problem
	path     string
}

func New() (*ProblemStore, error) {
	dir, err := getDataDir()
	if err != nil {
		return nil, err
	}

	// Create all necessary directories
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	return &ProblemStore{
		Problems: make(map[int]*Problem),
		path:     filepath.Join(dir, "problems.json"),
	}, nil
}

// Load the information from the data file
func (ps *ProblemStore) Load() error {
	data, err := os.ReadFile(ps.path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &ps.Problems)
}

// Save the information to the data file
func (ps *ProblemStore) Save() error {
	data, err := json.MarshalIndent(ps.Problems, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(ps.path, data, 0644)
}

// AddProblem adds a new LeetCode problem on to the existing list
func (ps *ProblemStore) AddProblem(numbers []string) error {
	for _, number := range numbers {
		id, err := strconv.Atoi(number)
		if err != nil {
			return err
		}

		problem := &Problem{
			ID:        id,
			Modified:  time.Now(),
			Completed: true,
		}

		// TODO: use custom errors (i.e., not all unique error or smth like that)
		// allow user to add the questions not in the list
		if _, ok := ps.Problems[id]; ok {
			return errors.New("problem already exists")
		}

		ps.Problems[id] = problem
	}

	return ps.Save()
}

func (ps *ProblemStore) UpdateProblem(number string) error {
	id, err := strconv.Atoi(number)

	if err != nil {
		return err
	}

	if _, ok := ps.Problems[id]; !ok {
		return errors.New("problem not found")
	}

	ps.Problems[id].Completed = !ps.Problems[id].Completed
	ps.Problems[id].Modified = time.Now()

	return ps.Save()
}

func (ps *ProblemStore) RemoveProblem(number string) error {
	id, err := strconv.Atoi(number)

	if err != nil {
		return err
	}

	if _, ok := ps.Problems[id]; !ok {
		return errors.New("problem not found")
	}

	delete(ps.Problems, id)
	return ps.Save()
}
