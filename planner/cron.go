package planner

import (
	"fmt"

	"github.com/LittleDrongo/fmn-lib/console/cmd/animation"

	"github.com/LittleDrongo/fmn-lib/console/color"

	"github.com/LittleDrongo/fmn-lib/console/cmd"

	"github.com/robfig/cron"
)

type TaskList struct{ *cron.Cron }

type CronRunArguments struct {
	Second     string
	Minute     string
	Hour       string
	DayOfMount string
	Mount      string
	DayOfWeek  string
}

/*Создать экземпляр структуры планировщика задач*/
func CreateNewTasklist() TaskList {
	return TaskList{cron.New()}
}

/*Создать добавить задачу в планировщик задач*/
func (crn *TaskList) AddTaskCron(args CronRunArguments, code func(), description ...string) {
	fmt.Println(color.BG_GREEN, " Cron ", color.BG_PURPLE, description, color.RESET, args)
	crn.AddFunc(args.Second+" "+args.Minute+" "+args.Hour+" "+args.DayOfMount+" "+args.Mount+" "+args.DayOfWeek, code)
}

/*Запуск выполнение кода с анимацией ожидания*/
func (crn *TaskList) Run() {

	crn.Start()
	fmt.Println()
	cmd.Waiting(animation.DOTS, color.BG_GREEN, "Планировщик в режиме ожидания", color.BG_RESET)

}
