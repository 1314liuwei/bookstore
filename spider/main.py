import re
import time

import MySQLdb
import requests
from lxml import etree


def get_content(u):
    proxies = {'http': 'http://127.0.0.1:54321', 'https': 'http://127.0.0.1:54321'}
    headers = {
        "cookie": "proxiesNotWorking=; domainsNotWorking=4lib.org%2C2lib.org",
        "referer": "https://zh.b-ok.cc/",
        "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) "
                      "Chrome/96.0.4664.55 Safari/537.36 Edg/96.0.1054.43",
    }
    content = requests.get(u, proxies=proxies, headers=headers)
    return content


def get_info(context):
    html = etree.HTML(context)
    info = {
        "cover": html.xpath('//div[@class="details-book-cover-content"]/a/@href')[0],
        "name": str(html.xpath('//h1[@itemprop="name"]/text()')[0]).strip(),
        "author": str(html.xpath('string(//div[@class="col-sm-9"]/i)')).strip(),
        "rate": float(html.xpath('//span[@class="book-rating-interest-score "]/text()')[0]),
        "description": re.sub(r"\s+", "", str(html.xpath('string(//div[@id="bookDescriptionBox"])'))),
        "category": str(html.xpath('//div[@class="bookDetailsBox"]/div[1]/div[2]/a/text()')[0]),
        # "ebook": "https://zh.b-ok.cc" + str(html.xpath('//div[@class="btn-group"]/a/@href')[0])
    }
    other = ["https://zh.b-ok.cc" + i for i in html.xpath('//div[@id="bMosaicBox"]/div/a/@href')]
    return info, other


def downloadEBook(name, url):
    with open("ebooks/" + name + ".epub", "wb") as f:
        f.write(get_content(url).content)


def storeInfo(info, conn: MySQLdb.connect):
    try:
        cursor = conn.cursor()
        cursor.execute(
            "INSERT IGNORE INTO categories (name) values (%s);",
            [info["category"]])

        cursor.execute(
            "INSERT IGNORE INTO books (name, author, description, ebook, cover,category_book) values (%s,%s,%s,%s,%s,%s)",
            [info["name"],
             info["author"],
             info["description"],
             info["ebook"],
             info["cover"],
             cursor.lastrowid])

        conn.commit()
        cursor.close()
    except Exception as e:
        print(e)
        conn.rollback()
    finally:
        print(f'《{info["name"]}》')


if __name__ == '__main__':
    conn = MySQLdb.connect(
        host='116.62.229.23',
        port=3306,
        user='root',
        passwd='104494lw!',
        db='bookstore',
    )

    urls = ["https://zh.b-ok.cc/book/3689456/e1b90b?dsource=mostpopular"]
    count = 0
    for url in urls:
        try:
            htm = get_content(url).text
            info, other = get_info(htm)
            info["ebook"] = url
            storeInfo(info, conn)
            urls.extend(other)
            count += 1
            if count > 300:
                break
        except Exception as e:
            print(e)

        time.sleep(5)
    conn.close()
