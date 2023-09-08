package main

import "fmt"

func LiveIn(name string, m *map[string]string) {
	addr := *m
	fmt.Printf("%s lives in %s", name, addr[name])
	fmt.Println()
}

func main() {
	m := make(map[string]string)
	m["Annie"] = "Osaka"
	m["Bale"] = "Fukuoka"
	m["Cindy"] = "Nagoya"
	m["Dean"] = "Tokyo"

	LiveIn("Annie", &m)
	LiveIn("Dean", &m)
}
