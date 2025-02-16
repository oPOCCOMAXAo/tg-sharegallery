build-docker:
	docker build -t poccomaxa/tg-gallery-bot:latest .

upload-docker:
	docker push poccomaxa/tg-gallery-bot:latest

update-prod:
	docker tag poccomaxa/tg-gallery-bot:latest poccomaxa/tg-gallery-bot:prod
	docker push poccomaxa/tg-gallery-bot:prod

prod: build-docker upload-docker update-prod
