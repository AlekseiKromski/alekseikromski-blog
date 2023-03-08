pipeline {
    agent any
    stages {
        stage('create docker image') {
            steps {
                sh "docker build . --tag docker.alekseikromski.com/blog"
            }
        }
        stage('push docker image') {
            steps {
                sh "docker push docker.alekseikromski.com/blog"
            }
        }
    }
}