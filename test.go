package main

import(
    "fmt"
    "time"
    "strings"

)


// utc to terrestrial time function
func tTime(lst [][]string, userDateTime time.Time)(float64){
  // this is a variable used for the UTC to TT conversion
  seconds := 32.184

  // next we check to make sure that the request conversation date
  // is at or older than 1972 or else we do not convert, the date is
  // pass to us in this format "1972-4-1" YYYY-MM-DD
  userDarray := strings.split(userDate, "-")
    if int(userDarray[0]) < 1972:
      return seconds

  return seconds
}


func main() {
  now := time.Now().UTC()
  fmt.Println(now)

  lsTableRaw := FetchLeapSeconds()
  lsTable := ParseLeapSeconds(lsTableRaw)
  fmt.Println(lsTable)

  // Signature: Year, Month, Day, Hour, Min, Sec, Nsec, Location
  userDate00 :=  time.Date(1972, time.Jan, 6, 15, 30, 0, 0, time.UTC)
  userDate01 := time.Date(2015, time.March, 31, 0, 0, 0, 0, time.UTC)

}
