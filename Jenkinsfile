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
                dockerImage = docker.build("docker.alekseikromski.com/blog:" + tag)
               }
            }
        }
        stage('Push image') {
            steps {
               script {
                 withDockerRegistry([ credentialsId: "docker-registry", url: "https://docker.alekseikromski.com" ]) {
                    dockerImage.push()
                 }
               }
            }
        }
    }
}

def getTag() {
    def now = new Date()
    return now.format("yyMMddHHmm", TimeZone.getTimeZone('UTC'))
}
