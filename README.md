# micro-cache

A simple, and light weight in-memory key-value store that can run independently, or be distributed across multiple systems. For, when all you need is a blob of data to be available on-demand, and available to multiple clients, theres no need for a large, cluncky database system. Micro-cache is made to be simple enough that it can even run effectivly on as small of computer as a raspberry pi. 

* Basic Key-value store.
* Made to be an extremely small, and simple. 
* Values are just stored as simple byte arrays.
    - This is to make it as flexable as possible.
    - **The client** is responsable for **security** using the data.

        
## TODO
- [ ] Persistance mode: Asyncronosly keep a copy of the in-memory data on disk, for recovery or logging.
- [ ] Create golang package for easier unmarshaling of data recieved from microcache.
- [ ] Add create a master node so it can be distributed accross different systems.
- [ ] Replace the storage data-structure with more performant alternative if possable.
- [ ] Security? Token?

