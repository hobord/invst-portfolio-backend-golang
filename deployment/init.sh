until nc -z -v -w30 mysql 3306
do
  echo "Waiting for database connection..."
  # wait for 5 seconds before check again
  sleep 5
done

echo "DB init start"
/app/portfolio-server migrate
echo "App start"
/app/portfolio-server serve

