import datetime
from time import sleep
from kubemq.commandquery import Responder, Sender, QueryRequest
from kubemq.commandquery.response import Response
from kubemq.subscription import SubscribeType, SubscribeRequest
from kubemq.tools import ListenerCancellationToken

if __name__ == "__main__":
    cancel_token = ListenerCancellationToken()
    try:
        responder = Responder("localhost:50000")
        subscribe_request = SubscribeRequest(
            channel="q1",
            subscribe_type=SubscribeType.Queries
        )


        def error_handler(error_msg):
            print("received error:%s'" % (
                error_msg
            ))


        def handler(request):
            if request:
                print("Subscriber Received request: Metadata:'%s', Channel:'%s', Body:'%s' tags:%s" % (
                    request.metadata,
                    request.channel,
                    request.body,
                    request.tags
                ))
                response = Response(request)
                response.body = "OK - got your message".encode('UTF-8')
                response.cache_hit = False
                response.error = "None"
                response.executed = True
                response.metadata = "OK"
                response.timestamp = datetime.datetime.now()
                return response
        responder.subscribe_to_requests(subscribe_request, handler, error_handler, cancel_token)
    except Exception as err:
        print('error:%s' % (
            err
        ))
        exit(1)
        # give some time to connect a receiver
    sleep(1)
    sender = Sender("localhost:50000")
    request = QueryRequest(
        channel="q1",
        timeout=1000,
        body="hello kubemq - sending a query, please reply".encode('UTF-8'),
    )
    try:
        resp = sender.send_request(request)
        print("response Received: IsExecuted: %s, ExecutionTime: %s', Body: %s" % (
            resp.executed,
            resp.timestamp,
            resp.body
        ))
    except Exception as err:
        print('error, error:%s' % (
            err
        ))

    cancel_token.cancel()
