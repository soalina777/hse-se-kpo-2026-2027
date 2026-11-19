package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"pos/internal/application"
	"pos/internal/domain"
	"strconv"
	"strings"
)

type CLI struct {
	facade *application.POSFacade
	reader *bufio.Reader
}

func NewCLI(facade *application.POSFacade) *CLI {
	return &CLI{
		facade: facade,
		reader: bufio.NewReader(os.Stdin),
	}
}

func (cli *CLI) ReadString(text string) (string, error) {
	fmt.Print(text)

	input, err := cli.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	if err == io.EOF && input == "" {
		return "", io.EOF
	}

	return strings.TrimSpace(input), nil
}

func (cli *CLI) ReadInt(text string) (int, error) {
	for {
		input, err := cli.ReadString(text)
		if err != nil {
			return 0, err
		}

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Введите целое число")
			continue
		}

		return value, nil
	}
}

func (cli *CLI) ReadFloat(text string) (float64, error) {
	for {
		input, err := cli.ReadString(text)
		if err != nil {
			return 0, err
		}

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Введите корректное число")
			continue
		}

		return value, nil
	}
}

func (cli *CLI) ShowOptions() {
	fmt.Println("=====POS=====")
	fmt.Println("1. Добавить студента")
	fmt.Println("2. Добавить критерий")
	fmt.Println("3. Показать студентов")
	fmt.Println("4. Показать критерии")
	fmt.Println("5. Выставить / изменить оценку")
	fmt.Println("6. Посмотреть оценки студента")
	fmt.Println("7. Рассчитать итоговую оценку")
	fmt.Println("8. Удалить оценку")
	fmt.Println("9. Экспорт результатов")
	fmt.Println("0. Выход")
}

func (cli *CLI) Run() {
	for {
		cli.ShowOptions()

		option, err := cli.ReadInt("Выберите действие: ")

		if err != nil {
			if err == io.EOF {
				fmt.Println("Завершение программы")
				return
			}

			fmt.Println("Ошибка чтения:", err)
			return
		}

		switch option {
		case 0:
			fmt.Println("Завершение программы")
			return
		case 1:
			cli.AddStudent()
		case 2:
			cli.AddCriterion()
		case 3:
			cli.ShowStudents()
		case 4:
			cli.ShowCriterions()
		case 5:
			cli.Evaluate()
		case 6:
			cli.ShowEvaluations()
		case 7:
			cli.CalculateResult()
		case 8:
			cli.DeleteEvaluation()
		case 9:
			cli.ExportResults()
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func (cli *CLI) AddStudent() {
	id, err := cli.ReadString("Введите ID студента: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	name, err := cli.ReadString("Введите имя студента: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	err = cli.facade.AddStudent(id, name)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Студент успешно добавлен!")
}

func (cli *CLI) AddCriterion() {
	id, err := cli.ReadString("Введите ID критерия: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	name, err := cli.ReadString("Введите название критерия: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	maxScore, err := cli.ReadInt("Введите максимальный балл: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	weight, err := cli.ReadFloat("Введите вес критерия: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if err := cli.facade.AddCriterion(id, name, maxScore, weight); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Критерий успешно добавлен!")
}

func (cli *CLI) ShowStudents() {
	students, err := cli.facade.GetStudents()

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if len(students) == 0 {
		fmt.Println("Список студентов пуст")
		return
	}

	fmt.Println("\n===== СПИСОК СТУДЕНТОВ =====")
	fmt.Printf("%-15s %-25s\n", "ID", "Имя")

	for _, student := range students {
		fmt.Printf(
			"%-15s %-25s\n",
			student.ID(),
			student.Name(),
		)
	}
}

func (cli *CLI) ShowCriterions() {
	criterions, err := cli.facade.GetCriterions()

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if len(criterions) == 0 {
		fmt.Println("Список критериев пуст")
		return
	}

	fmt.Println("\n===== СПИСОК КРИТЕРИЕВ =====")
	fmt.Printf(
		"%-12s %-25s %-15s %-10s\n",
		"ID",
		"Название",
		"Макс. балл",
		"Вес",
	)

	for _, criterion := range criterions {
		fmt.Printf(
			"%-12s %-25s %-15d %-10.2f\n",
			criterion.ID(),
			criterion.Name(),
			criterion.MaxScore(),
			criterion.Weight(),
		)
	}
}

func (cli *CLI) Evaluate() {
	studentID, err := cli.ReadString("Введите ID студента: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	criterionID, err := cli.ReadString("Введите ID критерия: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	score, err := cli.ReadFloat("Введите полученный балл: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	err = cli.facade.Evaluate(studentID, criterionID, score)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Оценка успешно сохранена или обновлена")
}

func (cli *CLI) ShowEvaluations() {
	studentID, err := cli.ReadString("Введите ID студента: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	evaluations, err := cli.facade.GetEvaluations(studentID)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if len(evaluations) == 0 {
		fmt.Println("У студента нет оценок")
		return
	}

	criterions, err := cli.facade.GetCriterions()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	criterionsByID := make(map[string]domain.Criterion)

	for _, criterion := range criterions {
		criterionsByID[criterion.ID()] = criterion
	}

	fmt.Println("\n===== ОЦЕНКИ СТУДЕНТА =====")
	fmt.Printf("%-30s %-10s\n", "Критерий", "Балл")

	for _, evaluation := range evaluations {
		criterion, ok := criterionsByID[evaluation.CriterionID()]
		if !ok {
			fmt.Printf(
				"%-30s %-10.2f\n",
				"Критерий не найден",
				evaluation.Score(),
			)
			continue
		}

		fmt.Printf(
			"%-30s %-10.2f\n",
			criterion.Name(),
			evaluation.Score(),
		)
	}
}

func (cli *CLI) CalculateResult() {
	studentID, err := cli.ReadString("Введите ID студента: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	totalScore, err := cli.facade.CalculateResult(studentID)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("\n===== ИТОГОВАЯ ОЦЕНКА =====")
	fmt.Printf("Студент ID: %s\n", studentID)
	fmt.Printf("Итоговый балл: %.2f / 10\n", totalScore)
}

func (cli *CLI) DeleteEvaluation() {
	studentID, err := cli.ReadString("Введите ID студента: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	criterionID, err := cli.ReadString("Введите ID критерия: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	err = cli.facade.DeleteEvaluation(studentID, criterionID)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Оценка успешно удалена!")
}

func (cli *CLI) ExportResults() {
	path, err := cli.ReadString("Введите путь к файлу: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	err = cli.facade.ExportResults(path)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результаты успешно экспортированы")
}
