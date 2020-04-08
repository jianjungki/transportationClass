
package railchina
type SearchRailChina struct{
  MerchantList []MerchantEntity
  FromStation string
  ToStation string
  ResultList []string
  Result string
}

func (s *SearchRailChina) Clear(){
  s.ResultList = nil
  s.Result = ""
}

func (s *SearchRailChina) NewSearch(merchants []MerchantEntity){
  //fmt.Printf("%v\n", merchants)
  s.MerchantList = merchants
}


func (s *SearchRailChina) Search(from string, to string){
  wg := sync.WaitGroup{}
  wg.Add(len(s.MerchantList))
  for _, m := range s.MerchantList{
    //fmt.Printf("%v\n", m)
    go func (m MerchantEntity){
      s.ResultList = append(s.ResultList, m.MerchantName+":  "+ m.Operator.Search(from, to))
      wg.Done()
    }(m)
  }
  wg.Wait()
}

func (s *SearchRailChina) MergeResult() {
  for _, res := range s.ResultList{
    //fmt.Printf("%v\n", res)
    s.Result += fmt.Sprintf("%s\n", res)
  }
}

func (s *SearchRailChina) ExportResult() string {
  //fmt.Printf("%v\n", s.MerchantList)
  //fmt.Printf("%v\n", s.ResultList)
  return s.Result
}

func (s *SearchRailChina) TrainInfo(f func (identify string, elm int) string, identify string, elm int) string{
  return f(identify, elm)
}

func (s *SearchRailChina) SearchOption(f func (from string, to string) string) string{
  return f(s.FromStation, s.ToStation)
}

func (s *SearchRailChina) StationSearch(keyword string) string{
  return fmt.Sprintf("find stations from keyword %s : xxxx", keyword)
}

