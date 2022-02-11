#!/bin/bash

declare -A teams
declare -A tT

while read line
do
    arr=($line)
    tmp=(0 0 0 0 0 0 0 0)
    teams[${arr[0]}]=${tmp[@]}
    tT[${arr[0]}]=${arr[1]}
done < top.dat

files=($*)
for i in ${files[@]:1:$#}
do
    while read line
    do
        split=($line)
        arr1=(${teams[${split[0]}]})
        arr2=(${teams[${split[2]}]})

        if [ ${split[1]} -gt ${split[3]} ]
        then
                # 1
                tmp=$((${arr1[0]} + 1)) # G
                arr1[0]=$tmp

                tmp=$((${arr1[1]} + 1)) # W
                arr1[1]=$tmp

                arr1[2]=${arr1[2]} # D

                arr1[3]=${arr1[3]} # L

                tmp=$((${arr1[4]} + 3)) # P
                arr1[4]=$tmp

                tmp=$((${arr1[5]} + ${split[1]}))
                arr1[5]=$tmp # GF

                tmp=$((${arr1[6]} + ${split[3]}))
                arr1[6]=$tmp # GA

                tmp=$((${arr1[5]} - ${arr1[6]})) # W
                arr1[7]=$tmp

                teams[${split[0]}]=${arr1[@]}


                # 2
                tmp=$((${arr2[0]} + 1)) # G
                arr2[0]=$tmp

                arr2[1]=${arr2[1]} # W

                arr2[2]=${arr2[2]} # D

                tmp=$((${arr2[3]} + 1))
                arr2[3]=$tmp # L

                tmp=$((${arr2[4]} + 0)) # P
                arr2[4]=$tmp

                tmp=$((${arr2[5]} + ${split[3]}))
                arr2[5]=$tmp # GF

                tmp=$((${arr2[6]} + ${split[1]}))
                arr2[6]=$tmp # GA

                tmp=$((${arr2[5]} - ${arr2[6]})) # W
                arr2[7]=$tmp

                teams[${split[2]}]=${arr2[@]}
        elif [ ${split[1]} -eq ${split[3]} ]
        then
                # 1
                tmp=$((${arr1[0]} + 1)) # G
                arr1[0]=$tmp

                arr1[1]=${arr1[1]}

                tmp=$((${arr1[2]} + 1))
                arr1[2]=$tmp # D

                arr1[3]=${arr1[3]} # L

                tmp=$((${arr1[4]} + 1)) # P
                arr1[4]=$tmp

                tmp=$((${arr1[5]} + ${split[1]}))
                arr1[5]=$tmp # GF

                tmp=$((${arr1[6]} + ${split[3]}))
                arr1[6]=$tmp # GA

                tmp=$((${arr1[5]} - ${arr1[6]})) # W
                arr1[7]=$tmp

                teams[${split[0]}]=${arr1[@]}

                # 2
                tmp=$((${arr2[0]} + 1)) # G
                arr2[0]=$tmp

                arr2[1]=${arr2[1]} # W

                tmp=$((${arr2[2]} + 1))
                arr2[2]=$tmp # D

                arr2[3]=${arr2[3]} # L

                tmp=$((${arr2[4]} + 1)) # P
                arr2[4]=$tmp

                tmp=$((${arr2[5]} + ${split[3]}))
                arr2[5]=$tmp # GF

                tmp=$((${arr2[6]} + ${split[1]}))
                arr2[6]=$tmp # GA

                tmp=$((${arr2[5]} - ${arr2[6]})) # W
                arr2[7]=$tmp

                teams[${split[2]}]=${arr2[@]}
        else
                # 1
                tmp=$((${arr1[0]} + 1)) # G
                arr1[0]=$tmp

                arr1[1]=${arr1[1]}

                arr1[2]=${arr1[2]} # D

                tmp=$((${arr1[3]} + 1))
                arr1[3]=$tmp # L

                tmp=$((${arr1[4]} + 0)) # P
                arr1[4]=$tmp

                tmp=$((${arr1[5]} + ${split[1]}))
                arr1[5]=$tmp # GF

                tmp=$((${arr1[6]} + ${split[3]}))
                arr1[6]=$tmp # GA

                tmp=$((${arr1[5]} - ${arr1[6]})) # W
                arr1[7]=$tmp

                teams[${split[0]}]=${arr1[@]}

                # 2
                tmp=$((${arr2[0]} + 1)) # G
                arr2[0]=$tmp

                tmp=$((${arr2[1]} + 1))
                arr2[1]=$tmp # W

                arr2[2]=${arr2[2]} # D

                arr2[3]=${arr2[3]}

                tmp=$((${arr2[4]} + 3)) # P
                arr2[4]=$tmp

                tmp=$((${arr2[5]} + ${split[3]}))
                arr2[5]=$tmp # GF

                tmp=$((${arr2[6]} + ${split[1]}))
                arr2[6]=$tmp # GA

                tmp=$((${arr2[5]} - ${arr2[6]})) # W
                arr2[7]=$tmp

                teams[${split[2]}]=${arr2[@]}
        fi
        done < $i
done

for i in ${!teams[@]}
do
    echo $i ${teams[$i]}
done

for i in ${!tT[@]}
do
    echo $i ${tT[$i]}
done

arr=(${!tT[@]})
result=(${!tT[@]})
echo "Result"
echo "${result[@]}"
i=0
while(( $i<= ${#arr[@]} ))
do
    tmp1=${arr[i]}
    j=$(($i + 1))
    let "i++"


    while(( $j < ${#arr[@]} ))
    do
        max=$tmp1
        tmp2=${arr[j]}
        let "j++"

        arr1=(${teams[$tmp1]})
        arr2=(${teams[$tmp2]})
        echo "${arr1[4]} ${arr2[4]}"
        if [ ${arr1[4]} -gt ${arr2[4]} ]
        then
            max=$tmp1
        elif [ ${arr1[4]} -lt ${arr2[4]} ]
        then
            max=$tmp2
            echo $max
        else
            if [ ${arr1[5]} -gt ${arr2[5]} ]
            then
                max=$tmp1
            elif [ ${arr1[5]} -lt ${arr2[5]} ]
            then
                max=$tmp2
            fi
        fi
    done

    result[$i]=$max
done

echo "Result"
for i in ${!result[@]}
do
    echo $i ${result[$i]}
done