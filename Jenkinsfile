pipeline {
    agent any

    tools {
        go 'Golang 1.20.5'
    }

    stages {
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
                to: 'karthikeyanvelu777@gmail.com',
                subject: "✅ JOB SUCCESS",
                body: '''Hello Karthikeyan,

Your Jenkins job completed successfully. 🎉

Regards,
Jenkins
'''
            )
        }

        failure {
            emailext(
                to: 'karthikeyanvelu777@gmail.com',
                subject: "❌ JOB FAILED",
                body: '''Hello Karthikeyan,

Your Jenkins job has failed. Please check the logs. ❌

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
