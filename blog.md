# Can we build url shortner without db?

hey, so few days ago, i got a ques in my mind, "Can we build url shortner without db?" yeah, no in
memory, just compute on the fly. 

but before that, lemme tell u what we gonna do thoughout this article
## know what is url shortner
## setup sample files to benchmark 
## try out different different approaches
## conclusion


## what is url shortner?
u can skip this section if u already know. 
it's what name say, something short the url. 
few famous examples are [bit.ly](https://bitly.com/), [tinyURL](https://tinyurl.com/) and there's lot more

### what they do??
they take ur long url, return another url. the return url is some url with hash, that u can use 
whereever u wish. whereever u use that url, it'll redirect to the ur original url. that's it. 

### basic idea, how most of them build?
when u submit the url, they generate a `key` for that url. store that in any sorta db, with the 
url as value. whenever u hit `/key` it'll redirect to `url` that already stored in db. and edge cases
like 404 etc, and if needed add more features, like analytics, rate limiting etc. it depends,
but basically, it is what it is.

### what im trying to do?
so i try to store the url itself in the key, and compute it on the demand. no db needed. 


## how we gonna benchmarking??
so, i got a repo, it contains massive amount of urls. 

a script, that'll take the urls one by one, short it, measure how much % it reduced, 
store the output.





--- 
repo use for benchmarking: https://github.com/ada-url/url-various-datasets
