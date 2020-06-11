RESULT=`mysqlshow --user=root --password="" testdb| grep -v Wildcard | grep -o testdb`
if [ "$RESULT" != "testdb" ]; then
    echo "Create database"
    mysql -e "CREATE DATABASE testdb;"
    #mysql -e "GRANT ALL PRIVILEGES ON *.* TO 'dbuser'@'localhost' IDENTIFIED BY 'secret';"

    echo "Run database migration"
    bin/portfolio-server migrate -H localhost -d testdb --db_password="" -u root -m infrastructure/mysql/migrations
fi
