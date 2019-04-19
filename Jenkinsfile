#!/usr/bin/env groovy

// docker tags must be in lower case
env.DOCKER_TAG = BUILD_TAG.toLowerCase()

stage("build") {
  node {
    checkout scm
    sh 'docker build -t $DOCKER_TAG .'
    sh 'docker run --rm $DOCKER_TAG'
  }
}
