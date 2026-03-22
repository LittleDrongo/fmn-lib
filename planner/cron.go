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

/*Create a new task scheduler instance.*/
func CreateNewTasklist() TaskList {
	return TaskList{cron.New()}
}

/*Add a task to the scheduler.*/
func (crn *TaskList) AddTaskCron(args CronRunArguments, code func(), description ...string) error {
	fmt.Println(description, args)
	return crn.AddFunc(args.Second+" "+args.Minute+" "+args.Hour+" "+args.DayOfMount+" "+args.Mount+" "+args.DayOfWeek, code)
}

/*Start execution with a waiting animation.*/
func (crn *TaskList) Run(msg string) {

	crn.Start()
	fmt.Println()
	cmd.Waiting(animation.DOTS, color.BG_GREEN, msg, color.BG_RESET)

}
