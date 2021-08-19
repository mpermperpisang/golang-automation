pipeline {
    agent any

    tools {
        go 'v1.17'
    }

    environment {
        GOROOT="/usr/local/go"
        GOPATH="$HOME/project/go"
        PATH="$PATH:$GOROOT/bin:$GOPATH/bin"
    }

    parameters {
        string(name: 'BRANCH', defaultValue: 'master', description: 'Please change repository branch')
        string(name: 'TAGS', defaultValue: '@scenarios', description: 'Please change scenario tags')
    }

    stages {
        stage('Pre Test') {
            steps {
                echo 'Dependencies'
                sh 'go version'
            }
        }

        stage('Checkout') {
            steps {
                echo 'Git Checkout'
                git branch: "${params.BRANCH}", url: 'https://github.com/mpermperpisang/golang-automation-v1.git'
            }
        }

        stage('Test') {
            steps {
                echo 'Running test'
                sh 'go mod download && cp env.sample .env'
                sh "godog --tags=${params.TAGS} --format=cucumber > test/report/cucumber_report.json --random"
            }
        }
    }

    post {
        always {
            cucumber '**/cucumber_report.json'
        }
    }
}
