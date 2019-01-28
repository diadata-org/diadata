checkBinary() {
which $1
if [ $? -ne 0 ]
then
echo missing $1
exit
fi


}
FILE="$1"
if [ ! -f $FILE ]
then
 	echo "File $FILE does not exist"; exit
fi
checkBinary jq
checkBinary jsonlint-py

jsonlint-py -f $FILE  | jq '.Coins|=sort_by(.Symbol)|.Coins|=sort_by(.ForeignName)' > /tmp/$$.txt
cp /tmp/$$.txt $FILE
