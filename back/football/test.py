def top():
    with open("top.dat", "r") as f:
        data = f.readlines()
        result = {}
        for i in data:
            split = i.split()
            result[split[0]] = split[1]
        return result


def score(teams):
    result = dict(zip(teams.keys(), ([0 for _ in range(8)] for _ in range(len(teams)))))
    with open("score1", "r") as f:
        data = f.readlines()
        for i in data:
            split = i.split()
            if int(split[1]) > int(split[3]):
                result[split[0]][0] += 1  # 总场次
                result[split[0]][1] += 1  # 赢
                # result[split[0]][2] += 1  # 平
                # result[split[0]][3] += 1  # 输
                result[split[0]][4] += 3  # 得分
                result[split[0]][5] += int(split[1])  # 进球
                result[split[0]][6] += int(split[3])  # 失球
                result[split[0]][6] = int(result[split[0]][5]) - int(result[split[0]][6])  # 差

                result[split[2]][0] += 1  # 总场次
#                 result[split[2]][1] += 1  # 赢
                # result[split[2]][2] += 1  # 平
                result[split[2]][3] += 1  # 输
#                 result[split[2]][4] +=   # 得分
                result[split[2]][5] += int(split[3])  # 进球
                result[split[2]][6] += int(split[1])  # 失球
                result[split[2]][6] = int(result[split[2]][5]) - int(result[split[2]][6])  # 差
            elif int(split[1]) == int(split[3]):
                result[split[0]][0] += 1  # 总场次
#                 result[split[0]][1] += 1  # 赢
                result[split[0]][2] += 1  # 平
                # result[split[0]][3] += 1  # 输
                result[split[0]][4] += 1  # 得分
                result[split[0]][5] += int(split[1])  # 进球
                result[split[0]][6] += int(split[3])  # 失球
                result[split[0]][6] = int(result[split[0]][5]) - int(result[split[0]][6])  # 差

                result[split[2]][0] += 1  # 总场次
#                 result[split[2]][1] += 1  # 赢
                result[split[2]][2] += 1  # 平
                # result[split[2]][3] += 1  # 输
                result[split[2]][4] += 1  # 得分
                result[split[2]][5] += int(split[3])  # 进球
                result[split[2]][6] += int(split[1])  # 失球
                result[split[2]][6] = int(result[split[2]][5]) - int(result[split[2]][6])  # 差
            else:
                result[split[0]][0] += 1  # 总场次
#                 result[split[0]][1] += 1  # 赢
                # result[split[0]][2] += 1  # 平
                result[split[0]][3] += 1  # 输
#                 result[split[0]][4] += 3  # 得分
                result[split[0]][5] += int(split[1])  # 进球
                result[split[0]][6] += int(split[3])  # 失球
                result[split[0]][6] = int(result[split[0]][5]) - int(result[split[0]][6])  # 差

                result[split[2]][0] += 1  # 总场次
                result[split[2]][1] += 1  # 赢
                # result[split[2]][2] += 1  # 平
                # result[split[2]][3] += 1  # 输
                result[split[2]][4] += 3  # 得分
                result[split[2]][5] += int(split[3])  # 进球
                result[split[2]][6] += int(split[1])  # 失球
                result[split[2]][6] = int(result[split[2]][5]) - int(result[split[2]][6])  # 差


        return result


if __name__ == '__main__':
    teams = top()
    print(score(teams))
