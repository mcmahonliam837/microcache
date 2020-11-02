# microcache

* Basic Key-value store.
* Made to be an extremely small, and simple. 
* Values are just stored as simple byte arrays.
    - This is to make it as flexable as possible.
    - **The client** is responsable for **security** using the data.

    
## Current State (Nov. 2 2020)

* Only in-memory volitile storage supported
* Basic hashmap to store the data.
    - Looking into alternatives

    
## TODO
- [ ] Persistance mode: Asyncronosly keep a copy if the in-memory data on disk.
    - This should not reduce speed by a noticable amount, as it will be done in a non-blocking go-routine pulling update requests from a queue. 
- [ ] Disk mode: Do not store data in-memory, save data with blocking-writes to a file.
- [ ] Replace the storage data-structure with more performant alternative if possable.
- [ ] Benchmark.
- [ ] Create golang package for easier unmarshaling of data recieved from microcache.
- [ ] Security? Token?

