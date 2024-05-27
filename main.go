package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetchAndSave(url, filename string) error {
	// Отправляем GET-запрос к указанному URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверяем код состояния ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("не удалось получить данные: статус %d", resp.StatusCode)
	}

	// Считываем тело ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Сохраняем тело ответа в файл
	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Ответ от %s сохранен в файл: %s\n", url, filename)
	return nil
}

func main() {
	// URL, на который отправляем запрос
	url := "https://cert.kz/"
	// Имя файла, в который будет сохранен ответ
	filename := "index.html"

	// Выполняем запрос и сохраняем ответ в файл
	err := fetchAndSave(url, filename)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
}
