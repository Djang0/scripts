#!/bin/bash
PS3='Choose action (or ctrl-c to quit): '
select type in "ssh" "sftp" "ssh-copy-id" "quit"
do
   	select item in \
	"administrator@crios-linux" \
	"git@crios-linux" \
	"phoenix@crios-linux" \
	"administrator@activemq.belgomedia.local" \
	"administrator@solr.belgomedia.local" \
	"administrator@test.occ.telepro.be" \
	"administrator@test.sso.telepro.be" \
	"osmc@lreenaers.ddns.net" \
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
