pipeline {
    agent {
        label master
    }

    stages {
        stage('List') {
            steps {
                sh 'ls'
            }
        }

        stage('Build') {
            steps {
                sh 'make build'
            }
        }
    }
}
