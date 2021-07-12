package weightconv

func KgToLbs(kg Kilograms) Pounds{
	return Pounds(kg * 2.2)
}

func LbsToKg(lb Pounds) Kilograms{
	return Kilograms(lb / 2.2)
}
