package fs

import (
	"encoding/json"
	"os"

	"github.com/drg-bots/telegram/models"
)

func WriteSalutes(s *models.Salutes) error {
	return writeToPath(s, "../store/data/salutes.json")
}

func WriteTrivias(t *models.Trivia) error {
	return writeToPath(t, "../store/data/trivia.json")
}

func WriteDeepDives(dd *models.DeepDives) error {
	return writeToPath(dd, "../store/data/deepdives.json")
}

func writeToPath[T models.DRGData](data *T, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = f.Write(j)
	return err
}
