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
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 'cp /tmp/example-app/bin/app /tmp/app'"
            }
        }

        stage('Deploy') {
            steps {
                // Kill active process running the app
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 killall app"
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
