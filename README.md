ht-golang
==================

Sample project that use golang for Heroku training.

Notes
==================

This project use custom buildpack.  
see [heroku-buildpack-go](https://github.com/kr/heroku-buildpack-go)

How to deploy
--------------

### Step1  

Clone app from Github  

```
$ git clone git@github.com:flect/ht-golang.git
$ cd ht-golang/
```

### Step2  

Create Heroku application

```
$ heroku create -b https://github.com/kr/heroku-buildpack-go
```

### Step3  

Deploy application
```
$ git push heroku master
```

### Step 4

Visit application

```
$ heroku open
```
