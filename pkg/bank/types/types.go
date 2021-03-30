package types
//Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д)
type Money int64
//Category представляет собой категорию, в которой был соверщён платёж (авто, аптеки рестороны и т.д.)
type Category string

//Payment предсавляет информатцию о платёже.
type Payment struct{
 ID int
 Amount Money
 Category Category
}

