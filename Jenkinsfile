pipeline {
    agent {
        label "master"
    }

    stages {
        stage('Build') {
            steps {
                sh 'make build'
            }
        }
    }
}
