# shorturl-go


## API:

init -> postShortURL

postShortURL:
 -> if already exist -> return existed URL (url exist state)
 -> not exist -> create a new one, and return it as response body (url created state)

getURL:
 -> if exist -> redirect to original URL (url exist state)
 -> not exist -> return 404 NotFound

```
GET /v1/url/<shortURL>
```


```
POST /v1/short
```

