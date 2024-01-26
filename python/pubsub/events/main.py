from time import sleep
from kubemq.events import Sender, Event
from kubemq.tools.listener_cancellation_token import ListenerCancellationToken
from kubemq.events.subscriber import Subscriber
from kubemq.subscription.subscribe_type import SubscribeType
from kubemq.subscription.subscribe_request import SubscribeRequest

if __name__ == "__main__":
    cancel_token = ListenerCancellationToken()
    try:
        subscriber = Subscriber("localhost:50000")
        subscribe_request = SubscribeRequest(
            channel="e1",
            subscribe_type=SubscribeType.Events)


        def error_handler(error_msg):
            print("received error:%s'" % (
                error_msg
            ))

        def handler(event):
            if event:
                print("Subscriber Received Event: Channel:'%s', Body:'%s\n" % (
                    event.channel,
                    event.body,
                ))
        subscriber.subscribe_to_events(subscribe_request, handler, error_handler,
                                       cancel_token)
    except Exception as err:
        print('error:%s' % (
            err
        ))
        exit(1)
    sleep(1)
    sender = Sender("localhost:50000")
    event = Event(
        body=("hello kubemq - sending event ".encode('UTF-8')),
        store=False,
        channel="e1",
    )
    try:
        sender.send_event(event)
    except Exception as err:
        print('error:%s' % (
            err
        ))
    sleep(1)
    cancel_token.cancel()
