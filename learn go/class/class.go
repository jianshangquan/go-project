package class


type Engine struct{
	EngineType string
	HorsePower int
	Torque int
}


type Car struct{
	Brand string
	Model string
	Year int
	Price float64
	EngineType *Engine
}



func (car *Car) Run(){
	println("Car is running")
	println("Brand: ", car.Brand)
	println("Model: ", car.Model)
	println("Year: ", car.Year)
	println("Price: ", car.Price)
	println("Engine Type: ", car.EngineType.EngineType)
	println("Horse Power: ", car.EngineType.HorsePower)
	println("Torque: ", car.EngineType.Torque)
}

func (car *Car) Drive(){
	println("Car is driving")
}







func NewCar() *Car{
	return &Car{
		Brand: "Toyota",
		Model: "Corolla",
		Year: 2020,
		Price: 20000.0,
		EngineType: &Engine{
			EngineType: "V6",
			HorsePower: 300,
			Torque: 400,
		},
	}
}





