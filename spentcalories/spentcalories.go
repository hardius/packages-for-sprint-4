package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	splitData := strings.Split(data, ",")
	if len(splitData) != 3 {
		return 0, "", 0, errors.New("некорректное количество строк")
	}

	steps, err := strconv.Atoi(splitData[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("некорректное количество шагов")
	}

	duration, err := time.ParseDuration(splitData[2])
	if err != nil {
		return 0, "", 0, err
	}

	return steps, splitData[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, typeOfActivity, duration, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		return "", err
	}

	distance := distance(steps, height)
	meanSpeed := meanSpeed(steps, height, duration)
	var calories float64
	switch typeOfActivity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	result := fmt.Sprintf("Тип тренировки: %s \n", typeOfActivity) +
		fmt.Sprintf("Длительность: %.2f ч. \n", duration.Hours()) +
		fmt.Sprintf("Дистанция: %.2f км. \n", distance) +
		fmt.Sprintf("Скорость: %.2f км/ч \n", meanSpeed) +
		fmt.Sprintf("Сожгли калорий: %.2f \n", calories)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("некорректные параметры")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	result := (weight * meanSpeed * durationInMinutes) / minInH
	return result, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("некорректные параметры")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	result := (weight * meanSpeed * durationInMinutes) / minInH * walkingCaloriesCoefficient
	return result, nil
}
