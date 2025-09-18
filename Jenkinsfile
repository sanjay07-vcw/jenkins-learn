pipeline {
       agent any
    tools {
        go 'Golang'
    }

    environment {
        EMAIL_RECIPIENT = 'karthikeyanvelu777@gmail.com'
    }

    stages {
            stage('Check Go') {
            steps {
                sh 'go version || echo "Go is NOT installed on this Jenkins server"'
            }
        }
        stage('Build') {
            steps {
                sh 'go build -o myapp main.go'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploying Go app...'
            }
        }
    }

    post {
        success {
            emailext(
                to: "${EMAIL_RECIPIENT}",
                subject: '✅ JOB SUCCESS',
                body: '''Hello Karthikeyan,

Your Jenkins job completed successfully. 🎉

Regards,
Jenkins
'''
            )
        }

        failure {
            emailext(
                to: "${EMAIL_RECIPIENT}",
                subject: '❌ JOB FAILED',
                body: '''Hello Karthikeyan,

Your Jenkins job has failed. ❌
Please check the build logs for details.

Regards,
Jenkins
'''
            )
        }

        always {
            echo 'Pipeline finished.'
        }
    }
}
