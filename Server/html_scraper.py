import mechanize
from bs4 import BeautifulSoup
import urllib2 
import cookielib
import selenium as selenium
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
import time
from selenium.webdriver.common.action_chains import ActionChains
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import sys, getopt

def main(argv):



	food_name = argv[0].replace("_", " ")
	serving_size = argv[1]

	print food_name
	print serving_size

	#driver = webdriver.Firefox()
	driver = webdriver.PhantomJS()
	driver.get("https://www.myfitnesspal.com/account/logout#fancy_login")
	username = driver.find_element_by_id("username")
	password = driver.find_element_by_id("password")

	username.send_keys("xxxxx")
	password.send_keys("xxxxx")
	print "1"
	driver.find_element_by_xpath('//form').submit()

	driver.get("http://www.myfitnesspal.com/food/search")
	print "2"
	search_food = driver.find_element_by_id("search")
	search_food.send_keys("Creamy Peanut Butter")
	driver.find_element_by_xpath('//form').submit()
	#driver.find_element_by_xpath('//li').click()
	div = driver.find_element_by_id("main")
	#first_element = div.find_element_by_xpath("//li[11]").click()

	html_list = driver.find_element_by_id("main")
	items = html_list.find_elements_by_tag_name("li")
	for item in items:
	    item.click()
	    print item.text
	    break

	time.sleep(3)
	food_quantity = driver.find_elements_by_xpath('//input[@class="text short"]')[1]
	food_quantity.clear()
	food_quantity.send_keys(serving_size)
	submit_button = driver.find_elements_by_xpath('//input[@class="button log"]')[1]
	submit_button.click()
	time.sleep(5)
	print "Done"

if __name__ == "__main__":
   main(sys.argv[1:])





