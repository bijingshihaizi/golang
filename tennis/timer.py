# -*- coding:utf-8 -*-
import time
import requests
import json

from datetime import datetime


class Timer(object):
    def __init__(self, sleep_interval=0.5):
        buy_time = '2021-11-10 00:00:00.000'
        # localtime = time.localtime(time.time())
        self.buy_time = datetime.strptime(buy_time,"%Y-%m-%d %H:%M:%S.%f")
        self.buy_time_ms = int(time.mktime(self.buy_time.timetuple()) * 1000.0 + self.buy_time.microsecond / 1000)
        self.sleep_interval = sleep_interval


    def jd_time(self):
        """
        从京东服务器获取时间毫秒
        :return:
        """
        url = 'https://a.jd.com//ajax/queryServerData.html'
        ret = requests.get(url).text
        js = json.loads(ret)
        return int(js["serverTime"])

    def local_time(self):
        """
        获取本地毫秒时间
        :return:
        """
        return int(round(time.time() * 1000))

    def local_jd_time_diff(self):
        """
        计算本地与京东服务器时间差
        :return:
        """
        return self.local_time() - self.jd_time()

    def start(self):
        while True:
            # 本地时间减去与京东的时间差，能够将时间误差提升到0.1秒附近
            # 具体精度依赖获取京东服务器时间的网络时间损耗
            if self.local_time() >= self.buy_time_ms:
                print("start")
                break
            else:
                time.sleep(self.sleep_interval)
