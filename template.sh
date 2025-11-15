#!/usr/bin/env bash

endpoint=
creatorId="henyathegenius"
firstId=
jq_arg=

while [[ -n "$1" ]]; do
	case "$1" in
		--listCreator)
			endpoint=listCreator
			;;

		--paginateCreator)
			endpoint=paginateCreator
			;;
		--creatorId)
			shift
			creatorId=$1
			;;

		--firstId)
			shift
			firstId=$1
			;;
		--jq)
			shift
			jq_arg="$1"
			;;
		*)
			echo "argument $1 is not valid " >&2
			exit 1
			;;
	esac
	shift
done

chatterino_args=(
	--get 
	-H 'Origin: https://henyathegenius.fanbox.cc'
)

paginate_creator=(
	--data-urlencode "creatorId=$creatorId"
	--data-urlencode "sort=newest"
	'https://api.fanbox.cc/post.paginateCreator?'
)

list_creator=(
	--data-urlencode 'firstPublishedDatetime=2025-11-08%2012%3A00%3A00'
	--data-urlencode "firstId=$firstId"
	--data-urlencode 'limit=10'
	--data-urlencode "creatorId=$creatorId"
	--data-urlencode "sort=newest"
	'https://api.fanbox.cc/post.listCreator'
)


case "$endpoint" in
	listCreator)
		chatterino_args+=("${list_creator[@]}")
		;;

	paginateCreator)
		chatterino_args+=("${paginate_creator[@]}")
		;;
	*)
		cat <<- 'EOF' 
		argument required.
			--listCreator
			--paginateCreator
		EOF
		exit
		;;
esac


curl "${chatterino_args[@]}" | jq "${jq_arg}"

