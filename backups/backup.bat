@echo off
set TIMESTAMP=%DATE:~10,4%%DATE:~4,2%%DATE:~7,2%%TIME:~0,2%%TIME:~3,2%%TIME:~6,2%
SET PGPATH="pg_dump.exe"
SET PGPASSWORD=linux
%PGPATH% --file -h 127.0.0.1 -p 5432  -U postgres --no-password --verbose --format=c --blobs --section=pre-data --section=data --section=post-data "%TIMESTAMP%-petshop.backup"  petshop
