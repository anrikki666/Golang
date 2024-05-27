package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func fetchAndSave(url, filename string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("не удалось получить данные: статус %d", resp.StatusCode)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }  
    err = ioutil.WriteFile(filename, body, 0644)
    if err != nil {
        return err
    }
    fmt.Printf("Ответ от %s сохранен в файл: %s\n", url, filename)
    return nil
}
func main() {

    url := "https://cert.kz/"
    filename := "index.html"
    err := fetchAndSave(url, filename)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }

}
