:: golangci-lint run
set _function="SmDelegatorCoverOracle"
echo %_function%
set GOOS=linux
cd src
gofmt -s -w . && ^
go build -o ..\dist\main main.go && ^
C:%HOMEPATH%\go\bin\build-lambda-zip.exe -o "%cd%\..\dist\main.zip" "%cd%\..\dist\main" && ^
aws s3 cp --profile=mb "%cd%\..\dist\main.zip" s3://stakemovr.com/cache/jars/%_function%.main.zip && ^
aws lambda update-function-code --profile=mb --function-name %_function% --s3-bucket stakemovr.com --s3-key cache/jars/%_function%.main.zip
cd ..