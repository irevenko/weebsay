package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var NoColor = &cobra.Command{
	Use:   "nocolor",
	Short: "Print witout colors",
	Long:  `weebsay nocolor`,
	Run: func(cmd *cobra.Command, args []string) {
		printNoColorQuote()
	},
}

func init() {
	RootCmd.AddCommand(NoColor)
}

func printNoColorQuote() {
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
	quote = append(quote, []string{"  © " + quoteResp.Char + " from " + quoteResp.Anime})

	if !strings.HasPrefix(quote[0][0], `"`) {
		quote[0][0] = `"` + quote[0][0]
		quote[len(quote)-3][len(quote[len(quote)-3])-1] = quote[len(quote)-3][len(quote[len(quote)-3])-1] + `"`
	}

	quoteCounter := 0
	for i, v := range imgSlice {
		if i > 0 && i < len(imgSlice)-1 {
			if i > 3 && quoteCounter < len(quote) {
				fmt.Println(v + strings.Join(quote[quoteCounter], " "))
				quoteCounter++
			} else {
				fmt.Println(v)
			}
		}
	}
}
