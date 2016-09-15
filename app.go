package main

import (
	"fmt"
	"os"
    "flag"
    "strings"
    "unicode/utf8"
	"github.com/caa-dev-apps/cefscan_v2/cef"
)

func mooi_log(a ...interface{}) (n int, err error) {
	return
}

func main() {

	l_root := flag.String("r", "", "Root folder")
	l_key := flag.String("k", "", "Search Key")
	l_info := flag.String("i", "", "info")
	l_limit := flag.Int("l", 1000, "Limit count")
    
	flag.Parse()
	if flag.NFlag() < 2 {
		flag.PrintDefaults()
        os.Exit(-1)
	}

    if l_info != nil {
        s1 := *l_key + `: ` + *l_info
        fmt.Println(s1)
        s2 := strings.Repeat("-", utf8.RuneCountInString(s1))
        fmt.Println(s2)
    } 
    
    ix := 0
    for l_args := range getScanArgs(l_root) {
        fmt.Println(l_args.CefFilename())

        l_in_meta := false
        
        for kv := range cef.ReadCefHeader_v2(&l_args) {
    
            if kv.Key() == *l_key {
                fmt.Println(kv.KVS(true))
                ix++
            } else if kv.Key() == "START_META" && kv.Val()[0] == *l_key {
                l_in_meta = true
                fmt.Println(kv.KVS(true))
            } else if l_in_meta == true {
                fmt.Println(kv.KVS(false))
                if kv.Key() == "END_META" {
                    fmt.Println(" ")
                    l_in_meta = false
                    ix++
                }
            }
            
            if ix >= *l_limit {
                fmt.Println("\n\n")
                return
            }
        }
    }
    
    if ix == 0 {
        fmt.Println("X")
    } 
    fmt.Println("\n\n")
    
}
