# cefscan_v2

A CEF scanner for the CAA team
    A utility for quickly obtaining a sample set of cef attributes and meta-data values exactly matching a key passed with the -k argument.
    The search routine scans all the cef files contained under the root folder (-r argument) and its sub folders including any ceh files 
    referenced. The search continues until either the max number of results is found (-l argument) or all cef files have been searched.

    

# go - download + install

    go get github.com/caa-dev-apps/cefscan_v2   (requires go + git)
        
        
# Local builds

    set GOOS=linux&&set GOARCH=amd64&& go build -v .
    set GOOS=linux&&set GOARCH=386&& go build -v .

    
# Sample command line    
    
    > cefscan_v2 -r C:/_CEF_CEH_EXAMPLES_2013_VALIDATOR_ -k MISSION -l 10 

    MISSION: 
    ---------
    C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\CEF\EDI\#00\C3_CP_EDI_EGD__20111009_V01.cef.gz
    START_META                     [MISSION]                                                   line:8          C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\HEADERS/CL_CH_MISSION.ceh
    ENTRY                          ["Cluster"]                                                 line:9           
    END_META                       [MISSION]                                                   line:10          
     
    C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\CEF\EDI\C3_CP_EDI_EGD__20111009_V01.cef.gz
    START_META                     [MISSION]                                                   line:8          C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\HEADERS/CL_CH_MISSION.ceh
    ENTRY                          ["Cluster"]                                                 line:9           
    END_META                       [MISSION]                                                   line:10          
     
    C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\CEF\EDI\C3_CP_EDI_EGD__20111020_V01.cef.gz
    START_META                     [MISSION]                                                   line:8          C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\HEADERS/CL_CH_MISSION.ceh
    ENTRY                          ["Cluster"]                                                 line:9           
    END_META                       [MISSION]                                                   line:10          
     
    C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\CEF\EDI\C3_CP_EDI_EGD__20111021_V01.cef.gz
    START_META                     [MISSION]                                                   line:8          C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\HEADERS/CL_CH_MISSION.ceh
    ENTRY                          ["Cluster"]                                                 line:9           
    END_META                       [MISSION]                                                   line:10          
     
    C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\CEF\EDI\C3_CP_EDI_EGD__20111022_V01.cef.gz
    START_META                     [MISSION]                                                   line:8          C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\HEADERS/CL_CH_MISSION.ceh
    ENTRY                          ["Cluster"]                                                 line:9           
    END_META                       [MISSION]                                                   line:10          
     
    C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\CEF\EDI\C3_CP_EDI_QZC__20111009_V01.cef.gz
    START_META                     [MISSION]                                                   line:8          C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\HEADERS/CL_CH_MISSION.ceh
    ENTRY                          ["Cluster"]                                                 line:9           
    END_META                       [MISSION]                                                   line:10          
     
    C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\CEF\EDI\C3_CP_EDI_QZC__20111020_V01.cef.gz
    START_META                     [MISSION]                                                   line:8          C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\HEADERS/CL_CH_MISSION.ceh
    ENTRY                          ["Cluster"]                                                 line:9           
    END_META                       [MISSION]                                                   line:10          
     
    C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\CEF\EDI\C3_CP_EDI_QZC__20111021_V01.cef
    START_META                     [MISSION]                                                   line:8          C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\HEADERS/CL_CH_MISSION.ceh
    ENTRY                          ["Cluster"]                                                 line:9           
    END_META                       [MISSION]                                                   line:10          
     
    C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\CEF\EDI\C3_CP_EDI_QZC__20111021_V01.cef.gz
    START_META                     [MISSION]                                                   line:8          C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\HEADERS/CL_CH_MISSION.ceh
    ENTRY                          ["Cluster"]                                                 line:9           
    END_META                       [MISSION]                                                   line:10          
     
    C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\CEF\EDI\C3_CP_EDI_QZC__20111022_V01.cef.gz
    START_META                     [MISSION]                                                   line:8          C:\_CEF_CEH_EXAMPLES_2013_VALIDATOR_\HEADERS/CL_CH_MISSION.ceh
    ENTRY                          ["Cluster"]                                                 line:9           
    END_META                       [MISSION]                                                   line:10          
     
        
    