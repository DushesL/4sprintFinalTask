package spentcalories

import (
	"errors"
	"fmt"
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
	slice := strings.Split(data, ",")

	if len(slice) != 3 {
		return 0, "", 0, errors.New("Длина слайса не равна 3")
	}
	stepCount, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, "", 0, err
	}
	if stepCount <= 0 {
		return 0, "", 0, errors.New("Кол-во шагов <= 0")
	}
	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return 0, "", 0, err
	}
	if duration <= 0 {
		return 0, "", 0, errors.New("Продолжительность <= 0")
	}
	return stepCount, slice[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	length := stepLength * float64(steps)
	length = length / mInKm
	return length
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distance := distance(steps, height)
	avgSpeed := distance / duration.Hours()
	return avgSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	stepCount, trainingType, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	switch trainingType {
	case "Ходьба":
		distance := distance(stepCount, height)
		speed := meanSpeed(stepCount, height, duration)
		spentCalories, _ := WalkingSpentCalories(stepCount, weight, height, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration.Hours(), distance, speed, spentCalories), nil
	case "Бег":
		distance := distance(stepCount, height)
		speed := meanSpeed(stepCount, height, duration)
		spentCalories, _ := RunningSpentCalories(stepCount, weight, height, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration.Hours(), distance, speed, spentCalories), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("Кол-во шагов <= 0")
	}
	if weight <= 0 {
		return 0, errors.New("Вес <= 0")
	}
	if height <= 0 {
		return 0, errors.New("Рост <= 0")
	}
	if duration <= 0 {
		return 0, errors.New("Длительность <= 0")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	spentCalories := duration.Minutes() * weight * avgSpeed / minInH
	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Некорректные входные данные")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	spentCalories := duration.Minutes() * weight * avgSpeed / minInH * walkingCaloriesCoefficient
	return spentCalories, nil
}
