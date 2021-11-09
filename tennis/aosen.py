from posix import PRIO_PGRP
from binascii import b2a_hex, a2b_hex
from Crypto import Cipher
from Crypto.Cipher import AES
import cv2
import os, base64
import binascii 
import requests
import json
import time
import urllib.parse

from timer import Timer
from requests.models import Response

# 提交订单
def addParkOrder(string):
    url = "http://tennis.bjofp.cn/TennisCenterInterface/pmPark/addParkOrder.action"

    payload={}
    payload['userid']='31459'
    
    payload['parkList']='[{"date":"2021-11-12","time":"19","parkid":"23","parkname":"A10"},{"date":"2021-11-12","time":"20","parkid":"23","parkname":"A10"}]'
    payload['paywaycode']='2'
    payload['cardnumber']=''
    payload['mobile']=''
    payload['ordercode']=''
    payload['captchaVerification']=string
    payload=urllib.parse.urlencode(payload)

    headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63040026)',
    'Content-Type': 'application/x-www-form-urlencoded',
    'Cookie': 'JSESSIONID=DB31273AAD0C3E24C80ACDF7221B9EB8; openid=oR_qexDMsNulkvM0F8E9GSY0ISkA'
    }
    response = requests.request("POST", url, headers=headers, data=payload)

    print("============================")
    print(response.text)

# 获取信息：secretKey、back和tar的base64、
def getPictureInfo():
    
    url = "http://tennis.bjofp.cn/TennisCenterInterface/imgCaptcha/api/get.action"

    payload="{\r\n    \"captchaType\":\"blockPuzzle\",\r\n    \"clientUid\":\"slider-4c6adace-84a1-4a1b-b061-ad5b028022ff\",\r\n    \"ts\":"+str(int(time.time()*1000))+",\r\n    \"userid\":\"31459\"\r\n}"
    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x6304001f)',
        'Content-Type': 'Application/json;charset=UTF-8',
        'Cookie': 'JSESSIONID=DB31273AAD0C3E24C80ACDF7221B9EB8; openid=oR_qexDMsNulkvM0F8E9GSY0ISkA'
    }

    response = requests.request("POST", url, headers=headers, data=payload)
    object = json.loads(response.text)
    return object

# check info 
def checkInfo(object):
    url = "http://tennis.bjofp.cn/TennisCenterInterface/imgCaptcha/api/check.action"

    payload="{\r\n    \"captchaType\":\"blockPuzzle\",\r\n    \"clientUid\":\"slider-4c6adace-84a1-4a1b-b061-ad5b028022ff\",\r\n    \"pointJson\":\""+object['pointJson']+"\",\r\n    \"token\":\""+object['token']+"\",\r\n    \"ts\":"+str(int(time.time()*1000))+",\r\n    \"userid\":\"31459\"\r\n}"
    headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63030073)',
    'Content-Type': 'Application/json;charset=UTF-8',
    'Cookie': 'JSESSIONID=DB31273AAD0C3E24C80ACDF7221B9EB8; openid=oR_qexDMsNulkvM0F8E9GSY0ISkA'
    }
    
    response = requests.request("POST", url, headers=headers, data=payload)
    object = json.loads(response.text)
    return object
    

# 获取 back 和 tar 图片
def base64_to_back(string):
    img_data = base64.b64decode(string)
    with open('back.png', 'wb') as f:
        f.write(img_data)

def base64_to_tar(string):
    img_data = base64.b64decode(string)
    with open('tar.png', 'wb') as f:
        f.write(img_data)

# 计算图片之间滑动距离：横坐标x的距离
def _tran_canny(image):
    """消除噪声"""
    image = cv2.GaussianBlur(image, (3, 3), 0)
    return cv2.Canny(image, 50, 150)


def detect_displacement(img_slider_path, image_background_path):
    """detect displacement"""
    # # 参数0是灰度模式
    image = cv2.imread(img_slider_path, 0)
    template = cv2.imread(image_background_path, 0)

    # 寻找最佳匹配
    res = cv2.matchTemplate(_tran_canny(image), _tran_canny(template), cv2.TM_CCOEFF_NORMED)
    # 最小值，最大值，并得到最小值, 最大值的索引
    min_val, max_val, min_loc, max_loc = cv2.minMaxLoc(res)

    top_left = max_loc[0]  # 横坐标
    # 展示圈出来的区域
    x, y = max_loc  # 获取x,y位置坐标

    w, h = image.shape[::-1]  # 宽高
    cv2.rectangle(template, (x, y), (x + w, y + h), (7, 249, 151), 2)
    # show(template)
    return top_left

# ECB加密
def AesDecryptECB(secretKey,text):
    bs = len(secretKey)
    endode_text=text.encode()
    length=len(endode_text)
    PADDING = lambda s: s + (bs - length % bs) * chr(bs - length % bs)
    entext = PADDING(text)
    aes = AES.new(str.encode(secretKey), AES.MODE_ECB)
    data=aes.encrypt(str.encode(entext))
    # aes_text=binascii.b2a_hex(aes.encrypt(str.encode(entext))) //hex
    encrypt_data = str(base64.b64encode(data),encoding='utf-8')
    return encrypt_data

if __name__ == '__main__':
    Timer().start()
    condition = True
    while condition:
        response = getPictureInfo()
        secretKey = response['repData']['secretKey']
        token = response['repData']['token']
        back = response['repData']['originalImageBase64']
        base64_to_back(back)
        tar = response['repData']['jigsawImageBase64']
        base64_to_tar(tar)
        top_left = str(detect_displacement("./tar.png", "./back.png"))
        position = '{"x":'+top_left+',"y":5}'
        pointJson = AesDecryptECB(secretKey,position)
        object={}
        object['token'] = token
        object['pointJson'] = pointJson
        if checkInfo(object)['repCode'] == '0000':
            break

    captchaVerification = AesDecryptECB(secretKey,token+"---"+position)
    addParkOrder(captchaVerification)





