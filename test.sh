echo -e "\nTests Using curl \n\n"
echo -e "\n Hitting /accounts/keys\n"

curl -n -X POST http://localhost:8000/accounts/keys \
-H "Content-Type:application/json" \
-d "{\"public_key\":\"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDSJZZmFq+OThseFy+h9xv5baXtLfzC/LnWOTjtAzgisvT4yk/vcMV1QxsUkZ+xpriJGqDqhBS/vpOEIStwd3oPpp24UwF+kzgE2f+CFTxJvVFTs5G2gnxfAjtvyykHU0yO2vpHIqHfVQMGaXGNPyEtZ2+DnkuvCZDNyOBda8lbFPAQdiS6AHq9EOfZ/Zq9Ia3jPXZc0IwCSOL0D90jJsRNpFjCtQEBWoPyTAX06YEi3JFNDjK7OKkwBVqr63rfBx8ggz0cAQIQaicOqgmRJRM8exnpKCt2Lf5Brwf+Jn8o1WN2mm2j5R+z8ofAdsALx01alUXJfbD7111uFIzohYuR user@server\"}"

echo -e "\n\n Hitting GET /account/key/{id} \n"
curl http://localhost:8000/account/keys/1

echo -e "\n\n Hitting DELETE /accounts/keys/{id}\n"
curl -n -X DELETE http://localhost:8000/account/keys/tobedeleted

echo -e "\n\n Hitting GET /account/keys\n"
curl -n -X GET http://localhost:8000/account/keys
