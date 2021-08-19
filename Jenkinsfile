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
        string(name: 'TAGS', description: 'Please change scenario tags')
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
                git branch: 'go/big-refactor-july-2021', url: 'https://github.com/mpermperpisang/golang-automation-v1.git'
            }
        }

        stage('Test') {
            steps {
                echo 'Running test'
                sh 'go mod download && cp env.sample .env'
                echo "${params}"
                sh 'godog --tags=@scenarios --format=cucumber > test/report/cucumber_report.json --random'
            }
        }
    }
}
