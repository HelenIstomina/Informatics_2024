package lab8

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func CreateNewFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func EditFile(name string, count int) {
	var value string
	file, err := os.OpenFile(name, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for i := 1; i <= count; i++ {
		fmt.Print("Добавление в файл: ")
		fmt.Fscan(os.Stdin, &value)
		file.WriteString(value + "\n")
	}
}

func ReadFile(name string) []string {
	var arr []string
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for {
		data := make([]byte, 64)
		for {
			n, err := file.Read(data)
			if err == io.EOF {
				break
			}
			arr = append(arr, string(data[:n]))
			break
		}
		return arr
	}
}

func SearchInFile(name string, ValueFile string) {
	arr := strings.Split(ReadFile(name)[0], "\n")
	for _, value := range arr {
		if value == ValueFile {
			fmt.Println("Значение в файле найдено")
			break
		}
	}
}


func RunLab8() []string {
	var NameFile string
	fmt.Print("Введите название файла: ")
	fmt.Fscan(os.Stdin, &NameFile)
	CreateNewFile(NameFile)

	var Count int
	fmt.Print("Введите сколько значений хотите занести в файл: ")
	fmt.Fscan(os.Stdin, &Count)
	EditFile(NameFile, Count)

	arr := strings.Split(ReadFile(NameFile)[0], "\n")
	arr = arr[0 : len(arr)-1]

	var Search string
	fmt.Print("Какое значение вы хотите найти в файле: ")
	fmt.Fscan(os.Stdin, &Search)
	SearchInFile(NameFile, Search)

	return arr
}
