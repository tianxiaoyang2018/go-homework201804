#!/bin/sh
set -x
# putong服务的错误日志报警
# tail一分钟的内容判断报错数量
# 发送给指定负责人 + henry + 杜圣辰
# author: 杜圣辰
# 微服务上线替换putong服务后可下线该脚本
# ^-^

function alert () {
  srv=$1
  num=$2
  curHour="`date +%H`"
  line=$limit_line

#    putong-core-service#500#18510866603#钟子奇
#    putong-account-service#500#18510866603#钟子奇
#
#    putong-payment-service#50#18510866603,18010105606#钟子奇,程宇航
#    putong-payment-auto-renewing#50#18510866603,18010105606#钟子奇,程宇航
#    putong-affiliate-service#50#18840829171#丁毅
#
#    putong-abtest-service#50#18810541723#于乐
#
#    tantan-push-worker#1000#13532288219,15210905429#ken,葛永德
#    tantan-push-service#1000#13532288219,15210905429#ken,葛永德
#    putong-push-service#2000#13532288219,15210905429,15120000437#ken,葛永德,高思晨
#    putong-sms-promotion-offline-resend-for-guangdong#200#15120000437#高思晨
#    putong-sms-promotion-offline-resend#200#15120000437#高思晨
#    putong-sms-promotion-offline#200#15120000437#高思晨
#
#    putong-sender-service#200#17610968213#孙本新
#    putong-promotion-web#200#17610968213,13261226863#孙本新,尹航
#    putong-idle-user-sms-push#200#17610968213#孙本新
#
#    putong-admin-web#200#13070168818,13910555747#杨凯,马健
#
#    putong-media-detection-service#200#15810108260,13693250813#张兵,李晨阳
  #服务名#报警阈值#报警电话#负责人
  conf=$(echo "


    tantan-gateway-service#30#15011536289#于有琪
    tantan-user-counter-restapi#30#15011536289#于有琪
    tantan-user-counter-grpc#30#15011536289#于有琪
    tantan-moment-service#30#15011536289#于有琪
    tantan-user-grpc#30#15011536289#于有琪


  "| grep $srv | awk -F"#" '{print $2"\t"$3}' )
# putong-sms-service#100#15810108260#张兵
  if [ -z "$conf" ];then
    return 0
  fi

  line=$(echo "$conf" | awk '{print $1}')
  telephone=$(echo "$conf" | awk '{print $2}')

  if [ "$curHour" = "20" ] || [ "$curHour" = "21" ] ;then
    line=$(($line + 200));
  fi

  if [ $num -lt $line ];then
    return
  fi

  telephone=$telephone
  hostname=`cat /etc/hostname`
  ip=`hostname --all-ip-address`
  content=$(echo $hostname $ip $srv'最近1分钟有'$num'条错误日志，请赶快处理！')
  content=$(echo $content | sed 's/ //g')
  url="http://10.191.224.16:8080/sms_service?mobile="$telephone"&content="$content

  echo "$srv $num $telephone"
  curl $url
}

cur_path=$(dirname $(readlink -f $0))
log_file="/data/putong_log/putong_error.log"
alert_file="$cur_path/alert_by_sever.sh"
limit_line=30
stats_time=60

#加锁，独占式跑
lock_file="/tmp/.stats.file.lock"
exec 3> ${lock_file}
if ! flock -n 3 ;then
  echo "file is running,exit"
  exit 1
fi


waitmax $stats_time tail -n 0 -f $log_file  | awk -F"[\\\[: \\\]]" '{print $6}' | sort | uniq -c | sort -nr | while read err_num server_name
do
    alert $server_name $err_num
done

#sleep 1m
echo "finish"

