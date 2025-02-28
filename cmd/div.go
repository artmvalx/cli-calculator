/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// divCmd represents the div command
var divCmd = &cobra.Command{
	Use:   "div [num1] [num2]",
	Short: "Деление первого числа на второе",
	Args:  cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		num1, err1 := strconv.ParseFloat(args[0], 64)
		num2, err2 := strconv.ParseFloat(args[1], 64)
		if err1 != nil || err2 != nil {
			fmt.Println("ОШИБКА: оба аргумента должны быть числами.")
			return
		}
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль.")
			return
		}
		result := num1 / num2
		fmt.Printf("Результат: %.2f\n", result)
	},
}

func init() {
	rootCmd.AddCommand(divCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// divCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
