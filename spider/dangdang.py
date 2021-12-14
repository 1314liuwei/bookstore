import csv
import random
import time
import pandas as pd

import requests
from lxml import etree


def get_dd_content(u):
    headers = {
        "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36 Edg/96.0.1054.53",
        'Cookie': "dest_area=country_id%3D9000%26province_id%3D111%26city_id%20%3D0%26district_id%3D0%26town_id%3D0; __permanent_id=20211206090856824122559833395345413; secret_key=aadf6094c0b768c3bfb7fabbfa33ec78; ddscreen=2; __rpm=%7Clogin_page...1639444890016; login.dangdang.com=.AYH=&.ASPXAUTH=O8LIA/OA+TNag9Y8cJ6j/cY8ag5b+o3kyE2tXmI3HkFycOpaqbmVVA==; dangdang.com=email=MTU1MjA1MzIyMDk3ODk5NkBkZG1vYmlscGhvbmVfX3VzZXIuY29t&nickname=yrG85MLD0NDV3w==&display_id=5268210381764&customerid=0ZFDdVFUI7b+H+J98d4LqA==&viptype=&show_name=155%2A%2A%2A%2A2209; ddoy=email=1552053220978996%40ddmobilphone__user.com&nickname=%CA%B1%BC%E4%C2%C3%D0%D0%D5%DF&validatedflag=0&agree_date=1; sessionID=pc_1143e32f454fd78e2b56ed758c1e4f053da5cef4db3885d8e28a4ef94de35c44; bind_cust_third_id=ocil5uGRgr3NWKP54hYjN_diNZvE; bind_custid=723210923; tx_open_id=oqh4kuOTki4mQ0dRTk76Ci1NNjsM; tx_nickname=yrG85MLD0NDV3w==; tx_figureurl=https://thirdwx.qlogo.cn/mmopen/vi_32/9Otwibfa5VXS3uLUPhF7QEn5GdA5VbPkGbDRviclKn1Xx7Po1ia53QGibFssw015pb3TYTCiaJeTjk1MyP7aiaXb2p6Q/132; order_follow_source=-%7C-O-123%7C%2311%7C%23login_third_weixin%7C%230%7C%23; LOGIN_TIME=1639464359676; __visit_id=20211214144559682376872721641487482; __out_refer=; __trace_id=20211214144621128430315625037497252",
    }
    # proxies = {'http': 'http://10.0.0.10:1080', 'https': 'http://10.0.0.10:1080'}
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
    f1 = open("data.csv", "a", encoding="utf-8")
    reader = csv.reader(f)
    writer = csv.writer(f1)

    start = 490
    count = 0
    for i in reader:
        count += 1
        if count >= start:
            try:
                i.append(get_category(i[1]))
                writer.writerow(i)
                print(i)
                f1.flush()
                time.sleep(random.randint(20, 40))
            except Exception as e:
                print(i[1])

    f1.close()
    f.close()


if __name__ == '__main__':
    # main()
    # remove_duplicate("data.csv")
    update_csv()
