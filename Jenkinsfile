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

environment {
  TimeStamp="${currentBuild.startTimeInMillis}"
  Service="${JOB_BASE_NAME}"
  Branch_Name= 'master'
  //gitlab webhook 回调功能
  Branch="${env.gitlabTargetBranch}"
}  


node('jenkins-slave') {

      stage('克隆代码') {
        //PrintMes("1.代码克隆和准备阶段", "green")
        // git credentialsId: 'e63825bc-e13c-4734-a3cd-2e33d81a2c4d', url: 'git@github.com:tomlinux/jenkins-demo.git'
 
          // PrintMes("$Service","blue")  错误
          echo "===================="
          printenv()
          checkout scm
          script {
                BRANCH_NAME_TAG = sh(returnStdout: true, script: 'git rev-parse --abbrev-ref HEAD').trim()
                build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
                Branch_Name = 'master'
                if (Branch_Name != 'master') {
                    build_tag = "${Branch_Name}-${build_tag}"
                }
          }       
          PrintMes($build_tag,"blue")   
        
               

      }
      stage('测试项目') {
      

         PrintMes("2.测试项目", "green")

        

       
      }
      stage('构建镜像') {
    
        PrintMes("3.Build Docker Image Stage", "green")
      
        
      }
      stage('上传镜像') {
         
          PrintMes("4.Push Docker Image Stage", "green")
        
        
      }
      stage('发布') {
          
            PrintMes("5. Deploy Stage", "green")
          
        

      } 

}



