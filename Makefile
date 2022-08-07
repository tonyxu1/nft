progres:
	docker run --name postgres -d -p 5432:5432 \
	 -e POSTGRES_USER=nft \
	 -e POSTGRES_PASSWORD=nft \
	 -e POSTGRES_DB=nft \
	 library/postgres
