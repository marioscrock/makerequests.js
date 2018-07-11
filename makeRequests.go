//Code snippets from 
// - http://tleyden.github.io/blog/2016/11/21/tuning-the-go-http-client-library-for-load-testing/

package main

import (
  "fmt"
  "io/ioutil"
  "math"
  "math/rand"
  "net/http"
  "strconv"
  "sync"
  "syscall"
  "time"
)

var myClient *http.Client;
var MAX_IDLE_CONNS = 1000; 
var MAX_IDLE_CONNS_PER_HOST = 1000;

func startMakeRequests(wg *sync.WaitGroup, delay float64, url string) {
  
  time.Sleep(time.Second * time.Duration(delay/1000))

  client := myClient

  req, err := http.NewRequest("GET", urlComplete, nil)
  if err != nil {
    panic(fmt.Sprintf("Got error: %v", err))
  }

  //req.Header.Set("key", "value")
  resp, _ := client.Do(req)

  defer resp.Body.Close()

  if resp.StatusCode == http.StatusOK {
    bodyBytes, _ := ioutil.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)
  }

  wg.Done()

}

func makeRequests(times []float64, url string) {

  var tv syscall.Timeval
  var wg sync.WaitGroup

  // Customize the Transport to have larger connection pool
  defaultRoundTripper := http.DefaultTransport
  defaultTransportPointer, ok := defaultRoundTripper.(*http.Transport)
  if !ok {
      panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
  }
  defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
  //Improve available concurrency
  defaultTransport.MaxIdleConns = MAX_IDLE_CONNS
  defaultTransport.MaxIdleConnsPerHost = MAX_IDLE_CONNS_PER_HOST

  myClient = &http.Client{Transport: &defaultTransport}

  for _, v := range times {
    syscall.Gettimeofday(&tv)
    tw := math.Round(v)
    fmt.Println("GET request " + strconv.FormatInt(int64(tv.Sec)*1e3 + int64(tv.Usec)/1e3 + int64(tw), 10))
    wg.Add(1)
    go startMakeRequests(&wg, v, url)
  }

  wg.Wait()
  fmt.Println("Requests completed!")

}