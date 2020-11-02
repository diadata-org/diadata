
DB="dia"

#CREATE RETENTION POLICY "30d" ON "$DB" DURATION 30d REPLICATION 1


for fnc in mean sum
	do
	for a in 5m 30m 1h 4h 1d 1w
		do
		QUERY="SELECT $fnc(\"value\") AS value INTO \"a_year\".\"filters_${fnc}_$a\" FROM \"filters\" GROUP BY time($a), symbol, filter, exchange"
		echo $QUERY
		#echo  "CREATE CONTINUOUS QUERY \"cq_filters_${fnc}_$a\" ON \"$DB\" BEGIN  $QUERY END"
	done
done

exit

for fnc in mean sum
	do
	for a in 5m 30m 1h 4h 1d 1w
		do
		SELECT * from "a_year"."filters_mean_5m" orderby asc
		echo  "select count(value) from \"a_year\".\"filters_${fnc}_$a\""
		echo  "SELECT * from \"a_year\".\"filters_${fnc}_$a\""
		echo  "SELECT * INTO  \"a_year\".\"filters_${fnc}_$a\" FROM filters WHERE time > now() - 100w and time < now() - 90w GROUP BY *"
	done
done
