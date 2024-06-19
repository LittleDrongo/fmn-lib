package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LittleDrongo/fmn-lib/planner"
	"github.com/LittleDrongo/fmn-lib/secret"
	"github.com/LittleDrongo/fmn-lib/utils/jsn"
	"github.com/LittleDrongo/fmn-lib/utils/yam"

	"github.com/LittleDrongo/fmn-lib/console/cmd"
	"github.com/LittleDrongo/fmn-lib/console/cmd/loading"
	"github.com/LittleDrongo/fmn-lib/console/color"
)

type mySettingsTestTest struct {
	Ms myStructTestTest
	Me myEmployTestTest
}

type myStructTestTest struct {
	Date       time.Time
	Coast      float64
	otherField float64
}

type myEmployTestTest struct {
	Name string
	Age  int
	Done bool
}

type loginPassword struct {
	Login    string
	Password string
	Admin    bool
	Age      int
	Sum      float64
	Banned   bool
}

const KEY = "1234567890123456"

func main() {

	mySet := mySettingsTestTest{
		Ms: myStructTestTest{
			Date:  time.Now(),
			Coast: 53.4343,
		},

		Me: myEmployTestTest{
			Name: "Alexey",
			Age:  35,
			Done: true,
		},
	}

	yam.ColorPrint(mySet)

	yam.ColorPrint(loginPassword{
		Login:    "mylogin",
		Password: "y12345",
		Admin:    false,
		Banned:   true,
		Age:      53,
		Sum:      321.6454,
	})

}

func secretSample() {
	key := "examplekey123456" // 16, 24 или 32 байта длина ключа для AES-128, AES-192 и AES-256 соответственно
	plaintext := cmd.Password("Введите ваш секрет: ")

	encrypted, err := secret.Encrypt(key, plaintext)
	if err != nil {
		fmt.Println("Ошибка при шифровании:", err)
		return
	}

	fmt.Println("Зашифрованный текст:", encrypted)

	decrypted, err := secret.Decrypt(key, encrypted)
	if err != nil {
		fmt.Println("Ошибка при расшифровке:", err)
		return
	}

	fmt.Println("Расшифрованный текст:", decrypted)
}

func yamlExportSample() {

	mySet := mySettingsTestTest{
		Ms: myStructTestTest{
			Date:  time.Now(),
			Coast: 53.4343,
		},

		Me: myEmployTestTest{
			Name: "Alexey",
			Age:  35,
			Done: true,
		},
	}

	yam.Export(mySet, "data/myfile.yaml")

}

func yamlImportAndPrintSample() {

	var mySetTwo mySettingsTestTest

	yam.Import("data/myfile.yaml", &mySetTwo)
	yam.Print(mySetTwo)

}

func jsonExportSample() {

	mySet := mySettingsTestTest{
		Ms: myStructTestTest{
			Date:  time.Now(),
			Coast: 53.4343,
		},

		Me: myEmployTestTest{
			Name: "Alexey",
			Age:  35,
			Done: true,
		},
	}

	jsn.Export(mySet, "export_settings.json")
	jsn.Print(mySet)
}

func jsonImportSample() {

	var settingsFromFile mySettingsTestTest
	_ = jsn.Import("import_settings.json", &settingsFromFile)
	fmt.Printf("type is: %T\n", settingsFromFile)

}

func timerSample() {
	fmt.Println()

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
