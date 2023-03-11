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
    }
}

def getTag() {
    def date = new Date()
    def sdf = new SimpleDateFormat("MM/dd/yyyy HH:mm:ss")
    return sdf.format(date)
}