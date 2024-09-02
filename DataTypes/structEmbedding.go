// it is very similar to nested structs where an attrubute of a struct has the structure of another strucut
// it is like a nested oject in javascript
package main;
import "fmt";


// In this example, the driver struct embeds the car struct. This means that driver has all the fields of car, in addition to its own fields name and age.
type car struct {
	model string
	licencePlate string
}

type driver struct {
	name string
	age int
	car
}

func (carMethod car) honk() string{
	return "Beep!"
}

func main(){

	var driver1 driver = driver{
		name: "Salman Khan",
		age:53,
		car: car{
			model:"AP010001",
			licencePlate: "GH02J0P9",
		},
	}


	fmt.Println(driver1) //{Salman Khan 53 {AP010001 GH02J0P9}};
	fmt.Println(driver1.car.model);
	fmt.Println(driver1.car.honk());




}