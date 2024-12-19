DAY=${1}

if [ -z "${DAY}" ]; then
    echo Day must be passed in first argument.
    exit 1
fi

cp -r ./template "./${DAY}"

sed -i "s/package template/package day${DAY}/g" "./${DAY}/part1.go"
sed -i "s/package template/package day${DAY}/g" "./${DAY}/part2.go"
