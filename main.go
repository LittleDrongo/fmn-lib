package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LittleDrongo/fmn-lib/planner"

	"github.com/LittleDrongo/fmn-lib/console/cmd"
	"github.com/LittleDrongo/fmn-lib/console/cmd/loading"
	"github.com/LittleDrongo/fmn-lib/console/color"
)

func main() {

	cmd.Timeout(loading.BAR, 1*time.Second, color.BG_RED, "Автозапуск:", color.BG_RESET)

	fmt.Println()
	fmt.Print(color.BG_RED, " </> ", color.RESET, color.BG_BLACK, " Тестовое приложение", color.DBOLD, " v.0.0 ", color.RESET)
	fmt.Println()

	cron := planner.CreateNewTasklist()

	cron.AddTaskCron(planner.CronRunArguments{
		Second:     "0",
		Minute:     "29",
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
		Minute:     "27",
		Hour:       "*",
		DayOfMount: "*",
		Mount:      "*",
		DayOfWeek:  "*",
	}, func() {
		fmt.Println()
		log.Println(color.CYAN, "Выполнение кода 2", color.RESET)
		log.Println(color.BOLD, "Выполнение кода 2", color.RESET)
		log.Println("Выполнение кода 2")
		log.Println("Выполнение кода 2")
		log.Println("Выполнение кода 2")
	},
		"Тестовый метод2")

	cron.Run()
	// cmd.Waiting(animation.DOTS, "Ожидание: ")
}
