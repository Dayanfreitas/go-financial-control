build-image:
	docker build -t dayanfretias/finance .

run-app:
	docker-compose -f .devops/app.yml up -d

