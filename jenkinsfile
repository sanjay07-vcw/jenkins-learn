pipeline {
    agent any

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
                body: """Hello Karthikeyan,


Regards,  
Jenkins
"""
            )
        }

        failure {
            emailext(
                to: 'karthikeyanvelu777@gmail.com',
                subject: "❌ JOB FAILED",
                body: """Hello Karthikeyan,

Regards,  
Jenkins
"""
            )
        }
    }
}
