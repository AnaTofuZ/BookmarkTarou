#!/bin/bash
set -xe

mysqladmin -uroot create bookmark_tarou
mysqladmin -uroot create bookmark_tarou_test

mysql -uroot bookmark_tarou < /app/db/mysql/schema.sql
mysql -uroot bookmark_tarou_test < /app/db/mysql/schema.sql
