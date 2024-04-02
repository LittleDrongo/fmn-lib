package errors

type Reporter interface {
	Report(error, ...interface{})
}

func Report(err error, message ...interface{}) {
	if err != nil {
		//TODO Добавить метод который будет отправлять отчёт об ошибке на почту
	}
}
