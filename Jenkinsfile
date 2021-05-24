pipeline {
    agent {
        label "master"
    }

    parameters {
        string(name: "password", description: "Deploy Password")
        string(name: "accessKey", description: "AWS Access Key")
        string(name: "secretKey", description: "AWS Secret Key")
        string(name: "account"  , description: "AWS Account Number")
    }

    stages {
        stage('Clone') {
            steps {
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 git clone https://github.com/eaneto/example-app /tmp/example-app"
            }
        }

        stage('Build') {
            steps {
                sh "make build"
            }
        }

        stage('Deploy') {
            steps {
                // Kill active process running the app, always return success
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 killall app || true"
                // Move the binary to /tmp
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 'cp /tmp/example-app/bin/app /tmp/app'"
                // launch binary on machine nohup
                sh "sshpass -p '${params.password}' ssh -oStrictHostKeyChecking=no vagrant@192.168.33.12 \" nohup /tmp/app -accessKey=${accessKey}> -secretKey=${secretKey} -account=${account} > /tmp/app.out &\"&"
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
