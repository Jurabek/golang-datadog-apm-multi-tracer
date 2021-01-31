aws dynamodb create-table \
    --table-name Items \
    --attribute-definitions \
        AttributeName=Id,AttributeType=S \
        AttributeName=Name,AttributeType=S \
    --key-schema \
        AttributeName=Id,KeyType=HASH \
        AttributeName=Name,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=10,WriteCapacityUnits=5 --endpoint-url http://localhost:8000

aws dynamodb put-item \
    --table-name Items  \
    --item \
        '{"Id": {"S": "b198785c-3e03-4045-be10-fd9feef4ac17"}, "Name": {"S": "User 1"} }' \
    --endpoint-url http://localhost:8000
