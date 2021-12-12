import csv
import time
import pandas as pd

import requests
from lxml import etree


def get_dd_content(u):
    headers = {
        "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) "
                      "Chrome/96.0.4664.55 Safari/537.36 Edg/96.0.1054.43",
    }
    content = requests.get(u, headers=headers)
    return content


def get_book(context):
    html = etree.HTML(context)

    titles, urls, imgs, prices = [], [], [], []
    titles.extend(html.xpath('//li/div[2]/a/img/@title'))
    urls.extend(html.xpath('//li/div[2]/a/@href'))
    imgs.extend(html.xpath('//li/div[2]/a/img/@src'))
    prices.extend(html.xpath(f'//li/div[7]/p[1]/span[1]/text()'))

    print(len(titles), len(urls), len(imgs), len(prices))
    return list(zip(titles, urls, imgs, prices))


def main():
    with open("data.csv", "a", encoding="utf-8") as f:
        writer = csv.writer(f)
        for j in range(1, 26):
            #  今日畅销榜
            # url = f"http://bang.dangdang.com/books/bestsellers/01.00.00.00.00.00-24hours-0-0-1-{j}"
            # 2020 畅销榜
            url = f"http://bang.dangdang.com/books/bestsellers/01.00.00.00.00.00-year-2020-0-1-{j}"
            result = get_dd_content(url)
            for i in get_book(result.text):
                print(i)
                writer.writerow(i)

            f.flush()
            time.sleep(5)


def get_category(url):
    html = get_dd_content(url).text
    html = etree.HTML(html)
    category = html.xpath('//div[@class="breadcrumb"]/a[2]/text()')[0]
    return category


def remove_duplicate(file):
    frame = pd.read_csv(file)
    data = frame.drop_duplicates(keep='first', inplace=False)
    data.to_csv(file, encoding='utf8')


def update_csv():
    f = open("data.csv", "r", encoding="utf-8")
    f1 = open("data1.csv", "w", encoding="utf-8")
    reader = csv.reader(f)
    writer = csv.writer(f1)

    for i in reader:
        i.append(get_category(i[1]))
        writer.writerow(i)
        print(i)
        f1.flush()
        time.sleep(5)

    f1.close()
    f.close()


if __name__ == '__main__':
    # main()
    # remove_duplicate("data.csv")
    update_csv()
