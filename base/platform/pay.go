package platform
//PaymentOperator 支付操作
type PaymentOperator interface{
  Pay(string) //支付
  PaymentStatus(string) //支付状态
  Refund(string) //发起退款
  GetDetail(string) //支付详情
  ManualRefund(string) //手动退款（特殊退款）
}
