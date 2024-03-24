### ローカルマイグレーション
```
aws2 dynamodb create-table --region us-west-2 --endpoint-url http://localhost:8000 --table-name MyDataModel --attribute-definitions AttributeName=Id,AttributeType=S AttributeName=DataType,AttributeType=S AttributeName=DataValue,AttributeType=S --key-schema AttributeName=Id,KeyType=HASH AttributeName=DataType,KeyType=RANGE --global-secondary-indexes 'IndexName=DataValueIndex,KeySchema=[{AttributeName=DataValue,KeyType=HASH},{AttributeName=Id,KeyType=RANGE}],Projection={ProjectionType=ALL}' --billing-mode PAY_PER_REQUEST
```