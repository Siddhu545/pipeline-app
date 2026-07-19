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
        stage('Update Manifest in Git') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'github-token', usernameVariable: 'GIT_USER', passwordVariable: 'GIT_TOKEN')]) {
                    sh '''
                        # Update the image tag in deployment.yaml to the new build number
                        sed -i "s|${DOCKER_IMAGE}:.*|${DOCKER_IMAGE}:${BUILD_NUMBER}|" k8s/deployment.yaml

                        # Configure git identity for the commit
                        git config user.email "jenkins@pipeline.local"
                        git config user.name "Jenkins CI"

                        # Commit and push the change
                        git add k8s/deployment.yaml
                        git commit -m "Update image to ${BUILD_NUMBER} [ci skip]"
                        git push https://${GIT_USER}:${GIT_TOKEN}@github.com/Siddhu545/pipeline-app.git HEAD:main
                    '''
                }
            }
        }
    }
}