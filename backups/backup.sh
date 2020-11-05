DATE=`date '+%Y%m%d%H%M%S'`
export PGPASSWORD=linux
pg_dump --file  "./backups/${DATE}-petshop.backup"  --host "localhost" --port "5432" --username "postgres" --no-password \
  --verbose --format=c --blobs --section=pre-data --section=data --section=post-data "petshop"
