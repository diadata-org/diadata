for a in `ls ../config/*.json`
do
	echo $a
	./jsonLintConfigs.sh $a
done
