# a cli for the vcloud api

## api spec

https://pubs.vmware.com/vcd-55/topic/com.vmware.ICbase/PDF/vcd_55_api_guide.pdf  
https://pubs.vmware.com/vca/index.jsp#com.vmware.vcloud.api.doc_56/GUID-F4BF9D5D-EF66-4D36-A6EB-2086703F6E37.html

## usage

set the following env vars
* VCD_URL
* VCD_USER
* VCD_PASSWORD
* VCD_ORG

if you are behind a proxy set also
* HTTPS_PROXY

explore the possiblities of the cli by using the help.  

the command structure of the vcloud-cli:  
`vcloud-cli --network query allocatedips`

`query` --> root command  
`allocatedips` --> sub command  
`--network` --> argument for the last command

at every level you can use the help:    
* `vcloud-cli query --help`
* `vcloud-cli query allocatedips --help`

### examples

get all vapps
`vcloud-cli query vapp`

get all vapps and the xml
`vcloud-cli query vapp --verbose true`
