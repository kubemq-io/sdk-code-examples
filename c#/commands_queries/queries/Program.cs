using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using KubeMQ.Grpc;
using KubeMQ.SDK.csharp.CommandQuery;
using KubeMQ.SDK.csharp.QueueStream;

namespace queries
{
    class Program
    {
        static async Task Main(string[] args)
        {
            var  responder = new KubeMQ.SDK.csharp.CommandQuery.Responder("localhost:50000");
            try
            {
                responder.SubscribeToRequests(new KubeMQ.SDK.csharp.Subscription.SubscribeRequest
                    {
                        Channel = "q1",
                        SubscribeType = KubeMQ.SDK.csharp.Subscription.SubscribeType.Queries,
                        ClientID = "subscriber-1",
                    }, (request) =>
                    {
                        Console.WriteLine($"Query Request Received: Body:{ System.Text.Encoding.UTF8.GetString(request.Body)}, responding... ");
                        return new KubeMQ.SDK.csharp.CommandQuery.Response(request)
                        {
                            Body = "hello kubemq - sending a query response"u8.ToArray(),
                            Executed = true,
                            Timestamp = DateTime.UtcNow,
                        };
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

            var sender = new KubeMQ.SDK.csharp.CommandQuery.Sender("localhost:50000");
            try
            {
                var result =  await sender.SendRequestAsync(new QueryRequest()
                {
                    Channel = "q1",
                    Body = "hello kubemq - sending a query message"u8.ToArray(),
                    Timeout = 2000,
                });
                Console.WriteLine($"Query Response Received: Executed:{result.Executed} Body:{ System.Text.Encoding.UTF8.GetString(result.Body)} ");
                
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);          
            }
        }
    }
}