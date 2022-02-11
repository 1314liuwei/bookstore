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

arr=(${!tT[@]})
result=(${!tT[@]})

i=0
while(( $i < ${#arr[@]} ))
do
    max=$i
    j=$(($i + 1))
    
    while(( $j < ${#arr[@]} ))
    do
        tmp1=${result[$max]}
        tmp2=${result[j]}

        arr1=(${teams[$tmp1]})
        arr2=(${teams[$tmp2]})

        if [ ${arr1[4]} -gt ${arr2[4]} ]
        then
            max=$max
        elif [ ${arr1[4]} -lt ${arr2[4]} ]
        then
            max=$j
        else
            if [ ${arr1[5]} -gt ${arr2[5]} ]
            then
                max=$max
            elif [ ${arr1[5]} -lt ${arr2[5]} ]
            then
                max=$j
            fi
        fi

        let "j++"
    done

    tmp1=${result[$i]}
    tmp2=${result[$max]}
    result[$max]=$tmp1
    result[$i]=$tmp2

    let "i++"
done


echo "Mini-league with ${#result[*]} teams"
echo -e "Rank\tTeam\tG\tW\tD\tL\tPoint\tGF\tGA\tGD"
for i in ${!result[@]}
do
    arr=(${teams[${result[$i]}]})
    echo -e "$(($i + 1))\t${tT[${result[$i]}]}\t${arr[0]}\t${arr[1]}\t${arr[2]}\t${arr[3]}\t${arr[4]}\t${arr[5]}\t${arr[6]}\t${arr[7]}"
done


