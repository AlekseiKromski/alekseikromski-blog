pipeline {
    environment {
        dockerImage = ''
        tag = getTag()
    }
    agent any
    stages {
        stage('Build image') {
            steps {
               script{
                dockerImage = docker.build("localhost:5000/blog:" + tag)
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
        stage('Run docker container') {
            steps {
                script {
                    sh "docker compose up -d"
                }
            }
        }
    }
}

def getTag() {
    def now = new Date()
    return now.format("yyMMddHHmm", TimeZone.getTimeZone('UTC'))
}
