# 肉

`arduino` で実装された低温調理器にコマンドを送る

## 動かし方

`./niku` を実行する

## 設定ファイル

ホームディレクトリに `niku.toml` というファイルを置いておく必要がある。ファイルの中身は次のようにする

```toml
[twilio]
sid = "hogehoge" # twilio の sdi。基本は変更しない
token = "hogehoge" # twilio の token。基本は変更しない
phone_number = "111" # twilio の発信元番号。基本は変更しない。
callback_url = "http://demo.twilio.com/docs/voice.xml" # twilio が発信する内容。変更しない。

[device]
device = "/dev/cu.usbmodem14341" # arduino のキャラクターデバイスのパス。
baudrate = 38400 # arudino とシリアル通信をするときのボードレート。

[cooker]
temperture = 55.5 # 料理をする温度。小数点を含む数字で指定する。
duration = 20 # 料理をする時間を秒で指定する。
```