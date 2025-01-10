# Berlin Station Board

--------------------

### Backend High-Level Design

### Using singleflight to Prevent Duplicate Calls
The singleflight package ensures that if multiple goroutines make the same request (in this case, refreshing the cache for departure times), only one request will actually go through to the API. The others will wait for the first request to complete and receive the same response. This is particularly useful during high traffic periods or when the cache for a popular station expires and needs refreshing.

### Backend Caching Mechanism:
The backend will maintain a cache of departure times.
Cache entries will have a TTL (Time To Live) calculated as departure time - time.Now().
When a cache entry expires (is evicted due to TTL), the next request for that station will trigger an API call to refresh the cache.

### Adaptive TTL Adjustments
Adaptively adjust TTLs based on the time of day or specific departure characteristics. For instance, during peak hours when departures are more frequent and potentially more volatile, use shorter TTLs. For late-night hours with less frequent changes, longer TTLs could be more appropriate.
1. Define Criteria for TTL Variation
   Start by identifying the criteria that will influence the TTL settings. For a station board application, consider:
    - Time of Day: Peak hours might need shorter TTL due to higher frequency of changes and user queries.
        - 7AM to 9AM - 5PM to 7PM
        - Weekends
        - holidays
    - Departure Time Proximity: Departures closer in time might need fresher data than those further away.


### User Experience During Data Refresh: // TODO
Users might see slightly outdated information if they request data right before a refresh.
Solution: Communicate clearly to users that data is being updated or show a countdown to the next refresh.

--------------------
## FrontEnd
### User Notifications: 
If the departure information changes (e.g., a delay or cancellation), consider how these updates are communicated to users, especially those who may have already checked the departure time.

### Adaptive Polling: 
Dynamically adjust the polling frequency based on user activity or time to departure. For example, increase the polling interval when a departure is further away and decrease it as the departure time approaches.

### Use Efficient HTTP Requests: 
Ensure that your polling requests are as lightweight as possible. For example, you could use HTTP GET requests that return only the necessary data. Consider using HTTP headers like If-Modified-Since to avoid fetching data that hasn't changed.


--------------------

## Monitoring
### Monitor and Adjust
Implement monitoring to track the effectiveness of your adaptive TTL adjustments. Collect metrics on cache hit/miss rates, data freshness, and user satisfaction. Use this data to fine-tune your TTL calculation criteria and adjustments.
