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
        stage('Install dependencies') {
            steps {
              dir('${env.WORKSPACE}/backend/src') {
                sh 'go get -d -v .'
                sh 'go install -v .'
              }
              dir('${env.WORKSPACE}/frontend/src') {
                sh 'npm ci'
              }
            }
        }
        stage('Quality checks') {
            environment { 
                CI = 'true'
            }
            steps {
              dir('${env.WORKSPACE}/backend/src') {
                sh 'go test ./... -count=1 -cover'
              }
              dir('${env.WORKSPACE}/frontend/src') {
                sh 'npm run test'
              }
            }
        }
        stage('Build') {
            steps {
              dir('${env.WORKSPACE}/backend/src') {
                sh 'go build'
              }
              dir('${env.WORKSPACE}/frontend/src') {
                sh 'npm run build'
              }
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