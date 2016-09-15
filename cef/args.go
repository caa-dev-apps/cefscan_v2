package cef

import (
	"errors"
	"flag"
//x 	"fmt"
)

type strslice []string

func (ss *strslice) String() string {
//x 	return fmt.Sprintf("%s", *ss)
	return ""
}

func (ss *strslice) Set(s string) error {
	*ss = append(*ss, s)
	return nil
}

type CefArgs struct {
	m_includes strslice
	m_cefpath  *string
}

func (a1s *CefArgs) init() error {
	err := error(nil)

	flag.Var(&a1s.m_includes, "i", "Include Folders")
	a1s.m_cefpath = flag.String("f", "", "Cef file path (.cef/.cef.gz)")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		err = errors.New("Error: Invalid number of cmdline args")
	}

	return err
}

func (a1s *CefArgs) Dump() {
//x 	println("cef path:  ", *a1s.m_cefpath)
//x 	println("include folders ", a1s.m_includes)
}

func (a1s *CefArgs) CefFilename() string {
    return *a1s.m_cefpath
}


func NewCefArgs() (args CefArgs, err error) {
	args = CefArgs{}
	err = args.init()

	return args, err
}


func NewScanArgs(i_includes []string, i_cefpath *string) (args CefArgs) {
	return CefArgs{i_includes, i_cefpath}
}

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        