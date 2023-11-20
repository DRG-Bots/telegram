package store

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/drg-bots/telegram/fs"
	"github.com/drg-bots/telegram/models"
)

type Store interface {
	GetSalute() string
	GetTrivia() string
	GetDeepDive() string
	GetEliteDeepDive() string
	GetDeepDives() string
}

type store struct {
	salutes   models.Salutes
	trivia    models.Trivia
	deepDives models.DeepDives
}

func New() (Store, error) {
	// read json files from data folder
	salutes, err := fs.ReadSalutes()
	if err != nil {
		return nil, err
	}

	trivia, err := fs.ReadTrivia()
	if err != nil {
		return nil, err
	}

	deepDives, err := fs.ReadDeepDives()
	if err != nil {
		return nil, err
	}

	return &store{
		salutes:   *salutes,
		trivia:    *trivia,
		deepDives: *deepDives,
	}, nil

}

func (s store) GetSalute() string {
	l := len(s.salutes.Data)
	return s.salutes.Data[rand.Intn(l)]
}

func (s store) GetTrivia() string {
	l := len(s.trivia.Data)
	return s.trivia.Data[rand.Intn(l)]
}

func (s *store) GetDeepDive() string {
	return s.getDeepDiveByType([]string{"Deep Dive"})[0].String()
}

func (s *store) GetEliteDeepDive() string {
	return s.getDeepDiveByType([]string{"Elite Deep Dive"})[0].String()
}

func (s *store) GetDeepDives() string {
	dd := s.getDeepDiveByType([]string{"Deep Dive", "Elite Deep Dive"})

	res := ""
	for _, d := range dd {
		res += d.String() + "\n"
	}

	return strings.Trim(res, "\n")
}

// getDeepDiveByType is a generic function that retrieves the deep dive, the elite deep dive or both.
func (s *store) getDeepDiveByType(types []string) []*models.DeepDive {
	err := s.tryFetchDeepDives()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	res := []*models.DeepDive{}

	for _, t := range types {
		for _, v := range s.deepDives.Variants {
			if v.Type == t {
				res = append(res, &v)
			}
		}
	}
	return res
}

func (s *store) tryFetchDeepDives() error {
	// if the deep dives are up to date early return
	if !s.deepDives.EndTime.Before(time.Now()) {
		return nil
	}

	res, err := http.Get("https://drgapi.com/v1/deepdives")
	if err != nil {
		return err
	}

	// decode response
	var dd models.DeepDives
	err = json.NewDecoder(res.Body).Decode(&dd)

	err = fs.WriteDeepDives(&dd)
	if err != nil {
		return err
	}

	return nil
}
