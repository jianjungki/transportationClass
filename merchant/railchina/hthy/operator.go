
package hthy
type MerchantHthy struct{
  FromStation string
  ToStation string
}

func (m *MerchantHthy) Search(from string, to string) string{
  m.FromStation = from
  m.ToStation = to
  return fmt.Sprintf("search from %s to %s", from, to)
}

func (m *MerchantHthy) SearchWithOutPrice(from string, to string) string{
  return fmt.Sprintf("search from %s to %s without price", from, to)
}

func (m *MerchantHthy) Deliver() string{
  return "merchant deliver service"
}


func (m *MerchantHthy) PassengerVerify() string{
  return "passenger verify"
}

func (m *MerchantHthy) HoldSeat() string{
  return fmt.Sprintf("success hold seat: %s to %s", m.FromStation , m.ToStation)
}

func (m *MerchantHthy) MerchantCancel() string{
  return "order cancel"
}

func (m *MerchantHthy) ConfirmOrder() string{
  return "order doing confirm "
}

func (m *MerchantHthy) MerchantOrderDetail() string{
  return fmt.Sprintf("order is 180 yuan, %s to %s ", m.FromStation, m.ToStation)
}

func (m *MerchantHthy) MerchantRefund() string{
  return "order refund apply"
}