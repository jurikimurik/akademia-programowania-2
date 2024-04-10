package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

var data = map[string]any{}

func main() {
	lFlag := flag.String("load", "", "Name of JSON file to load")
	flag.Parse()
	if *lFlag != "" {
		LoadJSON(*lFlag)
	}
	// Petla nadrzedna
	for {
		// Petla podrzedna
		var answer string = ""
		for {
			showMainMenu()
			answer = readInput("Podaj akcje:")

			if answer == "" {
				continue
			} else {
				break
			}
		}

		switch {
		case answer == "dodaj":
			AddParameter()
		case answer == "usun":
			DeleteParameter()
		case answer == "modyfikacja":
			ModParameter()
		case answer == "zapisz":
			SaveJSON()
		case answer == "zaladuj":
			LoadJSON("")
		case answer == "wyjdz":
			os.Exit(3)
		default:
			fmt.Println("Akcja niezdefiniowana. Powtorz.")
		}
	}

}

func showMainMenu() {
	fmt.Println("\n")
	showData()
	fmt.Println("\n")
	fmt.Println("Dodawanie parametru - dodaj\nUsuwanie parametru - usun\nModyfikacja parametru - modyfikacja")
	fmt.Println("Zapisz JSON do pliku - zapisz\nZaladuj JSON z pliku - zaladuj\nWyjdz z aplikacji - wyjdz")
}

func showData() {
	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}

func readInput(text string) string {
	fmt.Print(text)
	var answer string
	fmt.Scanln(&answer)
	return answer

}

func LoadJSON(toOpen string) {
	var newName string
	if toOpen == "" {
		newName = readInput("\nPodaj nazwe pliku (bez rozszerzenia):")
		newName += ".json"
	} else {
		newName = toOpen
	}

	fileContent, err := os.ReadFile(newName)
	if err != nil {
		fmt.Println("Nie moge odczytac tego pliku! Kurczaki!")
		return
	}

	text := string(fileContent)

	if err := json.Unmarshal([]byte(text), &data); err != nil {
		panic(err)
	}

	fmt.Println("Zaladowano JSON!")
}

func SaveJSON() {
	newName := readInput("\nPodaj nazwe pliku (bez rozszerzenia):")
	newName += ".json"

	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile(newName, file, 0664)
	fmt.Println("Zapisano JSON!")
}

func ModParameter() {
	parameterToModify := readInput("\nPodaj parameter do modyfikacji:")

	_, ok := data[parameterToModify]
	if !ok {
		fmt.Println("Parameter nie istnieje!")
		return
	}

	newDescription := readInput("Podaj znaczenie:")
	data[parameterToModify] = newDescription
	DataWasModificated()
}

func DataWasModificated() {
	data["Ostatnia modyfikacja"] = time.Now()
}

func DeleteParameter() {
	parameterToDelete := readInput("\nPodaj parameter do usuniecia:")

	_, ok := data[parameterToDelete]
	if !ok {
		fmt.Println("Parameter nie istnieje!")
		return
	}

	delete(data, parameterToDelete)
	DataWasModificated()
}

func AddParameter() {
	newParameter := readInput("\nPodaj parameter:")
	newDescription := readInput("Podaj znaczenie:")

	_, ok := data[newParameter]
	if ok {
		fmt.Println("Parameter juz istnieje!")
		return
	}

	data[newParameter] = newDescription
	DataWasModificated()
}
