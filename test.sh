HOST=localhost
PORT=8080

NUMBER_OF_CONNECTIONS=1000


for i in $(seq 1 $NUMBER_OF_CONNECTIONS); do
	echo "$i" | nc -q 0 $HOST $PORT  &
done


