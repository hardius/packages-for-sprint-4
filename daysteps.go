package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	splitData := strings.Split(data, ",")
	if len(splitData) != 2 {
		return 0, 0, errors.New("некорректное количество строк")
	}

	steps, err := strconv.Atoi(splitData[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, errors.New("некорректное количество шагов")
	}

	duration, err := time.ParseDuration(splitData[1])
	if err != nil {
		return 0, 0, err
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	distance := float64(steps) * stepLength / mInKm

	result := fmt.Sprintf(`Количество шагов: %d.
						   Дистанция составила %.2f км.
						   Вы сожгли %.2f ккал.`, steps, distance, calories)
	return result
}
