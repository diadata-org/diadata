from selenium import webdriver
import time

delayTime = 60
waitTime = 900

driver = webdriver.Firefox()
driver.get('https://explore.duneanalytics.com/dashboard/dia-bonding-curve-dashboard_1')

# login
username = driver.find_elements_by_xpath('//*[@id="signInFormUsername"]')[1]
username.send_keys("username")

# password
password = driver.find_elements_by_xpath('//*[@id="signInFormPassword"]')[1]
password.send_keys("password")

# click submit button
submit_button = driver.find_elements_by_xpath('/html/body/div[1]/div/div[2]/div[2]/div[2]/div[2]/div/div/form/input[3]')[0]
submit_button.click()

# get required dashboard
time.sleep(1)
driver.get('https://explore.duneanalytics.com/dashboard/dia-bonding-curve-dashboard_1')
time.sleep(1)

#get all refresh buttons
widgets = driver.find_elements_by_xpath('/html/body/section/app-view/div/dashboard-page/div/div[2]/div/div')

refresh_button = []
for i in range(1, len(widgets) + 1):
    refresh_button = refresh_button + driver.find_elements_by_xpath('/html/body/section/app-view/div/dashboard-page/div/div[2]/div/div[' + str(i) + ']/div[1]/dashboard-widget/div/div/div[3]')


while True:
    for i in range(1, len(refresh_button) + 1):
        refresh_button[i].click()
        time.sleep(delayTime)
    time.sleep(waitTime)


