using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using KubeMQ.Grpc;
using KubeMQ.SDK.csharp.QueueStream;

namespace basic
{
    class Program
    {
        static async Task Main(string[] args)
        {
            QueueStream client = new QueueStream("localhost:50000", "some-client-Id");
            await Task.Delay(1000);
            try
            {
                List<Message> messages = new List<Message>();
                messages.Add( new Message()
                {
                    Queue = "q1",
                    Body = System.Text.Encoding.UTF8.GetBytes("hello kubemq - sending single message")
                });

                await client.Send(new SendRequest(messages));
                Console.WriteLine("Message sent");
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
            }
            try
            {
                PollRequest pollRequest = new PollRequest()
                {
                    Queue = "q1",
                    WaitTimeout = 1000,
                    MaxItems = 1,
                    AutoAck = true,
                };
                PollResponse response = await client.Poll(pollRequest);
                Console.WriteLine($"Messages Received: {response.Messages.Count}");
                foreach (Message message in response.Messages)
                {
                    string bodyAsString = System.Text.Encoding.UTF8.GetString(message.Body);
                    Console.WriteLine($"MessageID:{message.MessageID}, Body:{bodyAsString}");
                }
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
            }
        }
    }
}
