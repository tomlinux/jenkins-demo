#!groovy
def PrintMes(value,color){
    colors = ['red'   : "\033[40;31m >>>>>>>>>>>${value}<<<<<<<<<<< \033[0m",
              'blue'  : "\033[47;34m ${value} \033[0m",
              'green' : "[1;32m>>>>>>>>>>${value}>>>>>>>>>>[m",
              'green1' : "\033[40;32m >>>>>>>>>>>${value}<<<<<<<<<<< \033[0m" ]
    ansiColor('xterm') {
        println(colors[color])
    }
}

pipeline{

  agent { node 'jenkins-slave' } 
  options {
    timestamps()
  }
  environment {
    TimeStamp="${currentBuild.startTimeInMillis}"
    Service="${JOB_BASE_NAME}"
    BRANCH_NAME='master'
    //gitlab webhook 回调功能
    Branch="${env.gitlabTargetBranch}"
  }

  stages {
      stage('克隆代码') {
        //PrintMes("1.代码克隆和准备阶段", "green")
        // git credentialsId: 'e63825bc-e13c-4734-a3cd-2e33d81a2c4d', url: 'git@github.com:tomlinux/jenkins-demo.git'
        checkout scm
        script {
              BRANCH_NAME_TAG = sh(returnStdout: true, script: 'git rev-parse --abbrev-ref HEAD').trim()
              build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
              if ($BRANCH_NAME != 'master') {
                  build_tag = "${BRANCH_NAME}-${build_tag}"
              }
        }
        PrintMes($build_tag,"blue")
       

      }
      stage('测试项目') {

        // PrintMes("2.测试项目", "green")
        // for(e in env){
        //   echo e + " is " + ${e}}
        // }
      }
      stage('构建镜像') {
        //PrintMes("3.Build Docker Image Stage", "green")
        sh "docker build -t  ccr.ccs.tencentyun.com/development/jenkins-demo:${build_tag} ."
      }
      stage('上传镜像') {
          //PrintMes("4.Push Docker Image Stage", "green")
          withCredentials([usernamePassword(credentialsId: '482f25c0-a6a0-48de-bd44-67242f69e8c1', passwordVariable: 'QclondRegistryPassword', usernameVariable: 'QclondRegistryUser')]) {
              sh "docker login  ccr.ccs.tencentyun.com -u ${QclondRegistryUser} -p ${QclondRegistryPassword}"
              sh "docker push  ccr.ccs.tencentyun.com/development/jenkins-demo:${build_tag}"
              sh "docker rmi  ccr.ccs.tencentyun.com/development/jenkins-demo:${build_tag}"
          }
      }
      stage('发布') {
        //PrintMes("5. Deploy Stage", "green")
        def userInput = input(
          id: 'userInput',
          message: 'Choose a deploy environment',
          parameters: [
              [
                  $class: 'ChoiceParameterDefinition',
                  choices: "Dev\nQA\nProd",
                  name: 'Env'
              ]
          ]
        )
        echo "This is a deploy step to ${userInput}"
        sh "sed -i 's/<BUILD_TAG>/${build_tag}/' k8s.yaml"
        sh "sed -i 's/<BRANCH_NAME>/${env.BRANCH_NAME}/' k8s.yaml"
        if (userInput == "Dev") {
        // deploy dev stuff
        } else if (userInput == "QA"){
        // deploy qa stuff
        } else {
        // deploy prod stuff
        }
        echo "发布成功"
        // sh "kubectl apply -f k8s.yaml -n default"
      }   

  }
}


