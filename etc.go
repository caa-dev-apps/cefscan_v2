package main

import (
//x 	"flag"
//x 	"fmt"
	"os"
    "strings"
	"path/filepath"
	"github.com/caa-dev-apps/cefscan_v2/cef"
    
)

func include_folders(i_root string) (fs []string, r_err error) {

	l_map := make(map[string]int)

	visit := func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == `.ceh` {
			l_map[filepath.Dir(path)] = 1
		}
		return nil
	}

	//x err := filepath.Walk(i_root, visit)
	_ = filepath.Walk(i_root, visit)
	//x fmt.Printf("filepath.Walk() returned %v\n", err)

	l_folders := []string{}
	for k := range l_map {
		l_folders = append(l_folders, k)
	}

	return l_folders, r_err
}

func cef_files(i_root string) (chan string) {

	output := make(chan string, 1)    
    ix := 0

	visit := func(path string, f os.FileInfo, err error) error {
        if strings.HasSuffix(path, "cef") || strings.HasSuffix(path, "cef.gz")  {
            output <- path
            ix++
		}
        
		return nil
	}

    go func() {
        defer close(output)
        
        _ = filepath.Walk(i_root, visit)
        //x fmt.Printf("filepath.Walk() returned %v\n", err)    
    }()

    return output
}


func getScanArgs(i_root *string) chan cef.CefArgs {
    output := make(chan cef.CefArgs, 1)

    go func() {
        defer close(output)
        
        l_includes, _ := include_folders(*i_root)
        
        for l_cefpath := range cef_files(*i_root) {
            l_args :=  cef.NewScanArgs(l_includes, &l_cefpath)
            //x fmt.Println(l_args)
            
            output <- l_args
        }
    }()
    
    return output
}

