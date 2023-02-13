.PHONY: run
run: ## запуск серверов
	echo "Starting msa-bank-client-cs server"
	go run .\msa-bank-client-cs\cmd\msa-bank-credit-cs-server\
	echo "Starting msa-bank-product-cs"
	