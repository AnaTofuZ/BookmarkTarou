#!/bin/bash
set -xe

mysqladmin -uroot create bookmark_tarou
mysqladmin -uroot create bookmark_tarou_test

mysql -uroot bookmark_tarou < /app/db/schema.sql
mysql -uroot bookmark_tarou_test < /app/db/schema.sql
