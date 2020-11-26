pipeline{

  agent { node 'jenkins-slave' } 
  options {
    timestamps()
  }
  environment {
    TimeStamp="${currentBuild.startTimeInMillis}"
    Service="${JOB_BASE_NAME}"
    ## gitlab webhook 回调功能
    ## Branch="${env.gitlabTargetBranch}"
  }

  stages {
    stage('Example Build') {
        steps {
            sh 'docker info'
            echo "${env.TimeStamp}"
            echo "${env.Service}"
        }
    }
}

}
