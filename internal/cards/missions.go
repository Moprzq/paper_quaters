package cards

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode/utf8"
)

type Mission struct {
	Buildings     []string
	MissionNumber int
	ScoreFirst    int
	ScoreSecond   int
	IsAlreadyDone bool
}

type GameMissons struct {
	First  Mission
	Second Mission
	Third  Mission
}

const missionCountInClass = 6

func GetMissons() GameMissons {
	missons := GameMissons{
		First:  GetOneOfFirstMissons(),
		Second: GetOneOfSecondMissons(),
		Third:  GetOneOfThirdMissons(),
	}
	return missons
}

func GetOneOfFirstMissons() Mission {
	ms := getFirstMissons()
	m := ms[rand.Intn(missionCountInClass)]
	return m
}

func GetOneOfSecondMissons() Mission {
	ms := getSecondMissons()
	m := ms[rand.Intn(missionCountInClass)]
	return m
}

func GetOneOfThirdMissons() Mission {
	ms := getThirdMissons()
	m := ms[rand.Intn(missionCountInClass)]
	return m
}

// 🏠
func getFirstMissons() [missionCountInClass]Mission {
	ms := [missionCountInClass]Mission{
		{Buildings: []string{"🏠🏠🏠🏠", "🏠🏠🏠🏠"}, ScoreFirst: 6, ScoreSecond: 3, MissionNumber: 1},
		{Buildings: []string{"🏠🏠🏠🏠🏠🏠", "🏠🏠🏠🏠🏠🏠"}, ScoreFirst: 10, ScoreSecond: 6, MissionNumber: 1},
		{Buildings: []string{"🏠🏠", "🏠🏠", "🏠🏠", "🏠🏠"}, ScoreFirst: 8, ScoreSecond: 4, MissionNumber: 1},
		{Buildings: []string{"🏠🏠🏠", "🏠🏠🏠", "🏠🏠🏠"}, ScoreFirst: 8, ScoreSecond: 4, MissionNumber: 1},
		{Buildings: []string{"🏠🏠🏠🏠🏠", "🏠🏠🏠🏠🏠"}, ScoreFirst: 8, ScoreSecond: 4, MissionNumber: 1},
		{Buildings: []string{"🏠", "🏠", "🏠", "🏠", "🏠", "🏠"}, ScoreFirst: 8, ScoreSecond: 4, MissionNumber: 1},
	}
	return ms
}

func getSecondMissons() [missionCountInClass]Mission {
	ms := [missionCountInClass]Mission{
		{Buildings: []string{"🏠🏠🏠", "🏠🏠🏠", "🏠🏠🏠🏠"}, ScoreFirst: 12, ScoreSecond: 7, MissionNumber: 2},
		{Buildings: []string{"🏠🏠🏠", "🏠🏠🏠🏠🏠🏠"}, ScoreFirst: 8, ScoreSecond: 4, MissionNumber: 2},
		{Buildings: []string{"🏠🏠🏠🏠", "🏠", "🏠", "🏠"}, ScoreFirst: 9, ScoreSecond: 5, MissionNumber: 2},
		{Buildings: []string{"🏠", "🏠", "🏠", "🏠🏠🏠🏠🏠🏠"}, ScoreFirst: 11, ScoreSecond: 6, MissionNumber: 2},
		{Buildings: []string{"🏠🏠🏠🏠🏠", "🏠🏠", "🏠🏠"}, ScoreFirst: 10, ScoreSecond: 6, MissionNumber: 2},
		{Buildings: []string{"🏠🏠🏠🏠", "🏠🏠🏠🏠🏠"}, ScoreFirst: 9, ScoreSecond: 5, MissionNumber: 2},
	}
	return ms
}

func getThirdMissons() [missionCountInClass]Mission {
	ms := [missionCountInClass]Mission{
		{Buildings: []string{"🏠🏠🏠", "🏠🏠🏠🏠"}, ScoreFirst: 7, ScoreSecond: 3, MissionNumber: 3},
		{Buildings: []string{"🏠", "🏠🏠", "🏠🏠", "🏠🏠🏠"}, ScoreFirst: 11, ScoreSecond: 6, MissionNumber: 3},
		{Buildings: []string{"🏠", "🏠🏠🏠🏠", "🏠🏠🏠🏠🏠"}, ScoreFirst: 13, ScoreSecond: 7, MissionNumber: 3},
		{Buildings: []string{"🏠🏠", "🏠🏠🏠🏠🏠"}, ScoreFirst: 7, ScoreSecond: 3, MissionNumber: 3},
		{Buildings: []string{"🏠", "🏠🏠", "🏠🏠🏠🏠🏠🏠"}, ScoreFirst: 12, ScoreSecond: 7, MissionNumber: 3},
		{Buildings: []string{"🏠🏠", "🏠🏠🏠", "🏠🏠🏠🏠🏠"}, ScoreFirst: 13, ScoreSecond: 7, MissionNumber: 3},
	}
	return ms
}

func (m *Mission) Mission() string {
	b := strings.Builder{}
	for _, building := range m.Buildings {
		b.WriteString(fmt.Sprintf("%s(%d) ", building, utf8.RuneCountInString(building)))
	}
	return b.String()
}
