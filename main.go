package main

import (
	"fmt"
	"log"

	"github.com/LittleDrongo/fmn-lib/planner"

	"github.com/LittleDrongo/fmn-lib/console/cmd"
	"github.com/LittleDrongo/fmn-lib/console/cmd/animation"
	"github.com/LittleDrongo/fmn-lib/console/cmd/loading"
	"github.com/LittleDrongo/fmn-lib/console/color"
)

func main() {
	timerSample()
}

func timerSample() {
	fmt.Println()

	cmd.Waiting(animation.CLOCK_COLOR, "dsadas")

	fmt.Println(color.BG_BLUE, "      Автозапуск через 10 секунд       ", color.BG_RESET)

	fmt.Println(color.DBOLD, "Нажмите Enter чтобы прервать автозапуск", color.RESET)
	fmt.Println()
	cmd.Autorun(loading.BAR, 5, func() { stopedCode() }, func() { cronSample() })

}

func stopedCode() {

	fmt.Println()
	fmt.Println("Атозапуск прерван")

}

func cronSample() {

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

	cron.Run("Планировщик в режиме ожидания")

}
