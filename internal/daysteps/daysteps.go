package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	slice := make([]string, 0, 2)
	var str string = ""
	for i, symb := range data {
		if symb != ',' {
			str += string(symb)
		} else {
			slice = append(slice, str)
			str = ""
		}
		if i == (len(data) - 1) {
			slice = append(slice, str)
		}
	}
	if len(slice) != 2 {
		return 0, 0, errors.New("")
	}
	stepCount, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, err
	}
	if stepCount <= 0 {
		return 0, 0, errors.New("")
	}
	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, errors.New("")
	}
	return stepCount, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	stepCount, duration, err := parsePackage(data)
	if err != nil {
		log.Println()
		return ""
	}
	if stepCount <= 0 {
		log.Println()
		return ""
	}
	length := stepLength * float64(stepCount)
	length = length / mInKm
	spentCalories, _ := spentcalories.WalkingSpentCalories(stepCount, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", stepCount, length, spentCalories)
}
