###Index

- Metrics
    - [cm_m_cpu](#cm_m_cpu)
    - [cm_m_filestat](#cm_m_filestat)
    - [cm_m_ram](#cm_m_ram)
    - [cm_m_tcp](#cm_m_tcp)
- Filters
    - [cm_f_false](#cm_f_false)
    - [cm_f_true](#cm_f_true)
    - [cm_f_regex](#cm_f_regex)
- Processors
    - [cm_p_bulk](#cm_p_bulk)
    - [cm_p_debounce](#cm_p_debounce)
    - [cm_p_eot2nl](#cm_p_eot2nl)
    - [cm_p_nl2eot](#cm_p_nl2eot)
    - [cm_p_message](#cm_p_message)
    - [cm_p_truncate](#cm_p_truncate)
    - [cm_p_watchdog](#cm_p_watchdog)
- Outputs
    - [cm_o_telegram](#cm_o_telegram)
    - [cm_o_opsgenie](#cm_o_opsgenie)

###Reference

####Metrics


#####cm_m_cpu  
Outputs cpu info

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Interval, seconds | N | 1

Example:
````
{"LoadAvg1":2.4023438,"LoadAvg2":2.2089844,"LoadAvg3":2.0039062}
````

#####cm_m_filestat  
Outputs file stat info

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Interval, seconds | N | 1
-f | File path | Y | "" 

Example:
````
{"Size":106,"ModTime":1557311714,"HasChanged":false}
````

#####cm_m_ram  
Outputs system memory info

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Interval, seconds | N | 1

Example:
````
{"TotalRam":16707883008,"FreeRam":2437763072,"AvailRam":8628068352,"BuffersRam":1228099584,"CacheRam":5788626944,"UsedRam":7253393408}
````

#####cm_m_tcp  
Check that remote address and port are open. Outputs "true" if open, "false" otherwise

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Interval, seconds | N | 1
-a | Address (host:port) | Y | ""
-n | Network to use (tcp, udp) | N | "tcp"

####Filters


#####cm_f_false  
Check that input is "false" or "0". Drop line otherwise

#####cm_f_true  
Check that input is "true" or "1". Drop line otherwise

#####cm_f_regex  
Check that input matches given regular expression. Drop line otherwise. Output depends on parameters

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-e | Regular expression | N | ".*"
-o | Output template. If given, cm_f_regex will output this string. {0} will hold whole match, {N}, where N>0 will hold N-th regex group match | N | ""

Example:
````
echo "123-345" | cm_f_regex -e '(\d+)-(\d+)' -o 'Whole match: {0}. Converted: {1}/{2}' 
Output: Whole match: 123-345. Converted: 123/345
````

####Processors
#####cm_p_bulk  
Buffer incoming data, and output it in bulks one at a time.

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-s | Max lines to buffer | N | 10
-fi | Flush interval. If specified and >0, than cm_p_bulk will flush internal buffer at most every -fi seconds | N | 0



#####cm_p_debounce  
Will output only first line at specified interval. Use it to send alerts only at given interval

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Debounce interval, seconds | N | 10 


#####cm_p_eot2nl  
Converts EOT to LF symbols. Use it to pipe data to other CLI commands


#####cm_p_nl2eot  
Converts LF to EOT symbols. Use it to pipe data from other CLI commands


#####cm_p_message  
Special processor. Use it before first output command. It will create json message in format needed for output commands

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-m | Message. If specified, msg.Message field will contain this sting. {stdin} in replaced by input data. Otherwise input data is used | N | ""
-s | Severity. One of ("debug", "info", "warn", "alert", "critical"). Actually you can use any string. But some output providers expect severity to be one of those mentioned above| N | "info 
-h | Host name | N | OS hostname 
Example:
````
echo "1111" | cm_p_message -m "Disk usage is: {stdin}" -s "alert"
Output: {"Severity":"alert","Message":"Disk usage is: 1111","Created":"2019-06-04T12:30:22.571812604+03:00","HostName":"ubuntu"}
````


#####cm_p_truncate  
Truncate input data to specified max length

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-l | Max length. If specified and >0, input will be truncated to this length | N | 0


#####cm_p_watchdog  
Outputs "true" if no input is given for more than specified seconds. Useful to act as a watchdog to monitor file changes f.e. or other stuff where bad condition is that something is not happening for some period of time  

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Max no action interval, seconds | N | 10


####Outputs

#####cm_o_telegram  
Send input data to telegram channel

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-cid | Channel ID | Y |
-token | Bot Api token | Y |


#####cm_o_opsgenie  
Send input data to opsgenie channel

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-apiToken | Api token | Y |
-apiEndpoint | Api endpoint. Example "https://api.eu.opsgenie.com/v2/alerts" | Y |
-responderType | Responser type. F.e. "team" | Y |
-responderId | Responser id | Y |