/*
Use UTC as your input.
Include a function to calculate TT from UTC. You'll need a table or function that knows the leap second count for a given date.
Perform all astronomical calculations (Julian Day, lunar/solar longitudes) using TT.
Convert the final result back to UTC for display.

Converting to TT = UTC + Delta T (LeapSeconds) + 32.184 seconds
Where Delta T (LeapSeconds) is the total number of leap seconds inserted between 1972 and the current date

*/






package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
    "strings"
    "unicode"
)

func parseLeapSeconds(leapSecDoc string)([][]string){

  //separate table from data
  index := strings.Index(leapSecDoc, "#NTP")

  //fmt.Println(index)
  leapSecTable := leapSecDoc[index:]

  //cut out everything but the table
  index = strings.Index(leapSecTable, "#")
  leapSecTable = leapSecTable[index+1:]
  index = strings.Index(leapSecTable, "#")
  leapSecTable = leapSecTable[index+1:]
  leapSecTable = strings.ReplaceAll(leapSecTable, "#", "")
  index = strings.Index(leapSecTable, "\n")
  leapSecTable = leapSecTable[index+1:]

  //splitting by new line character remove the rest of what we have
  separator := "\n"
  cleanTable := strings.Split(leapSecTable, separator)

  //find the lone white space
  for index, element := range cleanTable {
    if element == ""{
      cleanTable = cleanTable[:index]
      break
    }
  }

  //remove leading whitespace from each element
  for element := range cleanTable{
    cleanTable[element] = strings.TrimLeftFunc(cleanTable[element], unicode.IsSpace)
  }

 //create slice for each row, splitting by white space
  result := make([][]string, len(cleanTable))
  for i, element := range cleanTable{
    result[i] = strings.Fields(element)
  }
  //table if organized thus:
  //NTP Time      DTAI    Day Month Year
  //2287785600     11       1 Jul 1972
  //fmt.Println(result)


  // change jan and jul to integers
  for i := range result {
    for j := range result[i]{
      switch strings.ToLower(result[i][j]){
        case "jan":
         result[i][j] = "1"

        case "jul":
          result[i][j] = "6"
      }
    }
  }
  return result
}

// fetch official leapsecond data
func fetchLeapSeconds()(string, error){
    // The URL for the leap second list
    url := "https://data.iana.org/time-zones/data/leap-seconds.list"

    // Make the HTTP Get request
    resp, err := http.Get(url)
    if err != nil{
      panic(err)
    }
    defer resp.Body.Close()

    //Check the response status
    if resp.StatusCode != http.StatusOK{
      panic(fmt.Sprintf("HTTP %d: %s", resp.StatusCode, resp.Status))
    }

    // Read and Print the file content
    body, err := io.ReadAll(resp.Body)
    if err != nil{
      panic(err)
    }

    textBody := string(body)
    return textBody, nil
  }


func terriestrialTime(lst [][]string, y int, m int, d int)(string){
  /* To get terriestrial time we take UTC + leap Seconds
  + 32.184
  */

  for i := range lst{
    for j := range lst[i]{
      if y <= int(lst[i][4]):

    }
  }


//What should the algorithm do? If y is smaller than or equal to? What is the best
// way to accomplish what we want?

}


func main(){
  fmt.Println("current time:", time.Now())
  leapSecondsDoc, err := fetchLeapSeconds()

  if err != nil{
    panic(err)
  }

  table := parseLeapSeconds(leapSecondsDoc)

  fmt.Println(table)
  //year := 2015
  //day := 31
  //month := 3

  //ttTime := terriestrialTime(table, year, month, day)

}
