# Create wallet
curl -k -X POST 'https://0.0.0.0:8080/wallet'

# Add new transactions
curl -k -X POST 'https://0.0.0.0:8080/transaction' -d '{
	"wallet_id": "01E23TVYKC49YZD0BZK929KBKZ",
	"date": "2020-04-02T00:00:00.000Z",
	"sum": "1200.02"
}'
curl -k -X POST 'https://0.0.0.0:8080/transaction' -d '{
	"wallet_id": "01E23TVYKC49YZD0BZK929KBKZ",
	"date": "2020-04-02T00:24:00.000Z",
	"sum": "153.02"
}'
curl -k -X POST 'https://0.0.0.0:8080/transaction' -d '{
	"wallet_id": "01E23TVYKC49YZD0BZK929KBKZ",
	"date": "2020-04-02T03:34:00.000Z",
	"sum": "13.067"
}'

# Fetch wallet history
curl -k -X POST 'https://0.0.0.0:8080/wallet-history' -d '{
	"wallet_id": "01E23TVYKC49YZD0BZK929KBKZ",
	"start_date": "2020-04-02T23:00:00.000Z",
	"end_date": "2020-04-03T08:00:00.000Z"
}'

# Add transaction in past
curl -k -X POST 'https://0.0.0.0:8080/transaction' -d '{
	"wallet_id": "01E23TVYKC49YZD0BZK929KBKZ",
	"date": "2020-03-02T03:34:00.000Z",
	"sum": "13.067"
}'
