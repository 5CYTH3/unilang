powershell Set-ExecutionPolicy RemoteSigned -scope CurrentUser
powershell iwr -useb get.scoop.sh
powershell scoop install go
powershell go install golang.org/dl/go1.18rc1@latest
powershell go1.18rc1 download
powershell scoop install make