package cmd

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
    "github.com/spf13/cobra"
    "github.com/LilianAvry/code-of-road/app"
)

var filename string
var add string
var stat bool

var rootCmd = &cobra.Command{
    Use:   "hello",
    Short: "Hello !",
    Long: "Hello this is my command !",
    Run: func(cmd *cobra.Command, args []string) {
        if add != "none" {
                addAction(add)
        }
        if stat {
                statAction()
        }
    },
}

func init() {
    filename = "C:\\Users\\Lilian\\go\\src\\github.com\\LilianAvry\\code-of-road\\save.txt"
    rootCmd.PersistentFlags().StringVarP(&add, "add", "a", "none", "Add output")
    rootCmd.PersistentFlags().BoolVarP(&stat, "stat", "s", false, "Stat output")
}

func Execute() {
    rootCmd.Execute()
}

/*
 * Command Actions
 */

func addAction (value string) {
        serie := fmt.Sprintf(";%s", value)
        appendFile(serie)
        fmt.Printf("La valeur %v a été enregistrée !\n", value)
}

func statAction () {
        content := readFile()

        series := strings.Split(content, ";")
        list := app.NewList(series)

        statAll := list.StatAll()
        statLast := list.StatLast()

        fmt.Printf("Moyenne totale : %v\n", statAll)
        fmt.Printf("Moyenne des 5 dernières séries : %v\n", statLast)
}


/*
 * File management
 */

func appendFile(content string) {
        // Open a file
        file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_RDONLY, 0600)
        handleError(err)
        defer file.Close()

	// Append to a file
	_, err = file.WriteString(content)
	handleError(err)
}

func readFile () string {
	// Read a file
        content, err := ioutil.ReadFile(filename)
	handleError(err)
	return string(content)
}

/*
 * Utils
 */

func handleError(e error) {
    if e != nil {
        panic(e)
    }
}
