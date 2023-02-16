## Cache 
Cache makes our application performant by storing data in memory which otherwise would have taken 100-1000x time more time if queried from DB.

Cache can be used at places like 
- in web application layer to cache session objects.
- In Service layer to cache created objects if these objects are accessed frequently and their data is not modified very frequently.

### Redis

Redis is type of database. All of data is stored inside key value pair. 
Unlike normal Database that stores information on disk, Redis run inside RAM, So fast but volatile and is generally used for caching data.

#### Architecture

<img width="500" alt="image" src="https://user-images.githubusercontent.com/32810320/219389551-ef11cce4-a0fb-42a3-b2b1-67f534f09ff4.png">
Reference: https://nutanix.udemy.com/course/developer-to-architect/learn/lecture/24984768#questions from newtechways.com

It is [key -> Data Structure] store.
Where Data Structure includes, Strings, Lists, SortedSets, Maps

Now, we can directly append an entry to the lists instead of fetching whole blob and writing whole data back in Memcache.

We can use redis as data store too, because in redis we can store data on disk and it allows to take backup and
cluster can restarted and pre-populated with a backup.

For **scalability**, Keys are distributed to different master node, and each master node does **async replication** to 
some slave nodes. These slave nodes help in Read load distribution, as for read we can go to master or any of its slave.

And In case master node goes down then one of its slave can be elected as master and that provides us **high availability**.

**cons**:
data store mode requres fixed number of nodes but cache mode is suitable for node scaling,

### Redis Pub/Sub 

Redis is well known for in-memory data store. If you need messaging to send to subscriber without any persistence, then redis is a good choice.

#### Architecture
<img width="483" alt="image" src="https://user-images.githubusercontent.com/32810320/219392385-62582276-9fa9-411e-99b9-05920f99b05e.png">
Reference: https://nutanix.udemy.com/course/developer-to-architect/learn/lecture/24984794#questions from newtechways.com

Lets say in above image on node 1, subscriber 1 and 2 are connected. On node 2, subscriber 3 and 4 are connected.
And these are connected with long lived TCP connection. 
publisher has subscribed to node 2. and message delivered to node 2 will be delivered to subcriber 3 and 4 immediately and
node 2 also relay the message to node 1 and which is then sent to subscriber 1 and 2.
 
if subcriber goes down at that time then subscriber will not get message. Redis won't retry to deliver the message.

Useful for making leaderboard dashboard.

Pros: 
Million operations per second

Cons:
Delivery not guaranted.

Comparison  
Kafka  
No Push
Writes to log

RabbitMQ Transient  
Delivery acknowledgement  
Deletion of delivered messages.  

can also be used as messaging queue.

### Setup

`brew install redis`

### Commands
`redis-server`
-> starts redis server on port 6379

run redis-cli in different terminal
`redis-cli`
```
SET name khushboo
GET name
DEL name

Check if key exists
EXISTS name

KEYS \*

flushall -> to clear entire cache

ttl name -> time to live for that key
-1 -> never expires
-2 -> Expired.
Otherwise remaining sec to live

expire name 10

setex name 10 abc

lpush friends swarnam -> to add an item to start of list

redis-c
li
127.0.0.1:6379> SET name Khushboo
OK
127.0.0.1:6379> GET name
"Khushboo"
127.0.0.1:6379> GET namE
(nil)
127.0.0.1:6379> EXISTS name
(integer) 1
127.0.0.1:6379> KEYS \*
"name"
127.0.0.1:6379> ttl name
(integer) -1
127.0.0.1:6379> SET name BCD
OK
127.0.0.1:6379> GET name
"BCD"
127.0.0.1:6379> flushall
OK
127.0.0.1:6379> GET name
(nil)
127.0.0.1:6379> SET name BCD
OK
127.0.0.1:6379> ttl name
(integer) -1
127.0.0.1:6379> expire name 10
(integer) 1
127.0.0.1:6379> ttl name
(integer) 5
127.0.0.1:6379> ttl name
(integer) 3
127.0.0.1:6379> GET name
(nil)
127.0.0.1:6379> ttl name
(integer) -2
127.0.0.1:6379> KEYS \*
(empty array)
127.0.0.1:6379> ttl name
(integer) -2
127.0.0.1:6379> setex name 10 abc
OK
127.0.0.1:6379> ttl name
(integer) 7
127.0.0.1:6379> lpush friends mukul swarnam surabhi smriti shivangi
(integer) 5
127.0.0.1:6379> lrange friends 0 -1
1. "shivangi"
2. "smriti"
3. "surabhi"
4. "swarnam"
5. "mukul"

127.0.0.1:6379> DEL friends
(integer) 1

rpush to push on right
LPOP

Redis can be used as message queue.
lpush to push on left, RPOP to pop from right

SET -> prefix with S

127.0.0.1:6379> SADD hobbies "reading books"
(integer) 1
127.0.0.1:6379> SMEMBERS hobbies
"reading books"
127.0.0.1:6379> SADD hobbies "reading books"
(integer) 0
127.0.0.1:6379> SMEMBERS hobbies
"reading books"
127.0.0.1:6379> SREM hobbies "reading books"
(integer) 1
127.0.0.1:6379> SMEMBERS hobbies
(empty array)

Hash -> prefix with H
add name property on person
127.0.0.1:6379> HSET person name khushboo
(integer) 1
127.0.0.1:6379> HSET person college BITS
(integer) 1
127.0.0.1:6379> HGET person name
"khushboo"
127.0.0.1:6379> HGET person
(error) ERR wrong number of arguments for 'hget' command
127.0.0.1:6379> HGETALL person

1. "name"
2. "khushboo"
3. "college"
4. "BITS"

delete name property on person
127.0.0.1:6379> HDEL person college
(integer) 1
127.0.0.1:6379> HGETALL person

1. "name"
2. "khushboo"

redis-cli shutdown -> shutdown redis server
```
My Example Code: https://github.com/KHUSHBOO0012/TechStack/blob/main/Src/Redis/main.go

Reference: 
https://www.youtube.com/watch?v=jgpVdJB2sKQ&ab_channel=WebDevSimplified  
https://nutanix.udemy.com/course/developer-to-architect/learn/lecture/24984768#questions
