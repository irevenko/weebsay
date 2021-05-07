package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const importantFace = `
             ╓∩H▒▒▒▒▒▒▒▒▒∩╖.            
         ╓H▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒╖         
       ╖▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒╥.      
     ╖▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒╗     
    ▒▄▄▄▄▄▄▄▄▄▄▒▒▒▒▒▒▒▒▒▒▄▄▄▄▄▄▄▄▄▄▒    
   ▀███▓▓▓▓▓▓▓▓████▓▓████▓▓▓▓▓▓▓▓███▀   
  ╓▒▒██████████████▒▒▀█████████████▌▒╥  
  ║▒▒█████████████▒▒▒▒█████████████▒▒▒  
  ▒▒▒▒███████████▀▒▒▒▒▒███████████▒▒▒▒  
  ╙▒▒▒▒▒▀▀▀██▀▀▒▒▒▒▒▒▒▒▒▒▀▀███▀▀▒▒▒▒▒║  
   ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░▒▒▒▒▒▒▒▒   
    ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▄▄█▀░▒▒▒▒▒▒▒   
     ║▒▒▒▒▒▒▒▒▒█████████▀▀▒▒▒▒▒▒▒▒║    
      ╙║▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒╜      
         ╝▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒╝        
            ╙╝╣▒▒▒▒▒▒▒▒▒▒╣╝╜           
`

const importantFinger = `
                         ╓╗╖╖           
                        ╒▒▒▒▒╣          
                        ╞▒▒▒▒╢          
                        ╟▒▒▒▒╢          
                        ╟▒▒▒▒╢          
                        ╟╣▒▒▒╢          
              ╦╣╣╬╦▓▒▒▒▒▓╣▒▒▒╢          
        ╫▒▒▒▒▓╣▒▒▒╫▌▒▒▒▒╫╣▒▒▒╢          
        ╣▒▒▒▒▓▒▒▒▒╢▓▒▒▒▒╫╣▒▒▒╫          
        ▓╣╣╣▓▓▓▓╣╢╢▓╣╢▒▒╢╢▒▒▒╢═         
        ▓▓▓▓▓▓▓▓╣╣╣╣▒╢╢▒▒▒▒▒▒╢╣         
        ▓╣╣╣╣╢╢╢▓▓▓▓▓▓▓╣╣▒▒▒▒▒▒▓        
        ╟╣▒▒▒▒▒╢╢▓▓▓▓▓▓▓▓▒▒▒▒▒╢▓        
         ╣▒▒▒▒▒▒▒▒╢╢╢╢╣╣╣▒▒▒▒╢╫▓        
         ▐╣▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒╢╢▓▓         
          ▐▓╣▒▒▒▒▒▒▒▒▒▒▒╣╢╢▓▓▓          
            ▓▓▓╣╣▒▒▒▒╢╢╢▓▓▓▓            
               ╙▀▀▓▓▓▓▓▀▀╙ 
`

func main() {
	resp, err := http.Get("https://animechan.vercel.app/api/random")
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	var animeResp Response
	json.Unmarshal(body, &animeResp)

	imgSlice := strings.Split(importantFinger, "\n")
	// quoteSlice := strings.Split(animeResp.Quote, " ")

	for i, v := range imgSlice {
		if i > 0 && i < len(imgSlice)-1 {
			if i == 4 {
				fmt.Println(v + " " + animeResp.Quote)
			} else if i == 6 {
				fmt.Println(v + "                     " + animeResp.Char)
			} else {
				fmt.Println(v)
			}
		}
	}
}

type Response struct {
	Anime string `json:"anime"`
	Char  string `json:"character"`
	Quote string `json:"quote"`
}
