#!/bin/bash
MY_URL=https://hooks.slack.com/services/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
MY_DATE=`date +"%Y%m%d-%H%M%S"`
echo $MY_DATE
MY_MSG="payload={\"text\":\"`hostname`:$MY_DATE\"}"
echo $MY_MSG
#curl -X POST -d "payload={\"text\":\" test test test \"}" $MY_URL
curl -X POST -d $MY_MSG $MY_URL
