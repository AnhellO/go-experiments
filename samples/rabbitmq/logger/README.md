# RabbitMQ + Docker + Golang

Another sample taken from <https://github.com/Pungyeon/go-rabbitmq-example>

- [ ] I suggest to try to implement a simple API service, that uses RabbitMQ for event auditing. Sending all events of the API to the messaging broker and saving them as auditing logs
- [ ] This can then be extended to Event Sourcing, by using this log to regenerate state in your application, by going through the auditing logs and then based on those logs, recreating the data in your applications. This is somewhat complicated and there are a whole lot of considerations to be made.... but it's also really fun to experiment with :)
