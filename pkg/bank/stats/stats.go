package stats

import "github.com/Shavqat1/bank10.1/pkg/bank/types"

//Avg рассчитивает среднюю сумму платежа.
func Avg(payments []types.Payment)types.Money{
sum:=int64(0)
incr:=int64(0)
for _, payment := range payments {
sum+=int64(payment.Amount)
incr+=1
}
result:=int64(sum/incr)
return types.Money(result)
}