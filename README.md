[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-white.svg)](https://sonarcloud.io/dashboard?id=marve39_go-githubaction-sonarcloud) ![](https://github.com/marve39/go-githubaction-sonarcloud/workflows/Main%20Workflow/badge.svg)
# go-githubaction-sonarcloud
Sample for CI flow using github action and sonar cloud for go lang

Release Pipeline : 
- Code Scanner (Lint scanner, Code depedency scanner, Code security scan, Code test coverage) -> SonarCloud
- Unit Test -> SonarCloud
- Auto Tagging Release

Build Pipeline :
- Code Scanner (Code security scan, Code test coverage) -> SonarQube
- Unit Test -> ReportPortal 
- Build Image 
- Image Scanner -> Trivy

Deploy Pipeline : 
- Deploy Image -> Docker

Reference : https://riggaroo.dev/using-github-actions-to-automate-our-release-process/