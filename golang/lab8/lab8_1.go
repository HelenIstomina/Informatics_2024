package lab8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Parameters struct {
	A     float64
	B     float64
	x1    float64
	x2    float64
	dx    float64
	slice []float64
}

func ReadFromFile(filename string) (Parameters, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Parameters{}, fmt.Errorf("Ошибка открытия: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values []float64
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return Parameters{}, fmt.Errorf("ошибка парсинга float '%s': %w", line, err)
		}
		values = append(values, num)
	}

	if err := scanner.Err(); err != nil {
		return Parameters{}, fmt.Errorf("ошибка чтения файла: %w", err)
	}
	if len(values) < 6 {
		return Parameters{}, fmt.Errorf("Недостаточно значений. Введите еще раз")
	}
	params := Parameters{
		A:     values[0],
		B:     values[1],
		x1:    values[2],
		x2:    values[3],
		dx:    values[4],
		slice: values[5:],
	}
	return params, nil
}
