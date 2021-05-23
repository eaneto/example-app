pipeline {
    agent {
        label master
    }

    stages {
        stage('Clone') {
            steps {
                git 'https://github.com/eaneto/example-app.git'
            }
        }

        stage('Build') {
            steps {
                sh 'make build'
            }
        }
    }
}
