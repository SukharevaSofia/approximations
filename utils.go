package main

const (
	INFO = "\n***\nСухарева Софья \nP32131, ИСУ:334641" +
		"\nАППРОКСИМАЦИЯ ФУНКЦИИ МЕТОДОМ\nНАИМЕНЬШИХ КВАДРАТОВ***\n"
	CHOOSE_INPUT = "Выберите способ ввода: \"F\" (файл) или \"T\" (терминал)." +
		"\nF/T: "
	NUMBER_OF_DOTS = "Количество точек должно быть целым числом в пределах [7, 12]." +
		"\nКоличество точек: "
	INPUT_X       = "Введите значения X через пробел: "
	INPUT_Y       = "Введите значения Y через пробел: "
	INPUT_ERR     = "Данные введены неправильно. Повторите ввод."
	CHOOSE_OUTPUT = "Выберите способ вывода: \"F\" (файл) или \"T\" (терминал)." +
		"\nF/T: "
	MEAN_SQ_D   = "Среднеквадратичное отклонение: "
	BEST_APPROX = "Лучшее приближение : "
	LINEAR      = "Линейная аппроксимация"
	QUADRATIC   = "Квадратичная аппроксимация"
	CUBIC       = "Кубическая аппроксимация"
	EXP         = "Экспоненциальная аппроксимация"
	LOG         = "Логарифмическая аппроксимация"
	POWER       = "Степенная аппроксимация"
)

type Data struct {
	fuction             func(element float64) float64
	coefficents         []float64
	meanSquareDeviation float64
	xData               []float64
	yData               []float64
	phi                 []float64
	eps                 []float64
	nameOfApprox        string
}
