package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Cache interface {
	Get(key string) (string, bool)
	Set(key, value string)
	Delete(key string)
}

type cacheImpl struct {
	data map[string]string
}

func (c *cacheImpl) Get(key string) (string, bool) {
	value, ok := c.data[key]
	return value, ok
}

func (c *cacheImpl) Set(key, value string) {
	c.data[key] = value
}

func (c *cacheImpl) Delete(key string) {
	delete(c.data, key)
}

func main() {
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода из приложения нажмите Esc")

	cache := &cacheImpl{
		data: make(map[string]string),
	}

OuterLoop:
	for {
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильные аргументы в формате url описание теги")
				continue OuterLoop
			}

			url := args[0]
			description := args[1]
			tags := args[2:]

			cache.Set(url, fmt.Sprintf("Описание: %s\nТеги: %s\nДата: %s", description, strings.Join(tags, ", "), time.Now().Format(time.RFC3339)))
			fmt.Println("Запись успешно добавлена!")

		case 'l':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Список добавленных url:")

			for url, value := range cache.data {
				fmt.Println("Имя:", url)
				fmt.Println(value)
				fmt.Println()
			}

		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Введите имя ссылки, которую нужно удалить")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			url := strings.TrimSpace(text)

			cache.Delete(url)
			fmt.Println("Запись успешно удалена!")

		default:
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
