import requests



s = requests.Session()

headers = {
    'Host': 'www.google.com',
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36',
    'Accept-Language': 'en-US,en;q=0.9',
    #'cookie': 'AEC=AVcja2e6Z9bNIiMs5nNJbT6YIitRPZO5V_0eFqQ97bTmuvU3qUh6sxMblRQ; DV=Yww-JzI09RIRIDlq9acmfs7SXiEgURk; NID=521=uRHa0k6WBSWXPR1QZ-CxwHYRIcgZuKUIZ8IROxZk-ThvUcAgL9ztMj8AoIcVaYpxirOzdBdXIE4oLjDNY07m1BSZ2a1PSswMVJaZllR6s1msEKttFAuHjh2IuoungMogxmPa3b5iGMCCOrrCyVRmYGfB8v6GdmwR5NAvDIBbf1orf8aQyMKiQImgP0qxEBHbZRPV2CAWkGoSqwLzyqYcUxYkdzPzM1irEa6HG5Ob3EoXdRaCa2x6J7_ZiaeS; GOOGLE_ABUSE_EXEMPTION=ID=657daa78ee9ae13e:TM=1739763900:C=r:IP=131.125.11.1-:S=Xt4PDBQvJMso_TsE68jwmd8'
}

params = {
    'q': 'test',
    'start': '0'
}

s.get('https://www.google.com/search', params=params, headers=headers, verify=False)
response = s.get('https://www.google.com/search', params=params, headers=headers, verify=False)
print(response.text)
# data = response.text.split('"WEB_RESULT_INNER"')[1:]
# print(len(data))
# for v in data:
#     print(v + "\n")