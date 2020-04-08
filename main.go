package main
import "fmt"
var center Booking
func main(){
  initEntity()
  
  orderSequence()
  //joinWay()
}

func initEntity(){
  //BookingOperator
  book := &KlookBooking{}
  info := &KlookInfo{}
  payment := &KlookPayment{}
  center = NewBooking(book, info, payment)
}

func NewBooking(book BookingOperator, info InfoOperator, payment PaymentOperator) Booking{
  return Booking{
    Book: book,
    Payment: payment,
    Info: info,
  }
}

func orderSequence(){
  hthy := MerchantCenter{
    Entitys: []MerchantEntity{
      NewMerchantHthy(),
      NewMerchantCtrip(),
    },
    Order: &HthyOrderOperator{},
  }
  hthy.Search = &SearchRailChina{
      MerchantList: hthy.Entitys,
  }

  hthy.Search.Search("深圳", "上海")
  hthy.Search.MergeResult()
  fmt.Println(hthy.Search.ExportResult())

  center.Order.SubOrder = append(center.Order.SubOrder, center.Book.AddShoppingCar("深圳", "上海", hthy.Entitys[0]))
  center.Book.ShowShoppingCar(center.Order)

  hthy.Search.Clear()

  hthy.Search.MergeResult()
  hthy.Search.Search("上海", "北京")
  hthy.Search.MergeResult()

  fmt.Println(hthy.Search.ExportResult())

  center.Order.SubOrder = append(center.Order.SubOrder, center.Book.AddShoppingCar("上海", "北京", hthy.Entitys[1]))
  center.Book.ShowShoppingCar(center.Order)
  fmt.Println("")
  center.Book.ShowShoppingCarUpdate(center.Order)

  center.Payment
}

func NewMerchantHthy() MerchantEntity{
  merchant := &MerchantHthy{}
  merchantObj := MerchantEntity{
    Operator: merchant,
    MerchantName: "hthy",
    MerchantSystemWeight: 2, 
    MerchantManualWeight: 10,
    MerchantServiceFee: 2,
  }
  return merchantObj
}

func NewMerchantCtrip() MerchantEntity{
  merchant := &MerchantHthy{}
  merchantObj := MerchantEntity{
    Operator: merchant,
    MerchantName: "ctrip",
    MerchantSystemWeight: 2, 
    MerchantManualWeight: 10,
    MerchantServiceFee: 2,
  }
  return merchantObj
}

func joinWay(){
  j := JointWayOperator{}
  merchant := &MerchantHthy{
    FromStation: "深圳",
    ToStation: "上海",
  }
  merchantObj := MerchantEntity{
    Operator: merchant,
    MerchantName: "hthy",
    MerchantSystemWeight: 2, 
    MerchantManualWeight: 10,
    MerchantServiceFee: 2,
  }
  j.AddTrip(merchantObj)


  merchant = &MerchantHthy{
    FromStation: "上海",
    ToStation: "天津",
  }

  merchantNew := MerchantEntity{
    Operator: merchant,
    MerchantName: "ctrip",
    MerchantSystemWeight: 2, 
    MerchantManualWeight: 10,
    MerchantServiceFee: 2,
  }
  j.AddTrip(merchantNew)
  j.TripMerge()
}

//JointWayOperator 联程操作对象
type JointWayOperator struct{
  Trips []MerchantEntity //多个供应商的行程
}

//TripMerge 合并行程
func (o *JointWayOperator) TripMerge() *Booking{
  fmt.Println("\nMerge Order:")
  for _, trip := range o.Trips{
    fmt.Printf("%s,", trip.MerchantName + " "+ trip.Operator.MerchantOrderDetail())
  }
  return nil
}

//AddTrip 添加行程
func (o *JointWayOperator) AddTrip(merchant MerchantEntity) {
  fmt.Println("name:"+merchant.MerchantName+"\tdetail:"+ merchant.MerchantName+" "+ merchant.Operator.HoldSeat())
  o.Trips = append(o.Trips, merchant)
}

//ExcisionOperator 拆单操作对象
type ExcisionOperator struct{
  FromStation string
  ToStation string
  TrainNo string
  //SplitTrips []MerchantRoute
}