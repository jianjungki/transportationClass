package klook

type KlookBooking struct{}

func (c *KlookBooking) AddShoppingCar(from string, to string, merchant MerchantEntity) SubOrderEntity{
  fmt.Printf("add shoppingcar ticket : %s from %s to %s \n", merchant.MerchantName, from, to) 

  return SubOrderEntity{
    Merchant: merchant,
    FromStation: from,
    ToStation: to,
  }
}

func (c *KlookBooking) ShowShoppingCar(order OrderEntity) {
  fmt.Println("showing shoppingcar:")
  for _, item := range order.SubOrder{
    fmt.Printf("%s from %s to %s, guid: xxxx\n", item.Merchant.MerchantName, item.FromStation, item.ToStation) 
  }
  
}

func (c *KlookBooking) ShowShoppingCarUpdate(order OrderEntity) {
  fmt.Println("showing shoppingcar update:")
  for _, item := range order.SubOrder{
    fmt.Printf("%s from %s to %s, guid: xxxx\n", item.Merchant.MerchantName, item.FromStation, item.ToStation) 
  }
}

func (c *KlookBooking) ShowOrder(orderGuid string) {
  fmt.Printf("showing order, guid: %s\n", orderGuid) 
}

func (c *KlookBooking) ShowOrderUpdate(orderGuid string) {
  fmt.Printf("order info update, guid: %s\n", orderGuid) 
}

func (c *KlookBooking) Cancel(orderGuid string) {
  fmt.Printf("order cancel, guid: %s\n", orderGuid) 
}

func (c *KlookBooking) Expire(orderGuid string) {
  fmt.Printf("order expire, guid: %s\n", orderGuid) 
}
func (c *KlookBooking) StockOut(orderGuid string) {
  fmt.Printf("order stockout, guid: %s\n", orderGuid) 
}