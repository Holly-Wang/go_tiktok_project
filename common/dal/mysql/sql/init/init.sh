# init database for docker-compose test
set -e
echo '1. 导入数据...'

mysql -u root < /sql/database.sql

mysql -u root -D tiktok < /sql/schema.sql

mysql -u root -D tiktok < /sql/data.sql

echo '2. 导入完成...'
