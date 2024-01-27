using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using KubeMQ.Grpc;
using KubeMQ.SDK.csharp.QueueStream;

namespace events_store
{
    class Program
    {
        static async Task Main(string[] args)
        {
            var  subscriber = new KubeMQ.SDK.csharp.Events.Subscriber("localhost:50000");
            try
            {
                subscriber.SubscribeToEvents(new KubeMQ.SDK.csharp.Subscription.SubscribeRequest
                    {
                        Channel = "es1",
                        SubscribeType = KubeMQ.SDK.csharp.Subscription.SubscribeType.EventsStore,
                        EventsStoreType = KubeMQ.SDK.csharp.Subscription.EventsStoreType.StartNewOnly,
                        ClientID = "subscriber-1",
                    }, (eventReceive) =>
                    {
                        Console.WriteLine($"Event Store Received: EventID:{eventReceive.EventID} Body:{ System.Text.Encoding.UTF8.GetString(eventReceive.Body)} ");
                    },
                    (errorHandler) =>                 
                    {
                        Console.WriteLine(errorHandler.Message);
                    });
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
            }
            await Task.Delay(1000);

            var sender = new KubeMQ.SDK.csharp.Events.Sender("localhost:50000");
            try
            {
                var result = sender.SendEvent(new KubeMQ.SDK.csharp.Events.Event()
                {                  
                    Channel = "es1",
                    Body = "hello kubemq - sending an event store message"u8.ToArray(),
                    Store = true
                });
                
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);          
            }
            await Task.Delay(1000);
        }
    }
}