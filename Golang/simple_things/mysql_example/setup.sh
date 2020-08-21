if [[ $# -eq 0 ]]; then
  echo "You must pass at least one argument"
  exit 1
fi

if [[ $1 = "import" ]]; then
  cat db_dump.sql | mysql -h127.0.0.1 -uuser -ppassword mydb
  echo "import the db pls"
fi

if [[ $1 = "prune" ]]; then
  docker container prune --force
  docker image prune -a --force
  docker volume prune --force
fi