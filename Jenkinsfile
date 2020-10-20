void setBuildStatus(String message, String state) {
    echo "Set status on GitHub to: " + state + " with message: " + message
    step([
        $class: "GitHubCommitStatusSetter",
        reposSource: [$class: "ManuallyEnteredRepositorySource", url: "git@github.com:wh1tew0lf/musicreviewtool.git"],
        contextSource: [$class: "ManuallyEnteredCommitContextSource", context: "ci/jenkins/build-status"],
        errorHandlers: [[$class: "ChangingBuildStatusErrorHandler", result: "UNSTABLE"]],
        statusResultSource: [ $class: "ConditionalStatusResultSource", results: [[$class: "AnyBuildResult", message: message, state: state]] ]
    ]);
}

pipeline {
    agent any
    stages {
        stage ('Set GitHub Status') {
            steps {
                step([$class: 'GitHubSetCommitStatusBuilder', contextSource: [$class: 'ManuallyEnteredCommitContextSource']])
            }
        }
        stage('Build images') {
            steps {
              dir('backend') {
                sh 'docker build . -t "mrt-backend" -f backend.docker'
              }
              dir('frontend') {
                sh 'docker build . -t "mrt-frontend" -f frontend-test.docker'
              }
            }
        }
        stage('Quality checks') {            
            steps {
              dir('backend') {
                sh 'docker run --rm --name mrt-backend-unit-tests mrt-backend:latest go test ./... -count=1 -cover'
              }
              dir('frontend') {
                sh 'docker run --rm --env CI=true --name mrt-frontend-unit-tests mrt-frontend:latest npm run test'
              }
            }
        }
        stage('Tear down') {
          steps {
            sh 'docker rmi mrt-backend:latest'
            sh 'docker rmi mrt-frontend:latest'
          }
        }
    }
    post {
        success {
            setBuildStatus("Build success", "SUCCESS")
        }
        failure {
            setBuildStatus("Build failed", "FAILURE")
        }
    }
}