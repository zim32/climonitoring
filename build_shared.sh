#!/bin/bash

# EXECUTE AS ROOT

go install -buildmode=shared -linkshared std

# metrics
go build -linkshared -o ./bin/cm_m_filestat ./metrics/filestat
go build -linkshared -o ./bin/cm_m_cpu ./metrics/cpu
go build -linkshared -o ./bin/cm_m_ram ./metrics/ram
go build -linkshared -o ./bin/cm_m_tcp ./metrics/tcp

# filters
go build -linkshared -o ./bin/cm_f_false ./filters/false
go build -linkshared -o ./bin/cm_f_true ./filters/true
go build -linkshared -o ./bin/cm_f_regex ./filters/regex

# processors
go build -linkshared -o ./bin/cm_p_bulk ./processors/bulk
go build -linkshared -o ./bin/cm_p_message ./processors/message
go build -linkshared -o ./bin/cm_p_eot2nl ./processors/eot2nl
go build -linkshared -o ./bin/cm_p_nl2eot ./processors/nl2eot
go build -linkshared -o ./bin/cm_p_debounce ./processors/debounce
go build -linkshared -o ./bin/cm_p_truncate ./processors/truncate
go build -linkshared -o ./bin/cm_p_watchdog ./processors/watchdog
go build -linkshared -o ./bin/cm_p_multiline ./processors/multiline

# outputs
go build -linkshared -o ./bin/cm_o_telegram ./outputs/telegram
go build -linkshared -o ./bin/cm_o_opsgenie ./outputs/opsgenie
go build -linkshared -o ./bin/cm_o_smtp ./outputs/smtp