node() {
  stage('克隆代码') {
    echo "1.Clone Stage and Prepare"
    // git credentialsId: 'e63825bc-e13c-4734-a3cd-2e33d81a2c4d', url: 'git@github.com:tomlinux/jenkins-demo.git'
    checkout scm
    echo "${env.GIT_BRANCH}"
    sh "exit 3"
    script {
          build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
          if (env.GIT_BRANCH != 'master') {
              build_tag = "${env.GIT_BRANCH}-${build_tag}"
          }
    }
    echo "${build_tag}"
  }
  stage('测试项目') {
    echo "2.Test Stage"
    // for(e in env){
    //   echo e + " is " + ${e}}
    // }
  }
  stage('构建镜像') {
    echo "3.Build Docker Image Stage"
    sh "docker build -t  ccr.ccs.tencentyun.com/development/jenkins-demo:${build_tag} ."
  }
  stage('上传镜像') {
      echo "4.Push Docker Image Stage"
      withCredentials([usernamePassword(credentialsId: '482f25c0-a6a0-48de-bd44-67242f69e8c1', passwordVariable: 'QclondRegistryPassword', usernameVariable: 'QclondRegistryUser')]) {
          sh "docker login  ccr.ccs.tencentyun.com -u ${QclondRegistryUser} -p ${QclondRegistryPassword}"
          sh "docker push  ccr.ccs.tencentyun.com/development/jenkins-demo:${build_tag}"
          sh "docker rmi  ccr.ccs.tencentyun.com/development/jenkins-demo:${build_tag}"
      }
  }
  stage('发布') {
    echo "5. Deploy Stage"
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
    sh "sed -i 's/<BRANCH_NAME>/${env.GIT_BRANCH}/' k8s.yaml"
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


