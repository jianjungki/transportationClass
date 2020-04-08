package klook
//KlookInfo 操作
type KlookInfo struct{}


func (p *KlookInfo) OrderInfo(orderGuid string) {
  fmt.Printf("order info, guid: %s\n", orderGuid) 
}

func (p *KlookInfo) ProductInfo(orderGuid string) {
  fmt.Printf("order product info , guid: %s\n", orderGuid) 
}

func (p *KlookInfo) PaymentInfo(orderGuid string) {
  fmt.Printf("order payment info, guid: %s\n", orderGuid) 
}

func (p *KlookInfo) StockOutInfo(orderGuid string) {
  fmt.Printf("order stockout info, guid: %s\n", orderGuid) 
}

func (p *KlookInfo) OperatorInfo(orderGuid string) {
  fmt.Printf("order operator info, guid: %s\n", orderGuid) 
}
