# goUsingAPI

kayıt -> curl --location --request POST 'localhost:8080/kayit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "İbrahim",
    "surname": "Budak"
}'

listelemek -> curl --location --request GET 'localhost:8080/list'
