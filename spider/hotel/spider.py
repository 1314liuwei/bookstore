import json
import time

import requests
from lxml import etree


def download_content(page):
    url = f"https://www.booking.com/searchresults.zh-cn.html?aid=304142;label=wec301zr-1FCAEoggI46AdIM1gEaIkCiAEBmAEruAEGyAEM2AEB6AEB-AELiAIBqAIDuAK-9rGQBsACAdICJDU4Zjc1MDZkLWU1OGEtNGFmMS04Mzc0LWVjOTBmNTIzOWNkYdgCBuACAQ;sid=3320376ce094ca7be6c358e54a909d68;tmpl=searchresults;checkin_month=3;checkin_monthday=10;checkin_year=2022;checkout_month=3;checkout_monthday=11;checkout_year=2022;city=900047908;class_interval=1;dest_id=900047908;dest_type=city;dtdisc=0;from_sf=1;group_adults=2;" \
          f"group_children=0;inac=0;index_postcard=0;label_click=undef;no_rooms=1;offset={page * 25};" \
          f"order=popularity;postcard=0;room1=A%2CA;sb_price_type=total;shw_aparth=1;" \
          f"slp_r_match=0;src_elem=sb;srpvid=427b602e9d3a00e2;ss=%25E7%25A6%258F%25E5%2586%2588;" \
          f"ss_all=0;ssb=empty;sshis=0;ssne=%25E7%25A6%258F%25E5%2586%2588;ssne_untouched=%25E7%25A6%258F%25E5%2586%2588&;" \
          f"changed_currency=1;selected_currency=USD;top_currency=1"
    cookie = "_pxhd=hzhYC2hKgbEL1On6WnozWqKO6CJC83uR9sI9PLDSWTBFz9GO3ALFCz58zGQABb74LRbMK9OlpMHhPbKYcUatIw%3D%3D%3AT8ribxT2eABcQcaKR53Y4Rc%2Fjhcbxv" \
             "Rp%2FYu%2F7OXwa4saeBMGAqWvldznLyoFzD7EIXuZYSQurLUpXPbw5%2FAdlx1FcwRnrBqS0zaKm0Ww2uk%3D; cnfunco=1; cors_js=1; BJS=-;" \
             " bkng_sso_session=e30; pxcts=03c0fca5-8f1d-11ec-b5e7-454d6c446355; _pxvid=fc367fb2-8f1c-11ec-8c70-6a6173414a45; has_preloaded=1" \
             "; OptanonConsent=isGpcEnabled=0&datestamp=Wed+Feb+16+2022+21%3A40%3A45+GMT%2B0800+(GMT%2B08%3A00)&version=6.22.0&isIABGlobal=fals" \
             "e&hosts=&consentId=36da6ad4-12ab-4a31-8ce5-bb777f962154&interactionCount=1&landingPath=https%3A%2F%2Fwww.booking.com%2Fsearchresults" \
             ".zh-cn.html%3Flabel%3Dwec301zr-1FCAEoggI46AdIM1gEaIkCiAEBmAEruAEGyAEM2AEB6AEB-AELiAIBqAIDuAK-9rGQBsACAdICJDU4Zjc1MDZkLWU1OGEtNGFmMS04M" \
             "zc0LWVjOTBmNTIzOWNkYdgCBuACAQ%26sid%3D02a291898390b7d4099d06de5dcfb89d%26aid%3D304142%26sb%3D1%26src_elem%3Dsb%26error_url%3Dhttps%25253A%2" \
             "5252F%25252Fwww.booking.com%25252Fsearchresults.zh-cn.html%25253Flabel%25253Dwec301zr-1FCAEoggI46AdIM1gEaIkCiAEBmAEruAEGyAEM2AEB6AEB-AELi" \
             "AIBqAIDuAK-9rGQBsACAdICJDU4Zjc1MDZkLWU1OGEtNGFmMS04Mzc0LWVjOTBmNTIzOWNkYdgCBuACAQ%25253Bsid%25253D02a291898390b7d4099d06de5dcfb89d%25253Btm" \
             "pl%25253Dsearchresults%25253Bac_click_type%25253Db%25253Bac_position%25253D0%25253Bclass_interval%25253D1%25253Bdest_id%25253D900047908%252" \
             "53Bdest_type%25253Dcity%25253Bdtdisc%25253D0%25253Bfrom_sf%25253D1%25253Bgroup_adults%25253D2%25253Bgroup_children%25253D0%25253Biata%25253D" \
             "FUK%25253Binac%25253D0%25253Bindex_postcard%25253D0%25253Blabel_click%25253Dundef%25253Bno_rooms%25253D1%25253Boffset%25253D0%25253Bpostcard%25" \
             "253D0%25253Braw_dest_type%25253Dcity%25253Broom1%25253DA%2525252CA%25253Bsb_price_type%25253Dtotal%25253Bsearch_selected%25253D1%25253Bshw_aparth" \
             "%25253D1%25253Bslp_r_match%25253D0%25253Bsrc%25253Dindex%25253Bsrc_elem%25253Dsb%25253Bsrpvid%25253D17621ead87d8008b%25253Bss%25253D%252525E7%25252" \
             "5A6%2525258F%252525E5%252525B2%252525A1%252525E5%252525B8%25252582%2525252C%25252520%252525E7%252525A6%2525258F%252525E5%252525B2%252525A1%252525E7%25" \
             "25259C%2525258C%2525252C%25252520%252525E6%25252597%252525A5%252525E6%2525259C%252525AC%25253Bss_all%25253D0%25253Bss_raw%25253D%252525E7%252525A6%2525" \
             "258F%252525E5%252525B2%252525A1%25253Bssb%25253Dempty%25253Bsshis%25253D0%252526%25253B%26ss%3D%2525E7%2525A6%25258F%2525E5%252586%252588%26is_ski" \
             "_area%3D0%26ssne%3D%2525E7%2525A6%25258F%2525E5%252586%252588%26ssne_untouched%3D%2525E7%2525A6%25258F%2525E5%252586%252588%26city%3D900047908" \
             "%26checkin_year%3D2022%26checkin_month%3D3%26checkin_monthday%3D10%26checkout_year%3D2022%26checkout_month%3D3%26checkout_monthday%3D11%26g" \
             "roup_adults%3D2%26group_children%3D0%26no_rooms%3D1%26from_sf%3D1%26order%3Dpopularity&implicitConsentCountry=nonGDPR&implicitConsentDate=1" \
             "645018849705&groups=C0001%3A1%2CC0002%3A0%2CC0004%3A0; _gcl_au=1.1.418312906.1645018851; bkng_prue=1; _ga=GA1.1.1298771749.1645018851; " \
             "g_state={\"i_p\":1645026081233,\"i_l\":1}; _ga_FPD6YLJCJ7=GS1.1.1645018850.1.1.1645018896.0; bkng_sso_ses=eyJib29raW5nX2dsb2JhbCI6W3siaCI6" \
             "Ik00NU0zclNqZllYa1U0eWtLYTRyc0xwMHptbkVpT0w5bG9GZTZObi9SbFEiLCJhIjoxfV19; bkng=11UmFuZG9tSVYkc2RlIyh9YfDNyGw8J7nzPnUG3Pr%2Bfv6X5X4%2BLV2wNe" \
             "udoQO8K2P0sW5GDjbolNq9%2BnJiW7s4dXhZ26RVJ88uh6x6ENBaN0ycfw2V1AmdnhV36w6SfwHYsmK%2FnlEsuo%2FNcGEdHWazmK9YX6EkfaLw7hfJu%2FfL0%2F22%2F1NqfVo9K" \
             "WDPXipxWFSMJk4em42S7MIkq1Pc%2Bg9ufQ%3D%3D; _px3=417244eecf9659bf7d7e95482e07a95f1d14d67d6c643e710162194c1bd69f97:o2X9nSq2zcNJXVNoKLfTTmRIzp" \
             "YMNodY07PNBvLgbW0DBk+QBUyasD3Eeh3c1LZCmxjswhSLacRutfDPK+WMWg==:1000:vLY96t91jjkFSuRBLH/eJXcYqrFsl1nlrU2CugGZkICGpYuEnt8uywALdgznrFzyUSnU/Ym" \
             "eUGLuhFSViWDaGz5jrjMfw0SDYsbAzWoA4QvXSCBQkCPHhAMOrLh72AiKy6ZNOICMKwPNnJhYQbD9oLnVlv9TayqqhLTIrcwb1IZ8byGG37iXM5d6OuyVqghK1zEuep5uP85pwXOITI" \
             "OxdQ==; _pxde=f5552726f5650170faf8cc599ffa54fe52b085bd06a1a08779995ffe6186f6ef" \
             ":eyJ0aW1lc3RhbXAiOjE2NDUwMTg5NjEzMDcsImZfa2IiOjAsImlwY19pZCI6" \
             "W119; lastSeen=0"
    referer = "https://www.booking.com/searchresults.zh-cn.html?label=wec301zr" \
              "-1FCAEoggI46AdIM1gEaIkCiAEBmAEruAEGyAEM2AEB6AEB-AELiAIBqAIDuAK" \
              "-9rGQBsACAdICJDU4Zjc1MDZkLWU1OGEtNGFmMS04Mzc0LWVjOTBmNTIzOWNkYdgCBuACAQ&sid" \
              "=02a291898390b7d4099d06de5dcfb89d&aid=304142&sb=1&src_elem=sb&error_url=https%253A%252F%252Fwww" \
              ".booking.com%252Fsearchresults.zh-cn.html%253Flabel%253Dwec301zr" \
              "-1FCAEoggI46AdIM1gEaIkCiAEBmAEruAEGyAEM2AEB6AEB-AELiAIBqAIDuAK" \
              "-9rGQBsACAdICJDU4Zjc1MDZkLWU1OGEtNGFmMS04Mzc0LWVjOTBmNTIzOWNkYdgCBuACAQ%253Bsid" \
              "%253D02a291898390b7d4099d06de5dcfb89d%253Btmpl%253Dsearchresults%253Bac_click_type%253Db" \
              "%253Bac_position%253D0%253Bclass_interval%253D1%253Bdest_id%253D900047908%253Bdest_type%253Dcity" \
              "%253Bdtdisc%253D0%253Bfrom_sf%253D1%253Bgroup_adults%253D2%253Bgroup_children%253D0%253Biata%253DFUK" \
              "%253Binac%253D0%253Bindex_postcard%253D0%253Blabel_click%253Dundef%253Bno_rooms%253D1%253Boffset%253D0" \
              "%253Bpostcard%253D0%253Braw_dest_type%253Dcity%253Broom1%253DA%25252CA%253Bsb_price_type%253Dtotal" \
              "%253Bsearch_selected%253D1%253Bshw_aparth%253D1%253Bslp_r_match%253D0%253Bsrc%253Dindex%253Bsrc_elem" \
              "%253Dsb%253Bsrpvid%253D17621ead87d8008b%253Bss%253D%2525E7%2525A6%25258F%2525E5%2525B2%2525A1%2525E5" \
              "%2525B8%252582%25252C%252520%2525E7%2525A6%25258F%2525E5%2525B2%2525A1%2525E7%25259C%25258C%25252C" \
              "%252520%2525E6%252597%2525A5%2525E6%25259C%2525AC%253Bss_all%253D0%253Bss_raw%253D%2525E7%2525A6" \
              "%25258F%2525E5%2525B2%2525A1%253Bssb%253Dempty%253Bsshis%253D0%2526%253B&ss=%25E7%25A6%258F%25E5%2586" \
              "%2588&is_ski_area=0&ssne=%25E7%25A6%258F%25E5%2586%2588&ssne_untouched=%25E7%25A6%258F%25E5%2586%2588" \
              "&city=900047908&checkin_year=2022&checkin_month=3&checkin_monthday=10&checkout_year=2022" \
              "&checkout_month=3&checkout_monthday=11&group_adults=2&group_children=0&no_rooms=1&from_sf=1&order" \
              "=popularity "
    headers = {
        "cookie": cookie,
        "referer": referer,
        "Host": "www.booking.com",
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36 Edg/98.0.1108.50"
    }
    r = requests.get(url, headers=headers)
    with open("result.html", "w", encoding="utf-8") as f:
        f.write(r.text)


def get_data(r):
    with open("result.html", "r", encoding="utf-8") as f:
        data = f.read()

    html = etree.HTML(data)
    names = html.xpath("//div[@data-testid='title']/text()")
    addresses = html.xpath("//span[@data-testid='address']/text()")
    stars = html.xpath("//div[@data-testid='rating-stars']")
    links = html.xpath("//a[@data-testid='title-link']/@href")
    prices = html.xpath("//div[@data-testid='price-and-discounted-price']/span[last()]/text()")
    for i, j in enumerate(zip(names, addresses, stars, links, prices)):
        tmp = {
            "name": j[0],
            "address": j[1],
            "star": j[2].xpath("count(.//span)"),
            "link": j[3],
            "price": float(j[4][3:].replace(",", "")),
        }
        r.append(tmp)


if __name__ == '__main__':
    result = []
    for i in range(3):
        download_content(i)
        get_data(result)
        time.sleep(1)

    for i in sorted(result, key=lambda x: x["price"]):
        print(i)
    # with open("result.json", "w", encoding="utf-8") as f:
    #     f.write(json.dumps({
    #         "data": result
    #     }))
