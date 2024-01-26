from kubemq.queue.message_queue import MessageQueue
from kubemq.queue.message import Message

if __name__ == "__main__":
    sender = MessageQueue("q1", "sender_client_id", "localhost:50000")
    message = Message()
    message.body = "hello kubemq - sending single message".encode('UTF-8')
    try:
        result = sender.send_queue_message(message)
        if result.error:
            print('message enqueue error, error:' + result.error)
        else:
            print('message sent to the queue')
    except Exception as err:
        print('message enqueue error, error:%s' % (
            err
        ))

    receiver = MessageQueue("q1", "receiver_client_id", "localhost:50000", 2, 1)
    try:
        res = receiver.receive_queue_messages()
        if res.error:
            print(
                "'Received:'%s'" % (
                    res.error
                )
            )
        else:
            for message in res.messages:
                print(
                    "'Received :%s ,Body: sending:'%s'" % (
                        message.MessageID,
                        message.Body
                    )
                )
    except Exception as err:
        print(
            "'error sending:'%s'" % (
                err
            )
        )