#!/bin/sh

# Run Idler locally and then this script can be run
# Check status of jenkins
# If idle, send a build request and check if it is unidled, check time to unidle it. check time taken for the build to get finish
UnIdleTest()
{
    curl -X GET 'http://localhost:8080/api/idler/unidle/ksagathi-preview-jenkins?openshift_api_url=https://api.starter-us-east-2a.openshift.com/'
    sleep 5m
    isIdle=`curl -X GET 'http://localhost:8080/api/idler/isidle/ksagathi-preview-jenkins?openshift_api_url=https://api.starter-us-east-2a.openshift.com/'|jq .is_idle`
    if [ $isIdle == 'true' ]
    then
        echo -e "\033[33;31mJenkins is still idle eventhough the request to unidle it was sent 5 min ago"
    else
        echo -e "\033[33;32mJenkins is unidle since the request to unidle it was sent 5 min ago"
    fi
    echo -e "\e[0m"    
}

isIdle=`curl -X GET 'http://localhost:8080/api/idler/isidle/ksagathi-preview-jenkins?openshift_api_url=https://api.starter-us-east-2a.openshift.com/'|jq .is_idle`
if [ $isIdle == 'true' ]
then
    echo "Jenkins is idle. lets unidle it."
    UnIdleTest
else
    echo "Jenkins is unidle. Let's idle it and then unidle it."
    curl -X GET 'http://localhost:8080/api/idler/idle/ksagathi-preview-jenkins?openshift_api_url=https://api.starter-us-east-2a.openshift.com/'
    UnIdleTest
fi