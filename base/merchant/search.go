package merchant
import (
  "fmt"
)

type SearchOperator interface{
  NewSearch([]MerchantEntity)
  Search(from string, to string)
  SearchOption(func (from string, to string) string) string
  StationSearch(keyword string) string
  TrainInfo(func (string, int) string, string, int) string
  MergeResult() 
  ExportResult() string
  Clear()
} 