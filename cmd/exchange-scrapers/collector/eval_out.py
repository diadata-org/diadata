#!/usr/bin/env python3

from datetime import datetime
import numpy as np
import requests


def getCoins():
	r =requests.get("https://api.coingecko.com/api/v3/coins/list")
	return eval(r.content)

def findToken(symbol):
	content = getCoins()
	for coin in content:
		if coin["symbol"] == symbol.lower():
			return coin
	return

def getPrice(coin):
	coinSymbol = coin["symbol"]
	r = requests.get("https://api.coingecko.com/api/v3/simple/price?ids=" + coinSymbol + "&vs_currencies=USD")
	if len(r.content) > 5:
		return eval(r.content)[coinSymbol]["usd"]
	coinID = coin["id"]
	r = requests.get("https://api.coingecko.com/api/v3/simple/price?ids=" + coinID+ "&vs_currencies=USD")
	return eval(r.content)[coinID]["usd"]
	
def getPair(line):
	# Get trading pair from log entry of exchange scraper
	line = line.split("&")[1].split()[1]
	return line.split("-")

def priceFromLog(line):
	# get price from log entry of exchange scraper
	return line.split("&")[1].split()[2]

def getPricePair(pair):
	coin1 = findToken(pair[0])
	price1 = getPrice(coin1)

	coin2 = findToken(pair[1])
	price2 = getPrice(coin2)

	try:
		q = price1/price2
		return q
	except ZeroDivisionError:
		return 0 


if __name__ == '__main__':


	lines = ""
	prices = []
	with open('out.txt', 'r') as f:
	    lines = f.readlines()

	# Go through trades and print if price on UniSwap differs
	# from price calculation through coingecko by more than 5%.
	for line in lines:
	    if "Got Trade " in line:
	        pair = getPair(line)
	   		
	        logPrice =  priceFromLog(line)
	        coingeckoPrice = getPricePair(pair)
	        if abs(logPrice - coingeckoPrice)/logPrice > 0.05:
	        	print("trade: ", line)
	        	print("price from coingecko: ", coingeckoPrice)
	        prices.append(line.strip())




