FROM postgres:14.2

COPY up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]
