pipeline {
    agent {
        label "master"
    }

    parameters {
        string(name: "password")
    }

    stages {
        stage('Clone') {
            steps {
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 git clone https://github.com/eaneto/example-app /tmp/example-app"
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
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 \"nohup /tmp/app > /tmp/app.out &\"&"
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
