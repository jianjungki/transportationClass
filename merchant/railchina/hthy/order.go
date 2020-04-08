package hthy

type HthyOrderOperator struct{
  Price string
  OrderStatus string
}

func (h *HthyOrderOperator) Confirm(){
  fmt.Println("confirming your order")
}

func (h *HthyOrderOperator) HoldSeat(){
  fmt.Println("your order is trying to holdseat")
}

func (h *HthyOrderOperator) Cancel(){
  fmt.Println("your order cancel is processing")
}

func (h *HthyOrderOperator) Refund(){
  fmt.Println("refunding your order")
}

func (h *HthyOrderOperator) Status(){
  fmt.Println("your order status is : %s", h.OrderStatus)
}

func (h *HthyOrderOperator) PriceDetail(){
  fmt.Println("your order price is : %s", h.Price)
}

func (h *HthyOrderOperator) TripInfo(){
  fmt.Println("your trip is from %s to %s, and your seat is 15coach 3A")
}
