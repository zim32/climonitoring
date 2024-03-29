### Index

- Metrics
    - [cm_m_cpu](#cm_m_cpu)
    - [cm_m_filestat](#cm_m_filestat)
    - [cm_m_ram](#cm_m_ram)
    - [cm_m_tcp](#cm_m_tcp)
    - [cm_m_procinfo](#cm_m_procinfo)
    - [cm_m_netstat](#cm_m_netstat)
- Filters
    - [cm_f_false](#cm_f_false)
    - [cm_f_true](#cm_f_true)
    - [cm_f_regex](#cm_f_regex)
    - [cm_f_enable](#cm_f_enable)
- Processors
    - [cm_p_bulk](#cm_p_bulk)
    - [cm_p_debounce](#cm_p_debounce)
    - [cm_p_eot2nl](#cm_p_eot2nl)
    - [cm_p_nl2eot](#cm_p_nl2eot)
    - [cm_p_message](#cm_p_message)
    - [cm_p_truncate](#cm_p_truncate)
    - [cm_p_watchdog](#cm_p_watchdog)
    - [cm_p_bandwidth](#cm_p_bandwidth)
    - [cm_p_average](#cm_p_average)
    - [cm_p_multiline](#cm_p_multiline)
- Outputs
    - [cm_o_telegram](#cm_o_telegram)
    - [cm_o_opsgenie](#cm_o_opsgenie)
    - [cm_o_smtp](#cm_o_smtp)
- [Examples](#examples)

### Overview

Docs in progress. If you know Russian, you can read [this](https://habr.com/ru/post/454674/) post.

### Reference

#### Metrics


##### cm_m_cpu  
Outputs cpu info

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Interval, seconds | N | 1

Example:
````
{"LoadAvg1":2.4023438,"LoadAvg5":2.2089844,"LoadAvg15":2.0039062}
````

##### cm_m_filestat  
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

##### cm_m_ram  
Outputs system memory info

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Interval, seconds | N | 1

Example:
````
{"TotalRam":16707883008,"FreeRam":2437763072,"AvailRam":8628068352,"BuffersRam":1228099584,"CacheRam":5788626944,"UsedRam":7253393408}
````

##### cm_m_tcp  
Check that remote address and port are open. Outputs "true" if open, "false" otherwise

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Interval, seconds | N | 1
-a | Address (host:port) | Y | ""
-n | Network to use (tcp, udp) | N | "tcp"


##### cm_m_procinfo  
This metric collects information about process(es), specified by -pid parameter or -name parameter. 
If pid is given, than only the process with this PID is monitored. If name given, than cm_m_procinfo will monitor all processes whose procname mathces given pattern (-name is regex patter), and all values will be sum of values of individual processes. 

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Interval, seconds | N | 1
-pid | Process PID | N |
-name | Regex pattern to match procname against | N |

Example:
````
cm_m_procinfo -name '^chrome'
{
    "Pid":"10659|11680|11716|12318|12820|12850|15926|5912|5923|5927|5958|6013|7372",
    "ProcessName":"",
    "CommandLine":"",
    "ProcessState":"",
    "NumberOfThreads":173,
    "MemRss":1709178880,
    "MemRssAnon":890515456,
    "MemRssFile":805199872,
    "MemRssShared":13463552,
    "MemRssOwn":1695715328,
    "MemVirtual":11856023552,
    "CoreDumping":0,
    "NetInBytes":96580979009,
    "NetOutBytes":96580979009
    }
````


##### cm_m_netstat
Outputs network statistics

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Interval, seconds | N | 1

Example output:
````
{"NetInBytes":4581044682,"NetOutBytes":4581044682}
````

#### Filters


##### cm_f_false  
Check that input is "false" or "0". Drop line otherwise

##### cm_f_true  
Check that input is "true" or "1". Drop line otherwise

##### cm_f_regex  
Check that input matches given regular expression. Drop line otherwise. Output depends on parameters

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-e | Regular expression | N | ".*"
-o | Output template. If given, cm_f_regex will output this string. {0} will hold whole match, {N}, where N>0 will hold N-th regex group match | N | ""
--invert | Invert match behaviour | N | false

Example:
````
echo "123-345" | cm_f_regex -e '(\d+)-(\d+)' -o 'Whole match: {0}. Converted: {1}/{2}' 
Output: Whole match: 123-345. Converted: 123/345
````


##### cm_f_enable  
This filter reads content of file specified in -f parameter, compares it with -s string parameter and ignores input if they don't match.
Useful to enable/disable monitoring at all, by putting some string in file.

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-f | File path to read content from | Y |
-s | String to be interpreted as true value | N | "1"


#### Processors



##### cm_p_bulk  
Buffer incoming data, and output it in bulks one at a time.

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-s | Max lines to buffer | N | 10
-fi | Flush interval. If specified and >0, than cm_p_bulk will flush internal buffer at most every -fi seconds | N | 0



##### cm_p_debounce  
Will output only first line at specified interval. Use it to send alerts only at given interval

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Debounce interval, seconds | N | 10 


##### cm_p_eot2nl  
Converts EOT to LF symbols. Use it to pipe data to other CLI commands


##### cm_p_nl2eot  
Converts LF to EOT symbols. Use it to pipe data from other CLI commands


##### cm_p_message  
Special processor. Use it before first output command. It will create json message in format needed for output commands.
**There is a default limit of 60 messages per hour, to prevent self spamming. You can change it with -l parameter.**

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-m | Message. If specified, msg.Message field will contain this sting. {stdin} in replaced by input data. Otherwise input data is used | N | ""
-s | Severity. One of ("debug", "info", "warn", "alert", "critical"). Actually you can use any string. But some output providers expect severity to be one of those mentioned above| N | "info 
-h | Host name | N | OS hostname 
-l | Max messages per hour. If more than -l messages received during hour, new messages will be dropped. Counter is reset every hour | N | 60 
Example:
````
echo "1111" | cm_p_message -m "Disk usage is: {stdin}" -s "alert"
Output: {"Severity":"alert","Message":"Disk usage is: 1111","Created":"2019-06-04T12:30:22.571812604+03:00","HostName":"ubuntu"}
````


##### cm_p_truncate  
Truncate input data to specified max length

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-l | Max length. If specified and >0, input will be truncated to this length | N | 0


##### cm_p_watchdog  
Outputs "true" if no input is given for more than specified seconds. Useful to act as a watchdog to monitor file changes f.e. or other stuff where bad condition is that something is not happening for some period of time  

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Max no action interval, seconds | N | 10


##### cm_p_bandwidth
Outputs (input - prev_input) / time.  Useful for calculating bandwidth.

F.e:

    cm_m_procinfo -name '^chrome' | cm_p_eot2nl | jq -cM --unbuffered '.NetInBytes' | cm_p_nl2eot | cm_p_bandwidth | cm_p_eot2nl


##### cm_p_average
Outputs average of input values

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-i | Time interval for averaging, seconds | N | 10 


##### cm_p_multiline
Search input using provided regex pattern and insert EOT (End-Of-Transition) character before matched elements.
Use multiline processor when you need to properly handle multiline logs.

TODO: write more about EOT character and new line processing.

F.e. suppose your log file contains this text:

```
[error] This is some multiline
log message.
Here is another line
[warning] This is another message
with text in another line
```

And you want log messages to be splited by [level] block. You can do it like this (simplified ):

```
tail -F test.log | cm_p_multiline -p '\[(error|warning)\]' | cm_p_message | cm_o_telegram
```

Parameters:

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-p | Regex pattern to search | Y | 


#### Outputs

##### cm_o_telegram  
Send input data to telegram channel

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-cid | Channel ID | Y |
-token | Bot Api token | Y |


##### cm_o_opsgenie  
Send input data to opsgenie channel

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-apiToken | Api token | Y |
-apiEndpoint | Api endpoint. Example "https://api.eu.opsgenie.com/v2/alerts" | Y |
-responderType | Responser type. F.e. "team" | Y |
-responderId | Responser id | Y |


##### cm_o_smtp  
Send input data over smtp

Name | Description | Mandatory | Default
 --- | --- | --- | ---
-host | SMTP host | Y |
-port | SMTP port | Y |
-userName | SMTP user name | Y |
-userPass | SMTP user password | Y |
-from | Send from | Y |
-to | Send to | Y |


#### Examples

Monitor CPU load average
````
cm_m_cpu | cm_p_eot2nl | jq -cM --unbuffered 'if .LoadAvg1 > 1 then .LoadAvg1 else false end' | cm_p_nl2eot | cm_f_regex -e '\d+' | cm_p_debounce -i 60 | cm_p_message -m 'Load average is {stdin}' | cm_o_telegram
````

Monitor number of jobs ib RabbitMq queue
````
while true; do rabbitmqctl list_queues -p queue_name | grep -Po --line-buffered '\d+'; sleep 60; done | jq -cM '. > 10000' --unbuffered | cm_p_nl2eot | cm_f_true | cm_p_message -m 'There are more than 10000 tasks in rabbit queue' | cm_o_opsgenie
````
Monitor that nothing was written to file for more than 10 seconds
````
tail -f out.log | cm_p_nl2eot | cm_p_watchdog -i 10 | cm_p_debounce -i 3600 | cm_p_message -m 'No write to out.log for 10 seconds' -s 'alert' | cm_o_telegram
````
Monitor network input traffic
````
cm_m_netstat | cm_p_eot2nl | jq -cM --unbuffered '.NetInBytes' | cm_p_nl2eot | cm_p_bandwidth | cm_p_average | cm_p_eot2nl
````
Monitor docker service is alive
````
while true; do systemctl show docker; sleep 10; done | cm_p_nl2eot | cm_f_regex -e 'ActiveState=(.*)' -o '{1}' | cm_f_regex -e 'active' --invert | cm_p_debounce -i 3600 | cm_p_message -m 'Docker engine is down. State: {stdin}' | cm_o_opsgenie
````