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

    titles, urls, imgs = [], [], []
    for i in range(1, 15):
        titles.extend(html.xpath(f'//li[@class="line{i} "]/a/@title'))
        urls.extend(html.xpath(f'//li[@class="line{i} "]/a/@href'))
        imgs.extend(html.xpath(f'//li[@class="line{i} "]/a/img/@src'))

    return list(zip(titles, urls, imgs))


if __name__ == '__main__':
    url = "https://book.dangdang.com/"
    result = get_dd_content(url)
    # print(result.text)
    for i in get_book(result.text):
        try:
            if i[0]:
                print(i)
        except IndexError as e:
            print(i)
