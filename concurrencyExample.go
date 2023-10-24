package main

import (
  "net/url"
  "encoding/json"
  "fmt"
  "net/http"
  "sync"
  "time"
)



func fetchUniversity(university string, ch chan<- string, wg *sync.WaitGroup) interface{} {
  var data []struct {
	State string `json:"state-province"`
	Country string `json:"country"`
	Domains []string `json:"country"`
	WebPages []string `json:"web_pages"`
	AlphaCode string `json:"alpha_two_code"`
	Name string `json:"name"`
  } 
  
  
  defer wg.Done() 
  urlQuery :=  url.QueryEscape(university)
  url := fmt.Sprintf("http://universities.hipolabs.com/search?name=%s",urlQuery)
  resp, err := http.Get(url)
  if err != nil {
    fmt.Printf("Error fetching university for %s: %s \n", university, err)
    return data
  }

  defer resp.Body.Close()

  if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
  fmt.Printf("Error decoding university data for %s: %s \n", university, err)
  return data
  }
  ch <- fmt.Sprintf("This is the %s", university)
  return data
}
  
func main() {
  startNow := time.Now()
  
  universities := []string{"Loyola Marymount University", "University of California, Los Angeles", 
  "Pomona College", "University of California, Riverside"}
  
  ch := make(chan string)
  var wg sync.WaitGroup

  for _, university := range universities {
    wg.Add(1)
    go fetchUniversity(university, ch, &wg)
  } 
  
  go func() {
    wg.Wait()
    close(ch)
  }()

  for result := range ch {
    fmt.Println(result)
  } 
  fmt.Println("This operation took:", time.Since(startNow))
}
