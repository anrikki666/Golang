package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchAndSaveFromAPI(apiURL, filename string) error {
	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Ответ от API сохранен в файл: %s\n", filename)
	return nil
}

func main() {
	apiURL := "https://api.stormglass.io/v2"
	filename := "main.go"

	err := fetchAndSaveFromAPI(apiURL, filename)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
}
