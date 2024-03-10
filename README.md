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

This is pseudocode for the algorithm used to compute the number of packages of each size: 

1. First check if the desired capacity is less than the small package size, if it is, then it is just one package of the smallest size
2. If it is not, try to fit as many full packages as possible, starting from the biggest package size.
3. Once all full package sizes have been allocated, check if there is any remainder to deal with
4. If there is a remainder, first compute the excess item waste that would be accrued if we add on one package of the smallest size. The remainder will always be less than the smallest package size. 
5. Second, for each fully allocated package size computed in Step 2, compute what the item waste would be from making just that package size one size bigger (e.g. 250 -> 500 or 500 -> 1000 in the provided example) if the package size can be made bigger (the largest package size cannot be made any larger for instance)
6. Compute the item waste from making each package size bigger and get the minimum value of the range and compare it with the value computed in step 4.
7. Determine which approach accrues less item waste and go with that approach
