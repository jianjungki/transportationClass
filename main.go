package main

import "fmt"

func main() {
  fmt.Println("Hello World")
}

//OrderCenter 订单中心
type OrderCenter struct{
  Book BookingOperator //预定对象
  Refund RefundOperator //退款对象
  Info  InfoOperator //信息对象
}
//MerchantCenter 供应商中心
type MerchantCenter struct{
  Merchant MerchantOperator //供应商对象
  MerchantName string //供应商名称
  MerchantSystemWeight int //供应商系统权重
  MerchantManualWeight int //供应商人工干预权重
  MerchantServiceFee int //供应商单价
}
//MerchantOperator 供应商操作对象
type MerchantOperator interface{
  PassengerVerify() //验证乘客
  HoldSeat() //占座
  MerchantCancel() //供应商订单取消（确认订单前）
  ConfirmOrder() //确认订单
  MerchantOrderDetail() //供应商订单详情
  MerchantRefund() //供应商退款（退单）
}
//JointWayOperator 联程操作对象
type JointWayOperator struct{
  AddTrip(MerchantCenter) //添加行程
  TripMerge() //合并多个行程
  Trips []MerchantCenter //多个供应商合起来的行程
}

//BookingOperator 预定操作
type BookingOperator interface{
  AddShoppingCar() //加入购物车
  ShowShoppingCar()  //结算购物车
  ShowShoppingCarUpdate() //更新购物车
  ShowOrder() //订单结算
  ShowOrderUpdate() //更新订单结算
  Pay() //支付
  Cancel() //订单取消
  Expire() //订单超时取消
  StockOut() //出库
}
//RefundOperator 退款操作
type RefundOperator interface{
  Refund() //发起退款
  GetDetail() //退款详情
  Status() //退款状态
  ManualRefund() //手动退款（特殊退款）
}
//InfoOperator 信息操作
type InfoOperator interface{
  OrderInfo() //订单详情
  ProductInfo() //商品详情
  PaymentInfo() //支付详情
  RefundInfo() //退款详情
  StockOutInfo() //出库详情
  OperatorInfo() //操作详情
}


