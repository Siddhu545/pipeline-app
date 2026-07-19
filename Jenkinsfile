// pipeline {
//     agent any

//     stages {
//         stage('Checkout') {
//             steps {
//                 echo 'Getting the code ...'
//                 checkout scm
//             }
//         }
//         stage('SonarQube Scan') {
//             steps {
//                 echo 'Scanning code with SonarQube...'
//                 withCredentials([string(credentialsId: 'sonar-token', variable: 'SONAR_TOKEN')]) {
//                     sh '''
//                         docker run --rm \
//                           --network pipeline-net \
//                           -v "$(pwd):/usr/src" \
//                           sonarsource/sonar-scanner-cli \
//                           -Dsonar.host.url=http://sonarqube:9000 \
//                           -Dsonar.login=$SONAR_TOKEN \
//                           -Dsonar.projectKey=pipeline-app 
//                     '''
//                 }
//             }
//         }
//         stage('Build Docker Image') {
//             steps {
//                 echo 'Building docker image...'
//                 sh 'docker build -t pipeline-app:${BUILD_NUMBER} . '
//             }
//         }
//     }
// }

pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'sidpk/pipeline-app'
    }

    stages {
        stage('Checkout') {
            steps {
                echo 'Getting the code...'
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                echo 'Building the Docker image...'
                sh 'docker build -t ${DOCKER_IMAGE}:${BUILD_NUMBER} .'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                echo 'Pushing image to Docker Hub...'
                withCredentials([usernamePassword(credentialsId: 'dockerhub', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh '''
                        echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
                        docker push ${DOCKER_IMAGE}:${BUILD_NUMBER}
                    '''
                }
            }
        }
    }
}