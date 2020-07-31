.PHONY: deploy

build:
	go build server.go

build-linux:
	GOOS=linux go build server.go

clean:
	rm server

deploy-setup:
	cd deploy; pipenv install

deploy:
	cd deploy; pipenv run ansible-playbook deploy.yml

stop:
	cd deploy; pipenv run ansible-playbook stop.yml
