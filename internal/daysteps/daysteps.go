package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
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
	slice := strings.Split(data, ",")

	if len(slice) != 2 {
		return 0, 0, errors.New("slice length not equal 2")
	}
	stepCount, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, err
	}
	if stepCount <= 0 {
		return 0, 0, errors.New("step count <= 0")
	}
	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, errors.New("duration <= 0")
	}
	return stepCount, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	stepCount, duration, err := parsePackage(data)
	if err != nil {
		log.Print(err)
		return ""
	}
	if stepCount <= 0 {
		log.Print(err)
		return ""
	}
	length := stepLength * float64(stepCount)
	length = length / mInKm
	spentCalories, _ := spentcalories.WalkingSpentCalories(stepCount, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", stepCount, length, spentCalories)
}
