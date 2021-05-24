pipeline {
    agent {
        label "master"
    }

    parameters {
        string(name: "password")
        string(name: "AWS_ACCOUNT")
        string(name: "AWS_ACCESS_KEY_ID")
        string(name: "AWS_SECRET_ACCESS_KEY")
    }

    stages {
        stage('Clone') {
            steps {
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 git clone https://github.com/eaneto/example-app /tmp/example-app"
            }
        }

        stage('Set up Environment') {
            steps {
                // Set up aws account
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 'export AWS_ACCOUNT=${params.AWS_ACCOUNT}'"
                // Set up aws access key id
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 'export AWS_ACCESS_KEY_ID=${params.AWS_ACCESS_KEY_ID}'"
                // Set up aws secret acess key
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 'export AWS_SECRET_ACCESS_KEY=${params.AWS_SECRET_ACCESS_KEY}'"
            }
        }

        stage('Build') {
            steps {
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 'cd /tmp/example-app && make build'"
            }
        }

        stage('Deploy') {
            steps {
                // Kill active process running the app, always return success
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 killall app || true"
                // Move the binary to /tmp
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 'cp /tmp/example-app/bin/app /tmp/app'"
                // launch binary on machine nohup
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 'nohup /tmp/app &'&"
                // Wait two seconds
                sh "sleep 2"
            }
        }

        stage("Cleanup") {
            steps {
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 rm -rf /tmp/example-app"
            }
        }
    }
}
