#!/bin/bash

# Define the table name and index name
TABLE_NAME="MyDataModel"
INDEX_NAME="DataTypeDataValueIndex"

# Check if the table exists, and delete it if it does
if aws2 dynamodb describe-table --table-name $TABLE_NAME --endpoint-url http://localhost:8000 2>&1 | grep -q "Table"; then
    echo "Deleting existing table..."
    aws2 dynamodb delete-table --table-name $TABLE_NAME --endpoint-url http://localhost:8000 > /dev/null 2>&1
    echo "Waiting for table to be deleted..."
    aws2 dynamodb wait table-not-exists --table-name $TABLE_NAME --endpoint-url http://localhost:8000 > /dev/null 2>&1
fi

# Create a new table with a GSI
echo "Creating new table with GSI..."
aws2 dynamodb create-table \
    --table-name $TABLE_NAME \
    --attribute-definitions \
        AttributeName=Id,AttributeType=S \
        AttributeName=DataType,AttributeType=S \
        AttributeName=DataValue,AttributeType=S \
    --key-schema \
        AttributeName=Id,KeyType=HASH \
        AttributeName=DataType,KeyType=RANGE \
    --billing-mode PAY_PER_REQUEST \
    --global-secondary-indexes \
        "IndexName=$INDEX_NAME,KeySchema=[{AttributeName=DataType,KeyType=HASH},{AttributeName=DataValue,KeyType=RANGE}],Projection={ProjectionType=ALL}" \
    --endpoint-url http://localhost:8000 > /dev/null 2>&1

echo "Waiting for table to be active..."
aws2 dynamodb wait table-exists --table-name $TABLE_NAME --endpoint-url http://localhost:8000 > /dev/null 2>&1

echo "Table and index creation complete."
