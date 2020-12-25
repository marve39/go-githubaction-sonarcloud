# go-githubaction-sonarcloud
Sample for CI flow using github action and sonar cloud for go lang

Release Pipeline : 
- Code Scanner (Code security scan, Code test coverage) -> SonarCloud
- Unit Test -> SonarCloud
- Auto Tagging Release

Build Pipeline :
- Code Scanner (Code security scan, Code test coverage) -> SonarQube
- Unit Test -> ReportPortal 
- Build Image 
- Image Scanner -> Trivy

Deploy Pipeline : 
- Deploy Image -> Docker
