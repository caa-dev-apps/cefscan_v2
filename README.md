# cefscan_v2

A CEF scanner for the CAA team
    A utility for quickly obtaining a sample set of cef attributes and meta-data values exactly matching a key passed with the -k argument.
    The search routine scans all the cef files contained under the root folder (-r argument) and its sub folders including any ceh files 
    referenced. The search continues until either the max number of results is found (-l argument) or all cef files have been searched.

    
# Sample command line
    
    cefscan_v2 -r /root-cef-ceh_samples -k MISSION_TIME_SPAN -l 10


# go - download + install

    go get github.com/caa-dev-apps/cefscan_v2   (requires go + git)
        
        
# Local builds

    set GOOS=linux&&set GOARCH=amd64&& go build -v .
    set GOOS=linux&&set GOARCH=386&& go build -v .

    