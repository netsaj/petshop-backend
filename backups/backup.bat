@echo off
set TIMESTAMP=%DATE:~10,4%%DATE:~4,2%%DATE:~7,2%%TIME:~0,2%_%TIME:~3,2%_%TIME:~6,2%
SET PGPATH="pg_dump.exe"
SET PGPASSWORD=linux
%PGPATH%  --file "backups\%TIMESTAMP%-petshop.backup"  --host "localhost" --port "5432" --username "postgres" --no-password --verbose --format=c --blobs --section=pre-data --section=data --section=post-data "petshop"

