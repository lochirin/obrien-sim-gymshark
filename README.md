This is the code for the simple Golang API that I wrote for the Gymshark interview.

A sample request body is as follows: 

```
{
 "sizes": [
    250,500,1000,2000,5000
 ],
 "capacity" : 12001
}
```

Package sizes can be added and modified simply by modifying the request body in order to fulfill the optional requirement mentioned in the specification document.

A sample response to the sample request above is as follows:

```
{
    "1000": 0,
    "2000": 1,
    "250": 1,
    "500": 0,
    "5000": 2
}
```

The API is currently deployed using GCP Cloud Run and can be hit through the https://gymshark-dtvatad2ya-nw.a.run.app/api/packages endpoint using a client such as Postman. 

![image](https://github.com/lochirin/obrien-sim-gymshark/assets/162650499/20ac76ba-34ab-44d7-832a-f3268e966f5b)
