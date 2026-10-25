package main 


import ("fmt"
        "math"
	        )

func main(){

	var nombre1 , nombre2 float64
	var operation string
	var resultat float64
	var choix string

	effectuerCalcul := true

for effectuerCalcul {

    calculeVrai := true

    fmt.Println("Entrer le premier nombre")
	fmt.Scan(&nombre1)
	fmt.Println("Enter le deuxième nombre")
	fmt.Scan(&nombre2)
	fmt.Println("Quelle opération voulez vous éffectuez (+,-,/,* ou %)")
	fmt.Scan(&operation)


	if operation == "+"{
		resultat = nombre1 + nombre2
	} else if operation == "-"{
		resultat =nombre1 - nombre2 
	} else if operation =="*" {
		resultat = nombre1 * nombre2
	} else if operation == "/"{
		if nombre2 != 0 {
			resultat = nombre1 / nombre2
		} else{
			calculeVrai = false
			fmt.Println("impossible de diviser un nombre par 0")
		}
	} else if operation =="%"{
		if nombre2!=0{
			resultat = math.Mod(nombre1,nombre2)
		} else{
			calculeVrai = false
			fmt.Println("impossible de diviser un nombre par 0")
		}
	} else{
		calculeVrai = false
		fmt.Println("choissisez entre +, /, * , % , -")
	}
	if calculeVrai == true {
		fmt.Println(resultat)
	    }
	fmt.Println("voulez vous continuer(oui/non)")
    fmt.Scan(&choix)
	if choix == "oui"{
		effectuerCalcul = true
	} else {
		effectuerCalcul = false
	    }
	}
}