import datetime
from time import sleep
from kubemq.commandquery import Responder, Sender, CommandRequest
from kubemq.commandquery.response import Response
from kubemq.subscription import SubscribeType, SubscribeRequest
from kubemq.tools import ListenerCancellationToken

if __name__ == "__main__":
    cancel_token = ListenerCancellationToken()
    try:
        responder = Responder("localhost:50000")
        subscribe_request = SubscribeRequest(
            channel="c1",
            subscribe_type=SubscribeType.Commands
        )


        def error_handler(error_msg):
            print("received error:%s'" % (
                error_msg
            ))


        def handler(request):
            if request:
                print("Subscriber Received request: Body:'%s'" % (
                    request.body,
                ))
                response = Response(request)
                response.executed = True
                response.timestamp = datetime.datetime.now()
                return response


        responder.subscribe_to_requests(subscribe_request, handler, error_handler, cancel_token)
        responder_is_connected = True
    except Exception as err:
        print('error, error:%s' % (
            err
        ))
        exit(1)
    sleep(1)
    sender = Sender("localhost:50000")
    request = CommandRequest(
        channel="c1",
        timeout=1000,
        body="hello kubemq - sending a command, please reply".encode('UTF-8'),
    )
    try:
        resp = sender.send_request(request)
        print("response Received: IsExecuted: %s, ExecutionTime: %s'" % (
            resp.executed,
            resp.timestamp,
        ))
    except Exception as err:
        print('error, error:%s' % (
            err
        ))
    cancel_token.cancel()
