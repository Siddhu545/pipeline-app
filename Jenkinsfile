pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                echo 'Getting the code ...'
                checkout scm
            }
        }
        stage('SonarQube Scan') {
            steps {
                echo 'Scanning code with SonarQube...'
                withCredentials([string(credentialsId: 'sonar-token', variable: 'SONAR_TOKEN')]) {
                    sh '''
                        docker run --rm \
                          --network pipeline-net \
                          -v "$(pwd):/usr/src" \
                          sonarsource/sonar-scanner-cli \
                          -Dsonar.host.url=http://sonarqube:9000 \
                          -Dsonar.token=$SONAR_TOKEN
                    '''
                }
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