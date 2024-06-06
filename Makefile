build:
	goreleaser release --snapshot --rm-dist

release:
	goreleaser release --rm-dist

docker-build:
	docker build --platform=linux/amd64 -t chrispruitt/ssm-session-cli . 

docker-run:
	docker run --platform=linux/amd64 -it -e AWS_SESSION_TOKEN -e AWS_SECRET_ACCESS_KEY -e AWS_ACCESS_KEY_ID -e AWS_DEFAULT_REGION=us-east-1 chrispruitt/ssm-session-cli start

docker-build-slim:
	docker build --platform=linux/amd64 -f slim.Dockerfile -t chrispruitt/ssm-session-cli:slim . 

docker-run-slim:
	docker run --platform=linux/amd64 -it -e AWS_SESSION_TOKEN -e AWS_SECRET_ACCESS_KEY -e AWS_ACCESS_KEY_ID -e AWS_DEFAULT_REGION=us-east-1 chrispruitt/ssm-session-cli:slim start