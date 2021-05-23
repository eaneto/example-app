pipeline {
    agent {
        label master
    }

    stages {
        stage('List') {
            steps {
                sh 'ls -la'
            }
        }

        stage('Build') {
            steps {
                sh 'make build'
            }
        }
    }
}
