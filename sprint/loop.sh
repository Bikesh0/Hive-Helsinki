n=$1

if [ "$n" -gt 100 ]; then
    n=100
fi

i=1
while [ "$i" -le "$n" ]; do
    echo "This is loop number $i"
    i=$((i + 1))
done
