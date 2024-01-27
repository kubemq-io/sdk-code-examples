using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using KubeMQ.Grpc;
using KubeMQ.SDK.csharp.CommandQuery;
using KubeMQ.SDK.csharp.QueueStream;

namespace commands
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
                        Channel = "c1",
                        SubscribeType = KubeMQ.SDK.csharp.Subscription.SubscribeType.Commands,
                        ClientID = "subscriber-1",
                    }, (request) =>
                    {
                        Console.WriteLine($"Command Request Received: Body:{ System.Text.Encoding.UTF8.GetString(request.Body)} ");
                        return new KubeMQ.SDK.csharp.CommandQuery.Response(request)
                        {
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
                var result =  await sender.SendRequestAsync(new CommandRequest()
                {
                    Channel = "c1",
                    Body = "hello kubemq - sending a command message"u8.ToArray(),
                    Timeout = 1000,
                });
                Console.WriteLine($"Command Response Received: Executed:{result.Executed} Error:{result.Error} ");
                
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);          
            }
            await Task.Delay(1000);
        }
    }
}