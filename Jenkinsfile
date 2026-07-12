pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                echo 'Getting the code ...'
                checkout scm
            }
        }
        stage('Build Docker Image') {
            steps {
                echo 'Building docker image...'
                sh 'docker build -t pipeline-app:${BUILD_NUMBER} . '
            }
        }
    }
}