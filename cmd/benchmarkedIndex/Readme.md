# SFTP Scraper for benchmarked Index Values

## Needed environment variables
### SFTP_FILELIST
list of absolute path based on the ftp root directory to be downloaded and separated by ;

no default value set

Example: `/data/DIAData_DIA_Index_History.csv\;/data/DIAData_BTC_Index_History.csv`
### INFLUXURL 
URL to the influx server in the format: `http://influx:8086`
### SFTP_USERNAME
default: `dia`
### SFTP_PASSWORD
string containing the password
### SFTP_URL
default: `ftp.indexengine.com`
### SFTP_PORT
default: `22`
### INDEX_HISTORICAL_DATA
boolean value in string format true or false
This triggers a complete reimport of all data and overwrites the existing timestamps
default: `false`