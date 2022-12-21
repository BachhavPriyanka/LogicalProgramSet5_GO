package main
import (
	"fmt"
	"sort"
)

func main(){
	var states []string

	stateCapital := make(map[string]string)

	stateCapital["Maharashtra"] = "Mumbai"
	stateCapital["Punjab"] = "Amritsar"
	stateCapital["Rajasthan"] = "jaipur"
	stateCapital["Madhya Pradesh"] = "Bhopal"
	stateCapital["Tamil nadu"] = "Chennai"
	stateCapital["Gujrat"] = "Gandhinagar"
	stateCapital["Orissa"] = "Bhubaneswar"
	stateCapital["Bihar"] = "Patna"
	stateCapital["West Bengal"] = "Kolkata"
	stateCapital["Goa"] = "Panji"

	for state, capital := range stateCapital{
		fmt.Printf("State : %s | Capital : %s\n", state, capital)
	}

	for eachState := range stateCapital{
		states = append(states, eachState)
	}

	sort.Strings(states)
	fmt.Println("Sorted States : ",states)

}

