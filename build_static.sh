#!/bin/bash

# metrics
go build -o ./bin/cm_m_filestat ./metrics/filestat
go build -o ./bin/cm_m_cpu ./metrics/cpu
go build -o ./bin/cm_m_ram ./metrics/ram
go build -o ./bin/cm_m_tcp ./metrics/tcp
go build -o ./bin/cm_m_procinfo ./metrics/procinfo

# filters
go build -o ./bin/cm_f_false ./filters/false
go build -o ./bin/cm_f_true ./filters/true
go build -o ./bin/cm_f_regex ./filters/regex

# processors
go build -o ./bin/cm_p_bulk ./processors/bulk
go build -o ./bin/cm_p_message ./processors/message
go build -o ./bin/cm_p_eot2nl ./processors/eot2nl
go build -o ./bin/cm_p_nl2eot ./processors/nl2eot
go build -o ./bin/cm_p_debounce ./processors/debounce
go build -o ./bin/cm_p_truncate ./processors/truncate
go build -o ./bin/cm_p_watchdog ./processors/watchdog
go build -o ./bin/cm_p_bandwidth ./processors/bandwidth

# outputs
go build -o ./bin/cm_o_telegram ./outputs/telegram
go build -o ./bin/cm_o_opsgenie ./outputs/opsgenie
go build -o ./bin/cm_o_smtp ./outputs/smtp