from kubemq.queue.message_queue import MessageQueue
from kubemq.queue.message import Message

if __name__ == "__main__":
    sender = MessageQueue("q1", "sender_client_id", "localhost:50000")
    message = Message()
    message.body = "hello kubemq - sending single message".encode('UTF-8')
    try:
        result = sender.send_queue_message(message)
        if result.error:
            print('message sending error, error:' + result.error)
        else:
            print('message sent to the queue')
    except Exception as err:
        print('message sending error, error:%s' % (
            err
        ))

    receiver = MessageQueue("q1", "receiver_client_id", "localhost:50000", 2, 1)
    try:
        tr = receiver.create_transaction()
        stream = tr.receive(1, 5)
        print(stream.message.Body.decode('UTF-8'))
        tr.ack_message(stream.message.Attributes.Sequence)
        tr.close_stream()
    except Exception as err:
        print(
            "'error receiving:'%s'" % (
                err
            )
        )

