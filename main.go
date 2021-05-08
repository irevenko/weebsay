package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

const cherryBlossom = `
               ,╥╦╦  ╓╦╥,                 
             ╓▓▒▒▒▒▒▒▒▒▒▒▓╗               
             ▒▒▒▒▒░░░░▒▒▒▒▒N              
            j▒▒▒▒░░░░░░░▒▒▒[              
   ,╦╣╣╢╣╗╖  ╣▒▒▒░░░░░░░▒▒▒  ╓╗╣╢╢╬╦,     
 ╒▓▒▒▒▒▒▒▒▒▒╢╖▒▒▒▒▒▒▒▒▓█▒▒╗╣▒▒▒▒▒▒▒▒▒▓╕   
  ╣▒▒▒░░░░▒▒▒▒▒▒╢╢█▀╣╣ ╣▒▒▒▒▒▒░░░░░▒▒▒    
g╣▒▒░░░░░░░▒▀▀╙╨▓╢╢╢▒▒╫▓╝╜▀▀▒░░░░░░░▒▒╢@  
▓▒▒▒▒░░░░░░▒▒▒▓@╖▒▒▒▒▒▒╫@╣╢▒▒▄░░░░░▒▒▒▒╫  
 ▓▒▒▒▒▒▒▒▒▒▒▒╢▀▓▓▒▒▒▒▒▒▓φ@@@╫▀░▒▒▒▒▒▒▒╢   
  "╩╣▒▒▒▒▒▒▒╣╢╢╜░╠▓▓▓▓▓▒▓▓▒▒▒▒▒▒▒▒▒╣╩"    
        ╓╣▒▒▓▄╓╣▓▄╫╫╢░╟╢▒▀▒▒▒▒╢╦          
     ,@▒▒▒▒▒░░▒▒▒▒╢╫╣╣▄▓▒▒▒▒▒▒▒▒▒╢,       
     ▒▒▒▒▒░░░░░▒▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒L      
    ]╣▒▒▒░░░░░░░▒▒▒╫║▒▒▒░░░░░░░▒▒▒╢╛      
     ▓▒▒▒▒▒░░░░▒▒▒▒  ╣▒▒▒▒░░░░▒▒▒▒▓       
      "╙ ║▒▒▒▒▒▒▒╫    ╣▒▒▒▒▒▒▒▒           
          ╩╩╬╣Ñ╩        ╨╩╬╬▓╩            
`

const apiLink = "https://animechan.vercel.app/api/random"

func main() {
	resp, err := http.Get(apiLink)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var quoteResp AnimeChanResponse
	json.Unmarshal(body, &quoteResp)

	imgSlice := strings.Split(cherryBlossom, "\n")
	quoteSlice := strings.Split(quoteResp.Quote, " ")
	quote := chunkSlice(quoteSlice, 8)
	quote = append(quote, []string{""})
	quote = append(quote, []string{color.HiRedString("  © ") + quoteResp.Char + " from " + quoteResp.Anime})

	if !strings.HasPrefix(quote[0][0], `"`) {
		quote[0][0] = color.HiRedString(`"`) + quote[0][0]
		quote[len(quote)-3][len(quote[len(quote)-3])-1] = quote[len(quote)-3][len(quote[len(quote)-3])-1] + color.HiRedString(`"`)
	}

	quoteCounter := 0
	for i, v := range imgSlice {
		if i > 0 && i < len(imgSlice)-1 {
			if i > 3 && quoteCounter < len(quote) {
				fmt.Println(color.HiRedString(v) + strings.Join(quote[quoteCounter], " "))
				quoteCounter++
			} else {
				fmt.Println(color.HiRedString(v))
			}
		}
	}
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string

	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

type AnimeChanResponse struct {
	Anime string `json:"anime"`
	Char  string `json:"character"`
	Quote string `json:"quote"`
}
