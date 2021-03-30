package stats

import (
	"github.com/Shavqat1/bank10.1/pkg/bank/types"
    "fmt"
)
func ExampleAvg() {
	payments := []types.Payment{
	{	
		Amount:6_000_00,
	},
	{	
		Amount:15_000_00,
	},
	{	
		Amount:15_000_00,
	},
	{	
		Amount:16_000_00,
	},
	
}
	fmt.Println(Avg(payments))
	//Output:1300000 
}