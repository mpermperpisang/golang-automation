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
        stage('Checkout') {
            steps {
                echo 'Git Checkout'
                git branch: "${params.BRANCH}", url: 'https://github.com/mpermperpisang/golang-automation-v1.git'
            }
        }

        stage('Package') {
            steps {
                echo 'Properties'
                sh 'go mod download'
            }
        }

        stage('CP File') {
            steps {
                echo 'Properties'
                sh 'cp env.sample .env'
                sh 'cp capabilities-web.properties.sample capabilities-web.properties'
                sh 'cp capabilities-android.properties.sample capabilities-android.properties'
                sh 'cp capabilities-ios.properties.sample capabilities-ios.properties'
            }
        }

        stage('Test') {
            steps {
                echo 'Running test'
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
