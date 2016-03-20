PROJECT ?= sphere-state-service
EB_BUCKET ?= ninjablocks-sphere-docker

APP_NAME ?= sphere-state-service
APP_ENV ?= sphere-state-prod

SHA1 := $(shell git rev-parse --short HEAD | tr -d "\n")

DOCKERRUN_FILE := Dockerrun.aws.json
APP_FILE := ${SHA1}.zip

build:
	docker build -t "ninjasphere/${PROJECT}:${SHA1}" .

push:
	docker push "ninjasphere/${PROJECT}:${SHA1}"

services:
	docker run --name ninja-rabbit -p 5672:5672 -p 15672:15672 -d mikaelhg/docker-rabbitmq
	docker run --name ninja-redis -d redis

local:
	docker run -t -i --rm --link ninja-rabbit:rabbitmq --link ninja-redis:redis -e "RABBIT_URL=amqp://guest:guest@rabbitmq:5672" -e "REDIS_URL=redis://redis:6379/" -e "DEBUG=true" -p 6100:6100 -t "ninjasphere/${PROJECT}:${SHA1}"

deploy:
	sed "s/<TAG>/${SHA1}/" < Dockerrun.aws.json.template > ${DOCKERRUN_FILE}
	zip -r ${APP_FILE} ${DOCKERRUN_FILE} .ebextensions

	aws s3 cp ${APP_FILE} s3://${EB_BUCKET}/${APP_ENV}/${APP_FILE}

	aws elasticbeanstalk create-application-version --application-name ${APP_NAME} \
	   --version-label ${SHA1} --source-bundle S3Bucket=${EB_BUCKET},S3Key=${APP_ENV}/${APP_FILE}

	# # Update Elastic Beanstalk environment to new version
	aws elasticbeanstalk update-environment --environment-name ${APP_ENV} \
       --version-label ${SHA1}

clean:
	rm *.zip || true
	rm ${DOCKERRUN_FILE} || true
	rm -rf build || true

.PHONY: all build push local services deploy clean