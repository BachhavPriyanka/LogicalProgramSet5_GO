package main
import "fmt"

type Friend struct{
	name string
	mobNumber string
	alternateMobNum string
	address string
	city string

}

func main(){

	friend1 := Friend{"Priyanka","98657890", "34768889", "Township", "Delhi"}
	friend2 := Friend{"Rizwaan","987878765", "", "hitech", "Hyderbaad"}
	friend3 := Friend{"Sufiyan","434456567", "34768889", "ojhar", "Nashik"}
	friend4 := Friend{"Mac","78664799", "34768889", "laxmiRoad", "Pune"}
	friend5 := Friend{"Ken","987656649", "", "College-Road", "Nagpur"}
	friend6 := Friend{"Jenny","999556789", "", "Airforce", "Pune"}
	friend7 := Friend{"Diksha","964336777", "34768889", "camp", "Mumbai"}
	friend8 := Friend{"Tom","789089995", "", "fish-Maeket", "Kokan"}
	friend9 := Friend{"henrry","078978900", "34768889", "camp2", "lonawalaa"}
	friend10 := Friend{"Albert","098080786", "34768889", "Hinjawadi", "Jaipur"}

	Details := []Friend{
		friend1,
		friend2,
		friend3,
		friend4,
		friend5,
		friend6,
		friend7,
		friend8,
		friend9,
		friend10,
	}
	for _,friendDetails := range Details{
		fmt.Println(friendDetails.name,"|", friendDetails.mobNumber,"|",friendDetails.alternateMobNum,"|",friendDetails.address,"|",friendDetails.city)
	}
	fmt.Println("")
	fmt.Println("Printing Name and Mobile numbers")
	fmt.Println("")
	for _,friendDetails := range Details{
		fmt.Println(friendDetails.name,"|", friendDetails.mobNumber)
	}

	fmt.Println("")
	fmt.Println("Printing names Who are having Alternate Number")
	for _,friendDetails := range Details{
		if(friendDetails.alternateMobNum != "" ){
			fmt.Println(friendDetails.name)
		}
	}
	

	// xyz := make(map[string]string)
	// for _,friendDetails := range Details{
	// 	 xyz[friendDetails.city] = "xyz "
	// 	 fmt.Println(xyz)
	// } 

	


}