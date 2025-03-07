#!/bin/bash

# Jenkins container setup script

# Pull the latest Jenkins image
docker pull jenkins/jenkins:lts

# Create a Docker network for Jenkins
docker network create jenkins

# Run a Jenkins container
docker run --name jenkins -d \
  -p 8080:8080 -p 50000:50000 \
  -e JENKINS_OPTS="--accessLoggerClassName=winstone.accesslog.SimpleAccessLogger --simpleAccessLogger.format=combined --simpleAccessLogger.file=/dev/stdout" \
  --network jenkins \
  jenkins/jenkins:lts

echo "Jenkins is up and running on http://localhost:8080"
