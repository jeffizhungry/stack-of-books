## Microservices


## Integration Mechanisms
- *Message Brokers*
   - RabbitMQ
      * one master node
      * good for low-volume queued services
      * uses acks for reliable communication
   
   - Kafka
      * distributed
      * good for high-volume queued services
      * commit log for reliable communication, but requires consumer to keep track of offset

      - [Apache Documentation] (http://kafka.apache.org/documentation.html)
      
   - AWS Kinesis
      * very similar to kafka: distributed with commit log


- *JSON over HTTP*
   * good for RPC (request-response) services

- *Protocol Buffers*
   - [Overview] (https://developers.google.com/protocol-buffers/docs/overview)

   - [5 Reasons to Choose Protocol Buffers] (http://blog.codeclimate.com/blog/2014/06/05/choose-protocol-buffers/)

      * Schema Verification: optionals, types, versions

- *Database*
   * [Building Microservices] (https://www.nginx.com/wp-content/uploads/2015/01/Building_Microservices_Nginx.pdf) advocates not using databases as an integration mechanism because it restricts people to a database technology. 


## RESTful Services
- [Richardson Maturity Model] (http://martinfowler.com/articles/richardsonMaturityModel.html)

    * Abstraction levels of RESTful services


## Monitoring
- *Centralized Logging Solutions*
   - ELK Stack


## Component Testing
