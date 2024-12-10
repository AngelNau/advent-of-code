import requests
from bs4 import BeautifulSoup
from datetime import datetime as dt

day = dt.now().day
URL = f"https://adventofcode.com/2024/day/{day}/input"
page = requests.get(URL)

soup = BeautifulSoup(page.content, "html.parser")
print(soup.find(tag="pre"))
