#!/bin/bash
set -xe

mysqladmin -uroot --password="dondokodon" create bookmark_tarou
mysqladmin -uroot --password="dondokodon" create bookmark_tarou_test

mysql -uroot --password="dondokodon" bookmark_tarou < /app/db/mysql/schema.sql
mysql -uroot --password="dondokodon" bookmark_tarou_test < /app/db/mysql/schema.sql
