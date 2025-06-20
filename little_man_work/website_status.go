package main

import (
	"fmt"
    "sync"
    "time"
    "net/http"
	"os"

)
func analyse_status_websitev(url string, ch chan string, wg *sync.WaitGroup){

	defer wg.Done()

	var check string
	client := http.Client{
		Timeout: 3 *time.Second,
		} 
		zip := "https://"+url
		respond, err := client.Get(zip)
		if err != nil || respond.StatusCode == 404 {
			check = "retry"
		} else {
			check = "up"
		}

		// second try if the err != nil or the status code is not found (perhaps i should add a 500 code status right ?)
		// } else the check becomes up and the channels greps the the final thing
		if check == "retry" {
			fmt.Println("There seems to be in issue.")
			fmt.Printf("We will retry checking %s for a second time:\n", url)
			respond, err = client.Get(zip)	
			if err != nil || respond.StatusCode == 404 {
				check = "down"
			}		
		} 
		
		ch <-fmt.Sprintf("%s is %s", url,check)
	}
	func main() {
		start := time.Now()
		
		var wg  sync.WaitGroup 
		sites := []string{"google.com", "zip.com", "doesnotexist.com",}
		status := make(chan string, len(sites))
		file , err := os.Create("status.txt")
		defer file.Close()
		if err != nil {
			fmt.Println("the file had some management issues.\n")
			os.Exit(0)
		}

	for i:= 0 ; i<3; i++{
		wg.Add(1)
		go analyse_status_websitev(sites[i],status, &wg )
		
	}
	go func() {
		wg.Wait()
		close(status)
	}()	
	for msg := range status{
		fmt.Println(msg)
		message_line := msg+"\n"
		file.WriteString(message_line)

	}
	fmt.Printf("the time took for this programme to finish is %v seconds\n", start.Second())
}