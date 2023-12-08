package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var (
		path  string
		album string
	)

	app := &cli.App{
		Name:  "cloudphoto",
		Usage: "Приложение \"Фотоархив\"",
		//EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "upload",
				Aliases: []string{"u"},
				Usage:   "Отправка фотографий в облачное хранилище",
				Action: func(cCtx *cli.Context) error {
					//TODO: логика тут (вызов функции)
					fmt.Println("Ваш флаг: ", path)
					fmt.Println("Ваш флаг: ", album)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "album",
						Value:       "ALBUM",
						Usage:       "Название альбома",
						Destination: &album,
					},
					&cli.StringFlag{
						Name:        "path",
						Value:       "PHOTOS_DIR",
						Usage:       "Название каталога",
						Destination: &path,
					},
				},
			},
			{
				Name:    "download",
				Aliases: []string{"d"},
				Usage:   "Загрузка фотографий из облачного хранилища",
				Action: func(cCtx *cli.Context) error {
					//TODO: логика тут (вызов функции)
					fmt.Println("Ваш флаг: ", path)
					fmt.Println("Ваш флаг: ", album)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "album",
						Value:       "ALBUM",
						Usage:       "Название альбома",
						Destination: &album,
					},
					&cli.StringFlag{
						Name:        "path",
						Value:       "PHOTOS_DIR",
						Usage:       "Название каталога",
						Destination: &path,
					},
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "Вывод списка альбомов и фотографий в облачном хранилище",
				Action: func(cCtx *cli.Context) error {
					//TODO: логика тут (вызов функции)
					fmt.Println("Ваш флаг: ", path)
					fmt.Println("Ваш флаг: ", album)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "album",
						Value:       "ALBUM",
						Usage:       "Название альбома",
						Destination: &album,
					},
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"del"},
				Usage:   "Удаление альбомов и фотографий в облачном хранилище",
				Action: func(cCtx *cli.Context) error {
					//TODO: логика тут (вызов функции)
					fmt.Println("Ваш флаг: ", path)
					fmt.Println("Ваш флаг: ", album)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "album",
						Value:       "ALBUM",
						Usage:       "Название альбома",
						Destination: &album,
					},
				},
			},
			{
				Name:    "mksite",
				Aliases: []string{"mk"},
				Usage:   "Формирование и публикация веб-страниц фотоархива",
				Action: func(cCtx *cli.Context) error {
					//TODO: логика тут (вызов функции)
					fmt.Println("Ваш флаг: ", path)
					fmt.Println("Ваш флаг: ", album)
					return nil
				},
			},
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Инициализация программы",
				Action: func(cCtx *cli.Context) error {
					//TODO: логика тут (вызов функции)
					fmt.Println("Ваш флаг: ", path)
					fmt.Println("Ваш флаг: ", album)
					return fmt.Errorf("Все сломал!!!")
				},
			},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Println("Пожалуйста укажите команду. Чтобы узнать какие команды поддерживаются.")
			fmt.Println("Введите: ./cloudphoto help")
			//name := "Nefertiti"
			//if cCtx.NArg() > 0 {
			//	name = cCtx.Args().Get(0)
			//}
			//if cCtx.String("lang") == "spanish" {
			//	fmt.Println("Hola", name)
			//} else {
			//	fmt.Println("Hello", name)
			//}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
