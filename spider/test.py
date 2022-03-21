import threading
import time
import csv


if __name__ == '__main__':
    with open("data.csv", "r", encoding="utf-8") as f:
        reader = csv.reader(f)
        for i in reader:

            sql = f"INSERT IGNORE INTO `books` (`name`, `cover`, `price`, `category`) VALUES ('{i[0]}', '{i[2]}', {i[3]}, '{i[4]}');"
            print(sql)
    
    