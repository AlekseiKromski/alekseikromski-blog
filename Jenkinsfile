pipeline {
    environment {
        dockerImage = ''
    }
    agent any
    stages {
        stage('Build image') {
            steps {
               script{
                dockerImage = docker.build("localhost:5000/blog:testing")
               }
            }
        }
        stage('Push image') {
            steps {
               script {
                 withDockerRegistry([ credentialsId: "docker-registry", url: "http://localhost:5000" ]) {
                    dockerImage.push()
                 }
               }
            }
        }
    }
}
