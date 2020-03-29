package Covid19

import (
	"fmt"
)

// Information Displaying information about Covid19 cirus
func Information() {
	var input string
	status1, status2, status3, status4 := 0, 0, 0, 0
	count := 0
	status := true
	for status == true {
		display := " Please select an option: \n 1 – Print Covid19 cases in Pakistan \n 2 – Print Covid19 cases in SouthKorea \n " +
			"3 – Print Covid19 cases in France \n 4 – Print my personalized message to Coronavirus \n 0 – Exit "
		fmt.Println(display)
		fmt.Scanf("%s", &input)
		fmt.Println("Input", input)

		switch input {
		case "1":

			fmt.Println("***\nPrinting Covid19 cases in Pakistan: 1,526\n***")
			if status1 == 0 {
				count = count + 1
				status1 = 1
			}

		case "2":
			fmt.Println("***\nPrinting Covid19 cases in SouthKorea: 9,583\n***")
			if status2 == 0 {
				count = count + 1
				status2 = 1
			}
		case "3":
			fmt.Println("***\n Printing Covid19 cases in France: 37,575 \n***")
			if status3 == 0 {
				count = count + 1
				status3 = 1
			}
		case "4":
			fmt.Println("***\nThe disease can spread from person to person through small droplets from the nose or mouth" +
				"\nwhich are spread when a person with COVID-19 coughs or exhales.\n*** ")
			if status4 == 0 {
				count = count + 1
				status4 = 1
			}
		case "0":
			if count >= 4 {
				fmt.Println("***\nExiting\n***")
				status = false
			} else {
				fmt.Println("***\nCould not exit, all options have not explored yet\n***")
			}

		default:
			fmt.Printf("***\nEnter correct input format e.g. Enter 1 \n***")
		}

	}

}
