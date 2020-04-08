package base
import "fmt"
type OrderOperator interface{
  Confirm()
  HoldSeat()
  Cancel()
  Refund()

  Status()
  PriceDetail()
  TripInfo()
}
