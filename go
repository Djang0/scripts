#!/bin/sh
PS3='Choose action (or ctrl-c to quit): '
select type in "ssh" "sftp" "ssh-copy-id" "quit"
do
   	select item in \
	"administrator@crios-linux" \
	"git@crios-linux" \
	"administrator@172.31.6.51" \
	"phoenix@crios-linux" \
	"osmc@lreenaers.ddns.net -p 206" \
	"osmc@192.168.1.63"\
  "root@buggenie.belgomedia.local"
	do
		case "$type" in
           ssh ) ssh $item ;;
           sftp ) sftp $item ;;
           ssh-copy-id ) ssh-copy-id $item ;;
           quit ) break ;;
           "") please select one of the above or press ctrl-c ;;
       	esac
       	exit 0
	done
	exit 0
done
