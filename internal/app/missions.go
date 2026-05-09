package app

import (
	"math/rand"
	"path"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mission struct {
	Front *ebiten.Image
	Back  *ebiten.Image
	Done  bool
}

type GameMissions struct {
	First  Mission
	Second Mission
	Third  Mission
}

func newMissions(cache *imageCache) (GameMissions, error) {
	missionLevels := []string{"1", "2", "3"}
	missions := make([]Mission, 0, len(missionLevels))

	for _, level := range missionLevels {
		missionNumber := rand.Intn(missionsInLevel) + 1
		dir := "missions/" + level
		fileName := strconv.Itoa(missionNumber) + ".jpg"
		backFileName := strconv.Itoa(missionNumber) + "_rev.jpg"

		front, err := cache.load(path.Join(dir, fileName))
		if err != nil {
			return GameMissions{}, err
		}

		back, err := cache.load(path.Join(dir, backFileName))
		if err != nil {
			return GameMissions{}, err
		}

		missions = append(missions, Mission{
			Front: front,
			Back:  back,
		})
	}

	return GameMissions{
		First:  missions[0],
		Second: missions[1],
		Third:  missions[2],
	}, nil
}
