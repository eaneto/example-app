pipeline {
    agent {
        label master
    }

    stages {
        stage('Clone') {
            steps {
                git branch: "main", url "https://github.com/eaneto/example-app.git"
            }
        }

        stage('Build') {
            steps {
                sh 'make build'
            }
        }
    }
}
