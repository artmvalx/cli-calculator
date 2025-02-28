package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

// evalCmd представляет команду eval
var evalCmd = &cobra.Command{
	Use:   "eval [expression]",
	Short: "Вычисляет математическое выражение",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		expression := args[0]
		result, err := evaluateExpression(expression)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Printf("Результат: %.2f\n", result)
	},
}

func init() {
	rootCmd.AddCommand(evalCmd)
}

func evaluateExpression(expression string) (float64, error) {
	tokens := tokenize(expression)
	postfix, err := infixToPostfix(tokens)
	if err != nil {
		return 0, err
	}
	return evaluatePostfix(postfix)
}

func tokenize(expression string) []string {
	var tokens []string
	var number strings.Builder

	for _, char := range expression {
		if unicode.IsDigit(char) || char == '.' {
			number.WriteRune(char)
		} else if char == ' ' {
			continue
		} else {
			if number.Len() > 0 {
				tokens = append(tokens, number.String())
				number.Reset()
			}
			tokens = append(tokens, string(char))
		}
	}

	if number.Len() > 0 {
		tokens = append(tokens, number.String())
	}

	return tokens
}

func infixToPostfix(tokens []string) ([]string, error) {
	var postfix []string
	var stack []string
	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	for _, token := range tokens {
		if isNumber(token) {
			postfix = append(postfix, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, fmt.Errorf("несоответствие скобок")
			}
			stack = stack[:len(stack)-1] // удаляем "("
		} else if isOperator(token) {
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else {
			return nil, fmt.Errorf("неизвестный токен: %s", token)
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, fmt.Errorf("несоответствие скобок")
		}
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix, nil
}

func evaluatePostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if isNumber(token) {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("не удалось преобразовать %s в число", token)
			}
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, fmt.Errorf("недостаточно операндов для операции %s", token)
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("деление на ноль")
				}
				result = a / b
			default:
				return 0, fmt.Errorf("неизвестный оператор: %s", token)
			}
			stack = append(stack, result)
		} else {
			return 0, fmt.Errorf("неизвестный токен: %s", token)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("ошибка в выражении")
	}

	return stack[0], nil
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "%" || token == "^"
}
