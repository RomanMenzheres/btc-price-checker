# Super Duper BTC Price Checker
#
# IT'S MY FIRST PROGRAM ON GOLANG!
#
# In project I am using my fake email, so you can use it too
#
# Logic of /rate method
## 1. Getting btc price from third party service as json and check error.
## 2. a) Bad request response if error.
## 2. b) Getting price from json.
## 3. Doing response with price.
#
# Logic of /subscribe method
## 1. Getting new Email by json.
## 2. Opening file or creating if not found and checking error (db.txt where contains our all emails).
## 3. Getting email addresses from our file as big string with spaces.
## 4. a) If our big string does not have new email we write it in file, do response and return.
## 4. b) If our big string already have new email we do appropriate response.
#
# Logic of /sendEmails method
## 1. Getting reponse body with btc price (we call http://localhost:80/rate to get json).
## 2. Getting btc price from reponse body.
## 3. Getting emails from db file.
## 4. Sending mail to every email in our db file using 'for' (I tried do that without 'for' but unsuccessful).
## 5. Creating a message and seting sender, receiver, header and body of mail.
## 6. Sending mail and checking error.
## 7. a) Panic if error.
## 7. b) Doing response if all is ok.
#
# About Docker
## I tried to make the file lighter but every time I ran it I got permission denied :(
