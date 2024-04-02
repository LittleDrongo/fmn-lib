package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/LittleDrongo/fmn-lib/planner"

	"github.com/LittleDrongo/fmn-lib/errors"

	"github.com/LittleDrongo/fmn-lib/console/cmd"
	"github.com/LittleDrongo/fmn-lib/console/cmd/animation"
	"github.com/LittleDrongo/fmn-lib/console/cmd/loading"
	"github.com/LittleDrongo/fmn-lib/console/color"

	"gopkg.in/mail.v2"
)

func main() {

	// cmd.Input("Нажми: ")
	cmd.Timeout(loading.CUBES, 1*time.Second, color.BG_RED, "Автозапуск:", color.BG_RESET)
	cmd.Timeout(loading.BAR, 1*time.Second, color.BG_RED, "Автозапуск:", color.BG_RESET)

	cmd.Waiting(animation.DOTS, "Ожидание: ")

	pass := cmd.Input("dasd: ")
	fmt.Println(color.BG_RED, color.YELLO, pass, color.BG_RESET, color.RESET)
	fmt.Println(pass)

	setting := mail.Dialer{
		Host:     "",
		Port:     0,
		Username: "",
		Password: "",
	}

	fmt.Println(setting)

	file, err := os.Open("file.txt")
	errors.Print(err, "Тестовая отработка ошибки с логом")
	errors.Println(err, "Тестовая отработка ошибки с логом")
	// errors.Fatalln(err, "Ошибка с остановкой программы")
	fmt.Println(file)

	fmt.Println()
	fmt.Print(color.BG_RED, " </> ", color.RESET, color.BG_BLACK, " Тестовое приложение", color.DBOLD, " v.0.0 ", color.RESET)
	fmt.Println()

	cron := planner.CreateNewTasklist()

	cron.AddTaskCron(planner.CronRunArguments{
		Second:     "0",
		Minute:     "45",
		Hour:       "*",
		DayOfMount: "*",
		Mount:      "*",
		DayOfWeek:  "*",
	}, func() {
		fmt.Println()
		log.Println("Выполнение кода 1")
	},
		"Тестовый метод")

	cron.AddTaskCron(planner.CronRunArguments{
		Second:     "0",
		Minute:     "46",
		Hour:       "*",
		DayOfMount: "*",
		Mount:      "*",
		DayOfWeek:  "*",
	}, func() {
		fmt.Println()
		log.Println("Выполнение кода 2")
	},
		"Тестовый метод2")

	cron.Run()
}
