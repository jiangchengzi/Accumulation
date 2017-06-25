#!/bin/bash

set -e

declare -a needToUmountArray

declare -a needToUnmappedArray

pools=`ceph osd lspools|awk  'BEGIN{RS=",";ORS="\n"}{print $0}'|awk -F ' ' '{print $2}'`


unmap(){

	for i in ${needToUnmappedArray[@]};do
	
		rbd unmap ${i} --pool ${pool}
		
		echo -e "Unmap ${i} Successfully !\n" 
	
	done
}

umount(){

	for i in ${needToUmountArray[@]};do

                umount ${i}

                echo -e "Umount ${i} Successfully !\n" 

        done


}

start(){

	k8sMountedArray=`df -h |grep '/dev/rbd'|awk '{if (match($6,"kubernetes.io")) {print $1}}'`

	allMountedArray=`df -h |grep '/dev/rbd'|awk '{print $1}'`
		
        mappedArray=`rbd showmapped |grep ${pool}|grep '/dev/'|awk '{print $5}'`

        for i in ${k8sMountedArray[@]};do

                mountedDir=`mount | grep  ${i} |awk '{print $3}'`
                mountedPod=`mount |grep ${mountedDir}|grep -v ${i}|awk '{print $3}'|awk -F '/' '{print $6}'`

                if  ! docker ps |grep -q ${mountedPod} ;then

                        needToUmountArray=(${needToUmountArray[@]} ${i})

                fi

        done
        
	for i in ${mappedArray[@]};do
	
		flag=0
	
                for j in ${allMountedArray[@]};do
                        if [ ${j} == ${i} ] ;then
                                flag=1
				continue
                        fi
                done
               
		if [ ${flag} -eq 0 ];then
                        needToUnmappedArray=(${needToUnmappedArray[@]} ${i})
                fi
        done

        echo "need to umounted volume devices:" ${needToUmountArray[@]}
        echo "need to unmapped volume devices:" ${needToUnmappedArray[@]}

        read -p "Are you sure you want to release these resources? Y/n:" operation

        case $operation in
        "Y")
                umount 
                unmap
                ;;
        "n")
                ;;
        esac
}

select_pool(){
	echo "Please select a pool:"
	
	para= #assert option is valid
	
	select option in ${pools[@]} "Quit"
	do      
        	case $option in
       		"Quit") 
                	exit
                	;;
        	*) 
			if [ $option == $para ]; then
				echo "invalid input!"
				select_pool
			fi
                	echo "you choose $option"
                	echo "now starting to clean up unused resources in the storage pool...."
                	pool=$option
                	start
			select_pool
                	;;
        	esac
	done
}

select_pool
