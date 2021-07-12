//this program converts from kilograms to pounds and vice-versa
package weightconv

import (
	"fmt"
)

type Kilograms float64
type Pounds float64

func (kg Kilograms) String() string {
	return fmt.Sprintf("%g kg", kg)	
}
func (lb Pounds) String() string {
	return fmt.Sprintf("%g lbs", lb)	
}

