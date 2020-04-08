package main
import (
  "fmt"
)

//OrderBase 订单基类
type OrderBase struct{
  OrderStatus int
  TotalPrice string
  PaymentPrice string
  CouponPrice string
}

//SubOrderEntity 子订单实体
type SubOrderEntity struct{
  OrderBase
  Merchant MerchantEntity 
  FromStation string
  ToStation string
}
//OrderEntity 订单实体
type OrderEntity struct{
  SubOrder []SubOrderEntity
  OrderBase
}

//Booking 订单中心
type Booking struct{
  Book BookingOperator //预定对象
  Payment PaymentOperator //支付对象
  Info  InfoOperator //信息对象
  
  //Public
  Order OrderEntity //订单对象
}

//BookingOperator 预定操作
type BookingOperator interface{
  AddShoppingCar(string, string, MerchantEntity) SubOrderEntity//加入购物车
  ShowShoppingCar(OrderEntity)  //结算购物车
  ShowShoppingCarUpdate(OrderEntity) //更新购物车
  ShowOrder(string) //订单结算
  ShowOrderUpdate(string) //更新订单结算
  Cancel(string) //订单取消
  Expire(string) //订单超时取消
  StockOut(string) //出库
}


//InfoOperator 信息操作
type InfoOperator interface{
  OrderInfo(string) //订单详情
  ProductInfo(string) //商品详情
  PaymentInfo(string) //支付详情
  StockOutInfo(string) //出库详情
  OperatorInfo(string) //操作详情
}

