package main
import "fmt"

func main(){
	stateCapital := make(map[string]string)
	fmt.Println("")
	fmt.Printf("Data Type of State Capital : %T\n\n",stateCapital)

	stateCapital["Maharashtra"] = "Mumbai"
	stateCapital["Punjab"] = "Amritsar"
	stateCapital["Rajasthan"] = "jaipur"
	stateCapital["Madhya Pradesh"] = "Bhopal"
	stateCapital["Tamil nadu"] = "Chennai"
	stateCapital["Gujrat"] = "Gandhinagar"
	stateCapital["Telangana"] = "Hyderbad"
	stateCapital["Orissa"] = "Bhbaneswar"
	stateCapital["Bihar"] = "Patna"
	stateCapital["West Bengal"] = "Kolkata"
	stateCapital["Goa"] = "Panji"

	for state, capital := range stateCapital{
		fmt.Printf("State : %s | Capital : %s\n", state, capital)
	}
}
