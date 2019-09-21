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
var display bool
var last int
var list *app.List

var rootCmd = &cobra.Command{
    Use:   "cor",
    Short: "Code Of Road !",
    Long: "Code Of Road : Une petite application en ligne de commande pour noter le nombre d'erreur au code de la route !",
    Run: func(cmd *cobra.Command, args []string) {
        if add != "none" {
            addAction()
        }
        
        if stat {
            statAction()
        }
        
        if display {
            displayAction()
        }

        if last > 0 {
            lastAction()
        }
    },
}

func init() {
    filename = "C:\\Users\\Lilian\\go\\src\\github.com\\LilianAvry\\code-of-road\\save.txt"
    handleNoFile()

    rootCmd.PersistentFlags().StringVarP(&add, "add", "a", "none", "Add new serie")
    rootCmd.PersistentFlags().BoolVarP(&stat, "stat", "s", false, "Display Statistics")
    rootCmd.PersistentFlags().BoolVarP(&display, "display", "d", false, "Display series")
    rootCmd.PersistentFlags().IntVarP(&last, "last", "l", 0, "Display x last series")

    initializeList()
}

func Execute() {
    rootCmd.Execute()
}

/*
 * Command Actions
 */

func addAction () {
    serie := fmt.Sprintf(";%s", add)
    appendFile(serie)
    
    fmt.Printf("La valeur %v a été enregistrée !\n", add)

    initializeList()
}

func statAction () {
    if list.Length() > 0 {
        statAll := list.StatAll()
        fmt.Printf("Moyenne totale : %v\n", statAll)   
    } else {
        fmt.Println("Aucune série n'a été enregistrée")
    }

    if list.Length() > 5 {
        statLast := list.StatLast()
        fmt.Printf("Moyenne des 5 dernières séries : %v\n", statLast)     
    }  
}

func displayAction () {
    fmt.Println("Affichage de toutes les séries enregistrées :")

    for i, serie := range list.Series {
        fmt.Printf(" - Série numéro %v : %v fautes\n", i + 1, serie)
    }
}

func lastAction () {
    if last == 1 {
        fmt.Println("Affichage de la dernière série enregistrée :")
    } else {
        fmt.Printf("Affichage des %v dernières séries enregistrées :\n", last)
    }

    start := list.Length() - last
    base := start + 1

    for i, serie := range list.Series[start:] {
        fmt.Printf(" - Série numéro %v : %v fautes\n", base + i, serie)
    }
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

func initializeList () {
    content := readFile()
    series := strings.Split(content, ";")
    list = app.NewList(series)
}

/*
 * Utils
 */

func handleError(err error) {
    if err != nil {
        panic(err)
    }
}

func handleNoFile() {
    _, err := os.Open(filename)
    if err != nil {
        os.Create(filename)
    }
}
