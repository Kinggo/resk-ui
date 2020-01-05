package main

import (
	"encoding/json"
	"fmt"
	"github.com/Kinggo/resk/services"
)

func main() {

	d, e := json.Marshal(&acservices.AccountTransferDTO{})
	fmt.Println(e)
	fmt.Println(string(d))
}
