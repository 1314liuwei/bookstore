import threading
import time


def readFile(file):
    while True:
        line = file.readline()
        print(line)
        time.sleep(3)
        if not line:
            break


if __name__ == '__main__':
    with open("test.txt", "r", encoding="utf-8") as f:
        tasks = [threading.Thread(target=readFile, args=(f, )) for i in range(5)]
        for t in tasks:
            t.start()

        for t in tasks:
            t.join()
