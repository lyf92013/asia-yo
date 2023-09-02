package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
)

type Data struct {
	Currencies map[string]map[string]float64
}

func ConverseCurrency(sourceCurrency, targetCurrency string, amount float64) string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	data, err := os.ReadFile("./data/data.json")
	if err != nil {
		panic(err)
	}

	var d Data
	if err := json.Unmarshal(data, &d); err != nil {
		panic(err)
	}

	source, ok := d.Currencies[sourceCurrency]
	if !ok {
		panic(fmt.Sprintf("Source currency %s not support", sourceCurrency))
	}

	rate, ok := source[targetCurrency]
	if !ok {
		panic(fmt.Sprintf("Target currency %s not support", targetCurrency))
	}

	return toDollar(amount * rate)
}

func toDollar(n float64) string {
	t := strings.Split(fmt.Sprintf("%.2f", n), ".")

	i, err := strconv.ParseInt(t[0], 10, 64)
	if err != nil {
		panic(err)
	}

	f := humanize.Comma(i)
	if len(t) == 2 {
		f = f + "." + t[1]
	}

	return "$" + f
}
