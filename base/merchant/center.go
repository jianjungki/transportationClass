package merchant
import "fmt"

type MerchantCenter struct{
  Entitys []MerchantEntity
  Search SearchOperator
  Order OrderOperator
}
//MerchantEntity 供应商中心
type MerchantEntity struct{
  Operator MerchantOperator //供应商对象
  MerchantName string //供应商名称
  MerchantSystemWeight int //供应商系统权重
  MerchantManualWeight int //供应商人工干预权重
  MerchantServiceFee int //供应商单价
}

//MerchantOperator 供应商操作对象
type MerchantOperator interface{
  Search(string, string) string //搜索
  PassengerVerify() string//验证乘客
  HoldSeat() string//占座
  MerchantCancel() string//供应商订单取消（确认订单前）
  ConfirmOrder() string//确认订单
  MerchantOrderDetail() string//供应商订单详情
  MerchantRefund() string//供应商退款（退单）
  Deliver() string//供应商邮寄服务
}
