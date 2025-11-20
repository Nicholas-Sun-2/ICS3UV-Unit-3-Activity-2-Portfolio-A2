/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-20
 * @fileoverview This program prompts the user to enter the names of the ten Canadian provinces, three territories, and their capital cities.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	// Declaring variables
	var province1 string
	var capital1 string
	var province2 string
	var capital2 string
	var province3 string
	var capital3 string
	var province4 string
	var capital4 string
	var province5 string
	var capital5 string
	var province6 string
	var capital6 string
	var province7 string
	var capital7 string
	var province8 string
	var capital8 string
	var province9 string
	var capital9 string
	var province10 string
	var capital10 string
	var province11 string
	var capital11 string
	var province12 string
	var capital12 string
	var province13 string
	var capital13 string

	// Input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Province/Territory 1: ")
	province1, _ = reader.ReadString('\n')
	province1 = strings.TrimSpace(province1)
	fmt.Print("What is the capital of " + province1 + "? ")
	capital1, _ = reader.ReadString('\n')
	capital1 = strings.TrimSpace(capital1)

	fmt.Print("Province/Territory 2: ")
	province2, _ = reader.ReadString('\n')
	province2 = strings.TrimSpace(province2)
	fmt.Print("What is the capital of " + province2 + "? ")
	capital2, _ = reader.ReadString('\n')
	capital2 = strings.TrimSpace(capital2)

	fmt.Print("Province/Territory 3: ")
	province3, _ = reader.ReadString('\n')
	province3 = strings.TrimSpace(province3)
	fmt.Print("What is the capital of " + province3 + "? ")
	capital3, _ = reader.ReadString('\n')
	capital3 = strings.TrimSpace(capital3)

	fmt.Print("Province/Territory 4: ")
	province4, _ = reader.ReadString('\n')
	province4 = strings.TrimSpace(province4)
	fmt.Print("What is the capital of " + province4 + "? ")
	capital4, _ = reader.ReadString('\n')
	capital4 = strings.TrimSpace(capital4)

	fmt.Print("Province/Territory 5: ")
	province5, _ = reader.ReadString('\n')
	province5 = strings.TrimSpace(province5)
	fmt.Print("What is the capital of " + province5 + "? ")
	capital5, _ = reader.ReadString('\n')
	capital5 = strings.TrimSpace(capital5)

	fmt.Print("Province/Territory 6: ")
	province6, _ = reader.ReadString('\n')
	province6 = strings.TrimSpace(province6)
	fmt.Print("What is the capital of " + province6 + "? ")
	capital6, _ = reader.ReadString('\n')
	capital6 = strings.TrimSpace(capital6)

	fmt.Print("Province/Territory 7: ")
	province7, _ = reader.ReadString('\n')
	province7 = strings.TrimSpace(province7)
	fmt.Print("What is the capital of " + province7 + "? ")
	capital7, _ = reader.ReadString('\n')
	capital7 = strings.TrimSpace(capital7)

	fmt.Print("Province/Territory 8: ")
	province8, _ = reader.ReadString('\n')
	province8 = strings.TrimSpace(province8)
	fmt.Print("What is the capital of " + province8 + "? ")
	capital8, _ = reader.ReadString('\n')
	capital8 = strings.TrimSpace(capital8)

	fmt.Print("Province/Territory 9: ")
	province9, _ = reader.ReadString('\n')
	province9 = strings.TrimSpace(province9)
	fmt.Print("What is the capital of " + province9 + "? ")
	capital9, _ = reader.ReadString('\n')
	capital9 = strings.TrimSpace(capital9)

	fmt.Print("Province/Territory 10: ")
	province10, _ = reader.ReadString('\n')
	province10 = strings.TrimSpace(province10)
	fmt.Print("What is the capital of " + province10 + "? ")
	capital10, _ = reader.ReadString('\n')
	capital10 = strings.TrimSpace(capital10)

	fmt.Print("Province/Territory 11: ")
	province11, _ = reader.ReadString('\n')
	province11 = strings.TrimSpace(province11)
	fmt.Print("What is the capital of " + province11 + "? ")
	capital11, _ = reader.ReadString('\n')
	capital11 = strings.TrimSpace(capital11)

	fmt.Print("Province/Territory 12: ")
	province12, _ = reader.ReadString('\n')
	province12 = strings.TrimSpace(province12)
	fmt.Print("What is the capital of " + province12 + "? ")
	capital12, _ = reader.ReadString('\n')
	capital12 = strings.TrimSpace(capital12)

	fmt.Print("Province/Territory 13: ")
	province13, _ = reader.ReadString('\n')
	province13 = strings.TrimSpace(province13)
	fmt.Print("What is the capital of " + province13 + "? ")
	capital13, _ = reader.ReadString('\n')
	capital13 = strings.TrimSpace(capital13)

	// Printing
	fmt.Println("Province/Territory    Captial");
	fmt.Println(province1 + "    " + capital1);
	fmt.Println(province2 + "    " + capital2);
	fmt.Println(province3 + "    " + capital3);
	fmt.Println(province4 + "    " + capital4);
	fmt.Println(province5 + "    " + capital5);
	fmt.Println(province6 + "    " + capital6);
	fmt.Println(province7 + "    " + capital7);
	fmt.Println(province8 + "    " + capital8);
	fmt.Println(province9 + "    " + capital9);
	fmt.Println(province10 + "    " + capital10);
	fmt.Println(province11 + "    " + capital11);
	fmt.Println(province12 + "    " + capital12);
	fmt.Println(province13 + "    " + capital13);
}
