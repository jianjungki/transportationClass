package platform

type PirceCenter interface{
  //从各种参数来获取价格
  PriceDiscount(f func (args ...interface{}), args ...interface{})
}